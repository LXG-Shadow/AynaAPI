package provider

import (
	"AynaAPI/api/anime/core"
	"AynaAPI/api/anime/rule"
	"AynaAPI/api/core/e"
	"AynaAPI/api/httpc"
	"AynaAPI/utils/vhttp"
	"AynaAPI/utils/vstring"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SusuDm struct {
	BaseUrl    string
	SearchAPI  string
	PlayUrlAPI string
	Rules      rule.SusuDmRules
}

func (p *SusuDm) GetName() string {
	return "susudm"
}

func (p *SusuDm) Validate(meta core.ProviderMeta) bool {
	return meta.Name == p.GetName() &&
		regexp.MustCompile("^"+regexp.QuoteMeta(p.BaseUrl)).FindString(meta.Url) != ""
}

func _newSusuDm() *SusuDm {
	return &SusuDm{
		BaseUrl:    "http://www.susudm.com",
		SearchAPI:  "http://119.29.15.48:12334/ssszz.php?q=%s&top=%s",
		PlayUrlAPI: "http://d.gqyy8.com:8077/ne2/s%d.js?%d",
		Rules:      rule.InitializeSusuDmRules(),
	}
}

var SusuDmAPI *SusuDm

func init() {
	SusuDmAPI = _newSusuDm()
}

func (p *SusuDm) getSearchApi(keyword string) string {
	return fmt.Sprintf(p.SearchAPI, keyword, 1)
}

func (p *SusuDm) getPlayUrlAPI(id string) string {
	intid, _ := strconv.Atoi(id)
	return fmt.Sprintf(p.PlayUrlAPI, intid, time.Now().Unix())
}

func (p *SusuDm) GetAnimeMeta(meta core.ProviderMeta) (core.AnimeMeta, error) {
	aMeta := core.AnimeMeta{
		Provider: meta,
	}
	err := p.UpdateAnimeMeta(&aMeta)
	return aMeta, err
}

func (p *SusuDm) UpdateAnimeMeta(meta *core.AnimeMeta) error {
	if regexp.MustCompile("/[0-9]+/").FindString(meta.Provider.Url) == "" {
		return e.NewError(e.INTERNAL_ERROR)
	}
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         meta.Provider.Url,
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = strings.TrimSpace(result.GetSingle(p.Rules.Title))
	meta.Cover = result.GetSingle(p.Rules.Cover)
	meta.Tags = result.GetList(p.Rules.Tags)
	meta.Description = result.GetSingle(p.Rules.Desc)
	tmp := strings.Split(result.GetSingle(p.Rules.AreaYear), "---")
	if len(tmp) < 2 {
		tmp = strings.Split(result.GetSingle(p.Rules.AreaYear2), "---")
	}
	if len(tmp) < 2 {
		meta.Year = "-1"
	} else {
		meta.Year = tmp[1]
	}
	return nil
}

func (p *SusuDm) GetAnime(meta core.AnimeMeta) (core.Anime, error) {
	anime := core.Anime{
		AnimeMeta: meta,
	}
	err := p.UpdateAnime(&anime)
	return anime, err
}

func (p *SusuDm) UpdateAnime(anime *core.Anime) error {
	id := regexp.MustCompile("/[0-9]+/").FindString(anime.Provider.Url)
	if id == "" {
		return e.NewError(e.INTERNAL_ERROR)
	}
	id, _ = vstring.SliceString(id, 1, -1)
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         p.getPlayUrlAPI(id),
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeText,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	if anime.Playlists == nil {
		anime.Playlists = make([]core.Playlist, 0)
	}
	rawtext := result.(deepcolor.TentacleTextResult).Data.(string)
	for i := 0; i < 10; i++ {
		var playlistID string
		if i == 0 {
			playlistID = ""
		} else {
			playlistID = fmt.Sprintf("_%d", i)
		}
		pattern := fmt.Sprintf("playarr%s\\[[0-9]+\\]=\"[^\"]*\";", playlistID)

		if datas := regexp.MustCompile(pattern).FindAllString(rawtext, -1); len(datas) > 0 {
			playlist := core.Playlist{
				Name:   fmt.Sprintf("%d", i),
				Videos: make([]core.AnimeVideo, 0),
			}
			for _, data := range datas {
				data = regexp.
					MustCompile(fmt.Sprintf("playarr%s\\[[0-9]+\\]=\"", playlistID)).
					ReplaceAllString(data, "")
				data = regexp.MustCompile("\";").ReplaceAllString(data, "")
				videoData := strings.Split(data, ",")
				playlist.Videos = append(playlist.Videos, core.AnimeVideo{
					Provider: core.ProviderMeta{
						Name: videoData[len(videoData)-2],
						Url:  strings.Join(videoData[0:len(videoData)-2], ","),
					},
					Title: videoData[len(videoData)-1],
					Url:   strings.Join(videoData[0:len(videoData)-2], ","),
				})
			}
			anime.Playlists = append(anime.Playlists, playlist)
		}
	}
	return nil
}

func (p *SusuDm) UpdateAnimeVideo(video *core.AnimeVideo) error {
	video.Url = video.Provider.Url
	return nil
}

func (p *SusuDm) Search(keyword string) (core.AnimeSearchResult, error) {
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         p.getSearchApi(keyword),
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeText,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return core.AnimeSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	var sResults = make([]core.AnimeMeta, 0)
	jsonResult := gjson.Parse(
		strings.ReplaceAll(result.(deepcolor.TentacleTextResult).Data.(string), "\ufeff", ""))
	jsonResult.ForEach(func(key, value gjson.Result) bool {
		pMeta := core.ProviderMeta{
			Name: p.GetName(),
			Url:  vhttp.JoinUrl(p.BaseUrl, vhttp.GetUrlPath("http://"+value.Get("url").String())),
		}
		aMeta := core.AnimeMeta{
			Title:    value.Get("title").String(),
			Cover:    value.Get("thumb").String(),
			Year:     value.Get("time").String(),
			Provider: pMeta,
		}
		sResults = append(sResults, aMeta)
		return true

	})
	return core.AnimeSearchResult{Result: sResults}, nil
}
