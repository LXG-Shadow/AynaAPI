package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/anime/rule"
	"AynaAPI/api/core"
	"AynaAPI/api/e"
	"AynaAPI/api/httpc"
	"AynaAPI/config"
	"AynaAPI/utils/vhttp"
	"AynaAPI/utils/vstring"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"math"
	"regexp"
	"strings"
	"time"
)

type Agefans struct {
	BaseUrl    string
	InfoAPI    string
	SearchAPI  string
	PlayUrlAPI string
	Rules      rule.AgefansRules
}

func (p *Agefans) GetName() string {
	return "agefans"
}

func (p *Agefans) Validate(meta core.ProviderMeta) bool {
	return meta.Name == p.GetName() && regexp.MustCompile("^[0-9]+$").FindString(meta.Url) != ""
}

func _newAgefans(baseurl string) *Agefans {
	return &Agefans{
		BaseUrl:    baseurl,
		InfoAPI:    baseurl + "/detail/%s",
		SearchAPI:  baseurl + "/search?query=%s&page=%d",
		PlayUrlAPI: baseurl + "/_getplay?aid=%s&playindex=%s&epindex=%s",
		Rules:      rule.InitializeAgefansRules(),
	}
}

var AgefansAPI *Agefans

func init() {
	if config.APIConfig == nil {
		AgefansAPI = _newAgefans("https://www.agemys.com")
	} else {
		AgefansAPI = _newAgefans(config.APIConfig.AnimeAgefansBaseUrl)
	}
	anime.Providers.Add(AgefansAPI.GetName(), AgefansAPI)
}

func (p *Agefans) getInfoApi(aid string) string {
	return fmt.Sprintf(p.InfoAPI, aid)
}

func (p *Agefans) getSearchApi(keyword string) string {
	return fmt.Sprintf(p.SearchAPI, keyword, 1)
}

func (p *Agefans) getPlayUrlAPI(aid string, playindex string, epindex string) string {
	return fmt.Sprintf(p.PlayUrlAPI, aid, playindex, epindex)
}

