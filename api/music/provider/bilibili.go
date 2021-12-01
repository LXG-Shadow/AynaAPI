package provider

import (
	"AynaAPI/api/core"
	"AynaAPI/api/e"
	"AynaAPI/api/httpc"
	"AynaAPI/api/music"
	"AynaAPI/utils/vstring"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"regexp"
)

type BilibiliMusicQuality int

const (
	BilibiliMusicQualityHigh   BilibiliMusicQuality = 2
	BilibiliMusicQualityMedium BilibiliMusicQuality = 1
	BilibiliMusicQualityLow    BilibiliMusicQuality = 0
)

type BilibiliMusicQualityInfo struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
	BPS         string `json:"bps"`
}

var BilibiliMusicQualities = map[BilibiliMusicQuality]BilibiliMusicQualityInfo{
	BilibiliMusicQualityHigh: {
		Tag:         "HQ",
		Description: "高品质",
		BPS:         "320kbps",
	},
	BilibiliMusicQualityMedium: {
		Tag:         "HQ",
		Description: "标准",
		BPS:         "192kbps",
	},
	BilibiliMusicQualityLow: {
		Tag:         "Q",
		Description: "流畅",
		BPS:         "128kbps",
	},
}

type Bilibili struct {
	InfoApi   string
	FileApi   string
	SearchApi string
}

func _newBilibili() *Bilibili {
	return &Bilibili{
		InfoApi: "https://www.bilibili.com/audio/music-service-c/web/song/info?sid=%d",
		FileApi: "https://api.bilibili.com/audio/music-service-c/url?device=phone&mid=8047632&mobi_app=iphone&platform=ios&privilege=2&songid=%d&quality=%d",
		//FileApi:   "https://www.bilibili.com/audio/music-service-c/web/url?privilege=2&sid=%d&quality=%d",
		SearchApi: "https://api.bilibili.com/audio/music-service-c/s?search_type=music&keyword=%s&page=%d&pagesize=%d",
	}
}

var BiliMusicAPI *Bilibili

func init() {
	BiliMusicAPI = _newBilibili()
	music.Providers.Add(BiliMusicAPI.GetName(), BiliMusicAPI)
}

func (b *Bilibili) getInfoApi(sid int) string {
	return fmt.Sprintf(b.InfoApi, sid)
}

func (b *Bilibili) getFileApi(sid int, quality BilibiliMusicQuality) string {
	return fmt.Sprintf(b.FileApi, sid, quality)
}

func (b *Bilibili) getSearchApi(keyword string) string {
	return fmt.Sprintf(b.SearchApi, keyword, 1, 100)
}

func (b *Bilibili) parseSongId(url string) int {
	s := regexp.MustCompile("au[0-9]+").FindString(url)
	if s == "" {
		return -1
	}
	if s, ok := vstring.SliceString(s, 2, 0); ok {
		return cast.ToInt(s)
	}
	return -1
}

func (b *Bilibili) GetName() string {
	return "bilibilimusic"
}

func (b *Bilibili) Validate(meta core.ProviderMeta) bool {
	return meta.Name == b.GetName() && b.parseSongId(meta.Url) != -1
}

func (b *Bilibili) Search(keyword string) (music.MusicSearchResult, error) {
	resp := httpc.Get(b.getSearchApi(keyword), map[string]string{
		"user-agent": "BiliMusic/2.233.3",
	})
	if resp.String() == "" {
		return music.MusicSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	result := music.MusicSearchResult{Result: make([]music.MusicMeta, 0)}
	gjson.Get(resp.String(), "data.result").ForEach(func(key, value gjson.Result) bool {
		result.Result = append(result.Result, music.MusicMeta{
			Title:  value.Get("title").String(),
			Cover:  value.Get("cover").String(),
			Artist: value.Get("author").String(),
			Provider: core.ProviderMeta{
				Name: b.GetName(),
				Url:  fmt.Sprintf("au%s", value.Get("id").String()),
			},
		})
		return true
	})
	return result, nil
}

func (b *Bilibili) GetMusicMeta(meta core.ProviderMeta) (music.MusicMeta, error) {
	mMeta := music.MusicMeta{Provider: meta}
	if !b.Validate(meta) {
		return mMeta, e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	err := b.UpdateMusicMeta(&mMeta)
	return mMeta, err
}

func (b *Bilibili) UpdateMusicMeta(meta *music.MusicMeta) error {
	resp := httpc.Get(b.getInfoApi(b.parseSongId(meta.Provider.Url)), map[string]string{
		"user-agent": "BiliMusic/2.233.3",
	})
	if resp.String() == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	if gjson.Get(resp.String(), "data.title").String() == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = gjson.Get(resp.String(), "data.title").String()
	meta.Cover = gjson.Get(resp.String(), "data.cover").String()
	meta.Artist = gjson.Get(resp.String(), "data.author").String()
	meta.Album = meta.Title
	return nil
}

func (b *Bilibili) GetMusic(meta music.MusicMeta) (music.Music, error) {
	musicc := music.Music{
		MusicMeta: meta,
	}
	err := b.UpdateMusic(&musicc)
	return musicc, err
}

func (b *Bilibili) UpdateMusic(musicc *music.Music) error {
	//resp := httpc.Get(b.getFileApi(b.parseSongId(musicc.MusicMeta.Provider.Url),BilibiliMusicQualityHigh), map[string]string{
	//	"user-agent":"BiliMusic/2.233.3",
	//})
	//if resp.String() == ""{
	//	return e.NewError(e.EXTERNAL_API_ERROR)
	//}
	//
	//url := gjson.Get(resp.String(),"data.cdns.0").String()
	//if url == "" {
	//	return e.NewError(e.EXTERNAL_API_ERROR)
	//}
	musicc.Audio = music.MusicAudio{
		Url:       "",
		UserAgent: "BiliMusic/2.233.3",
		Provider: core.ProviderMeta{
			Name: "",
			Url:  musicc.MusicMeta.Provider.Url,
		},
	}
	musicc.Lyric = music.MusicLyric{}
	return nil
}

func (b *Bilibili) UpdateMusicAudio(audio *music.MusicAudio) error {
	resp := httpc.Get(b.getFileApi(b.parseSongId(audio.Provider.Url), BilibiliMusicQualityHigh), map[string]string{
		"user-agent": audio.UserAgent,
	})
	if resp.String() == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}

	url := gjson.Get(resp.String(), "data.cdns.0").String()
	if url == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	audio.Url = url
	audio.Type = BilibiliMusicQualities[BilibiliMusicQuality(gjson.Get(resp.String(), "data.type").Int())].Description
	audio.Size = int(gjson.Get(resp.String(), "data.size").Int())
	return nil
}
