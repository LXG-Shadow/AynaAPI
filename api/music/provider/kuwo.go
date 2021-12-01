package provider

import (
	"AynaAPI/api/core"
	"AynaAPI/api/e"
	"AynaAPI/api/httpc"
	"AynaAPI/api/music"
	"AynaAPI/utils/vhttp"
	"AynaAPI/utils/vstring"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"regexp"
)

type Kuwo struct {
	InfoApi      string
	FileApi      string
	SearchCookie string
	SearchApi    string
}

func _newKuwo() *Kuwo {
	return &Kuwo{
		InfoApi: "http://www.kuwo.cn/api/www/music/musicInfo?mid=%d&httpsStatus=1",
		//FileApi:      "http://www.kuwo.cn/api/v1/www/music/playUrl?mid=%d&type=music&httpsStatus=1",
		FileApi:      "http://antiserver.kuwo.cn/anti.s?type=convert_url&format=mp3&response=url&rid=MUSIC_%d",
		SearchCookie: "http://kuwo.cn/search/list?key=%s",
		SearchApi:    "http://www.kuwo.cn/api/www/search/searchMusicBykeyWord?key=%s&pn=%d&rn=%d",
	}
}

var KuwoMusicAPI *Kuwo

func (k *Kuwo) getInfoApi(songId int) string {
	return fmt.Sprintf(k.InfoApi, songId)
}

func (k *Kuwo) getFileApi(songId int) string {
	return fmt.Sprintf(k.FileApi, songId)
}

func (k *Kuwo) getSearchApi(keyword string) string {
	return fmt.Sprintf(k.SearchApi, vhttp.QueryEscapeWithEncoding(keyword, "utf-8"), 1, 64)
}

func init() {
	KuwoMusicAPI = _newKuwo()
	music.Providers.Add(KuwoMusicAPI.GetName(), KuwoMusicAPI)
}

func (k *Kuwo) parseSongId(url string) int {
	s := regexp.MustCompile("kuwo[0-9]+").FindString(url)
	if s == "" {
		return -1
	}
	if s, ok := vstring.SliceString(s, 4, 0); ok {
		return cast.ToInt(s)
	}
	return -1
}

func (k *Kuwo) GetName() string {
	return "kuwomusic"
}

func (k *Kuwo) Validate(meta core.ProviderMeta) bool {
	return meta.Name == k.GetName() && k.parseSongId(meta.Url) != -1
}

func (k *Kuwo) httpGet(url string) string {
	searchCookie := httpc.Get(fmt.Sprintf(k.SearchCookie, "any"), nil).RawResponse.Header.Get("set-cookie")
	kwToken, ok := vstring.SliceString(regexp.MustCompile("kw_token=([^;])*;").FindString(searchCookie), 9, -1)
	if !ok {
		return ""
	}
	return httpc.Get(url, map[string]string{
		"cookie":  "kw_token=" + kwToken,
		"csrf":    kwToken,
		"referer": "http://www.kuwo.cn/",
	}).String()
}

func (k *Kuwo) Search(keyword string) (music.MusicSearchResult, error) {
	resp := k.httpGet(k.getSearchApi(keyword))
	if resp == "" {
		return music.MusicSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	var result music.MusicSearchResult = music.MusicSearchResult{Result: make([]music.MusicMeta, 0)}
	gjson.Parse(resp).Get("data.list").ForEach(func(key, value gjson.Result) bool {
		result.Result = append(result.Result, music.MusicMeta{
			Title:  value.Get("name").String(),
			Cover:  value.Get("pic").String(),
			Artist: value.Get("artist").String(),
			Album:  value.Get("album").String(),
			Provider: core.ProviderMeta{
				Name: k.GetName(),
				Url:  fmt.Sprintf("kuwo%s", value.Get("rid").String()),
			},
		})
		return true
	})
	return result, nil
}

func (k *Kuwo) GetMusicMeta(meta core.ProviderMeta) (music.MusicMeta, error) {
	mMeta := music.MusicMeta{Provider: meta}
	if !k.Validate(meta) {
		return mMeta, e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	err := k.UpdateMusicMeta(&mMeta)
	return mMeta, err
}

func (k *Kuwo) UpdateMusicMeta(meta *music.MusicMeta) error {
	resp := k.httpGet(k.getInfoApi(k.parseSongId(meta.Provider.Url)))
	if resp == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	jresp := gjson.Parse(resp)
	if jresp.Get("data.musicrid").String() == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = jresp.Get("data.name").String()
	meta.Cover = jresp.Get("data.pic").String()
	meta.Artist = jresp.Get("data.artist").String()
	meta.Album = jresp.Get("data.album").String()
	return nil
}

func (k *Kuwo) GetMusic(meta music.MusicMeta) (music.Music, error) {
	musicc := music.Music{
		MusicMeta: meta,
	}
	err := k.UpdateMusic(&musicc)
	return musicc, err
}

func (k *Kuwo) UpdateMusic(musicc *music.Music) error {
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

func (k *Kuwo) UpdateMusicAudio(audio *music.MusicAudio) error {
	result := httpc.Get(k.getFileApi(k.parseSongId(audio.Provider.Url)), nil).String()
	if result == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	audio.Url = result
	return nil
}