func (p *Agefans) GetAnimeMeta(meta core.ProviderMeta) (anime.AnimeMeta, error) {
	aMeta := anime.AnimeMeta{Provider: meta}
	if !p.Validate(meta) {
		return aMeta, e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	err := p.UpdateAnimeMeta(&aMeta)
	return aMeta, err
}

func (p *Agefans) UpdateAnimeMeta(meta *anime.AnimeMeta) error {
	//id := regexp.MustCompile("/detail/[0-9]+").FindString(meta.Provider.Url)
	//if id == "" {
	//	return e.NewError(e.INTERNAL_ERROR)
	//}
	meta.Provider.Name = "agefans"
	//fmt.Println(p.getInfoApi(meta.Provider.Url))
	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(p.getInfoApi(meta.Provider.Url), "utf-8"),
		deepcolor.GetCORS,
		nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = result.GetSingle(p.Rules.InfoTitle)
	meta.Year = result.GetSingle(p.Rules.InfoYear)
	meta.Cover = result.GetSingle(p.Rules.InfoCover)
	meta.Description = result.GetSingle(p.Rules.InfoDesc)
	meta.Tags = strings.Split(result.GetSingle(p.Rules.InfoTag), " ")
	return nil
}

func (p *Agefans) GetAnime(meta anime.AnimeMeta) (anime.Anime, error) {
	animee := anime.Anime{AnimeMeta: meta}
	err := p.UpdateAnime(&animee)
	return animee, err
}
func (p *Agefans) UpdateAnime(animee *anime.Anime) error {
	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(p.getInfoApi(animee.Provider.Url), "utf-8"),
		deepcolor.GetCORS, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	ids := result.GetList(p.Rules.InfoVideos)
	urlNames := result.GetList(p.Rules.InfoVideoNames)
	animee.Playlists = make([]anime.Playlist, 0)
	current_playlist_id := "-1"
	current_playlist_index := -1
	for index, id := range ids {
		tmp := strings.Split(id, "-")
		if len(tmp) < 3 {
			continue
		}
		animeId, playlistId, epId := tmp[0], tmp[1], tmp[2]
		if playlistId != current_playlist_id {
			current_playlist_id = playlistId
			animee.Playlists = append(animee.Playlists, anime.Playlist{
				Name:   playlistId,
				Videos: make([]anime.AnimeVideo, 0),
			})
			current_playlist_index = len(animee.Playlists) - 1
		}
		animee.Playlists[current_playlist_index].Videos = append(animee.Playlists[current_playlist_index].Videos,
			anime.AnimeVideo{
				Title: urlNames[index],
				Url:   "",
				Provider: core.ProviderMeta{
					Name: "",
					Url:  strings.Join([]string{animeId, playlistId, epId}, "-"),
				},
			})
	}

	//for _, playlist := range anime.Playlists {
	//	for _, v := range playlist.Videos {
	//		err = p.UpdateAnimeVideo(&v)
	//		if err != nil{
	//		}
	//	}
	//}
	return nil
}

func (p *Agefans) getCookie(t1 int) string {
	timeNow := time.Now().UnixNano() / (1000000)
	t1Tmp := int64(math.Round(float64(t1)/1000)) >> 0x5
	k2 := (t1Tmp*(t1Tmp%0x1000)*0x3+0x1450f)*(t1Tmp%0x1000) + t1Tmp
	t2 := timeNow
	t2 = t2 - t2%10 + k2%10
	return fmt.Sprintf("t1=%d;k2=%d;t2=%d", t1, k2, t2)
}

func (p *Agefans) UpdateAnimeVideo(video *anime.AnimeVideo) error {
	tmp := strings.Split(video.Provider.Url, "-")
	if len(tmp) < 3 {
		return e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	url := p.getPlayUrlAPI(tmp[0], tmp[1], tmp[2])
	resp, err := httpc.Head(url, map[string]string{
		"referer": p.BaseUrl,
	})
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	initiator := regexp.MustCompile("t1=[^;]*;").FindString(resp.Header().Get("set-cookie"))

	if initiator == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	t1, _ := vstring.SliceString(initiator, 3, -1)

	authCookie := p.getCookie(cast.ToInt(t1))
	respString := httpc.GetBodyString(url, map[string]string{
		"referer": p.BaseUrl,
		"cookie":  authCookie,
	})
	video.Provider.Name = regexp.MustCompile("</?play>").
		ReplaceAllString(gjson.Parse(respString).Get("playid").String(), "")

	video.Url = vhttp.QueryUnescapeWithEncoding(gjson.Parse(respString).Get("vurl").String(), "utf-8")
	return nil
}

func (p *Agefans) Search(keyword string) (anime.AnimeSearchResult, error) {
	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(p.getSearchApi(keyword), "utf-8"),
		deepcolor.GetCORS, nil, nil)

	if err != nil {
		return anime.AnimeSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	var sResults = make([]anime.AnimeMeta, 0)
	urls := result.GetList(p.Rules.SearchURL)
	titles := result.GetList(p.Rules.SearchTitle)
	years := result.GetList(p.Rules.SearchYear)
	tags := result.GetList(p.Rules.SearchTag)
	covers := result.GetList(p.Rules.SearchCover)
	desc := result.GetList(p.Rules.SearchDesc)
	idreg := regexp.MustCompile("[0-9]+")
	for index, url := range urls {
		meta := anime.AnimeMeta{
			Title:       titles[index],
			Year:        years[index],
			Tags:        strings.Split(tags[index], " "),
			Cover:       covers[index],
			Description: desc[index],
			Provider: core.ProviderMeta{
				Name: p.GetName(),
				Url:  idreg.FindString(url),
			},
		}
		sResults = append(sResults, meta)
	}
	return anime.AnimeSearchResult{Result: sResults}, nil
}
