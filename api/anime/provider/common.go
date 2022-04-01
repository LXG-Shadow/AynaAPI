package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/anime/rule"
	"AynaAPI/api/core"
	"AynaAPI/api/e"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"strings"
)

type CommonProvider struct {
	Name          string
	BaseUrl       string
	InfoAPI       string
	SearchAPI     string
	SearchRules   rule.CommonSearchRules
	InfoRules     rule.CommonInfoRules
	PlaylistRules rule.CommonPlaylistRules
}

func (p *CommonProvider) GetName() string {
	return p.Name
}

func (p *CommonProvider) Search(keyword string) (anime.AnimeSearchResult, error) {
	return commonSearch(
		p.GetName(),
		fmt.Sprintf(p.SearchAPI, keyword),
		p.SearchRules)
}

func (p *CommonProvider) UpdateAnimeMeta(meta *anime.AnimeMeta) error {
	return commonUpdateMeta(
		p.GetName(),
		fmt.Sprintf(p.InfoAPI, meta.Provider.Url),
		p.InfoRules,
		meta)
}

func (p *CommonProvider) UpdateAnime(animee *anime.Anime) error {
	return commonUpdateAnime(
		fmt.Sprintf(p.InfoAPI, animee.Provider.Url),
		p.PlaylistRules,
		animee)
}

func commonSearch(providerName string, url string, rules rule.CommonSearchRules) (anime.AnimeSearchResult, error) {

	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(url, "utf-8"),
		deepcolor.GetCORS, nil, nil)

	if err != nil {
		return anime.AnimeSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	var sResults = make([]anime.AnimeMeta, 0)
	ids := result.GetList(rules.SearchID)
	titles := result.GetList(rules.SearchTitle)
	years := result.GetList(rules.SearchYear)
	tags := result.GetList(rules.SearchTag)
	covers := result.GetList(rules.SearchCover)
	desc := result.GetList(rules.SearchDesc)
	for index, id := range ids {
		meta := anime.AnimeMeta{
			Title:       titles[index],
			Year:        years[index],
			Tags:        strings.Split(tags[index], " "),
			Cover:       covers[index],
			Description: desc[index],
			Provider: core.ProviderMeta{
				Name: providerName,
				Url:  id,
			},
		}
		sResults = append(sResults, meta)
	}
	return anime.AnimeSearchResult{Result: sResults}, nil
}

func commonUpdateMeta(providerName string, url string, rules rule.CommonInfoRules, meta *anime.AnimeMeta) error {
	meta.Provider.Name = providerName
	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(url, "utf-8"),
		deepcolor.GetCORS,
		nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = result.GetSingle(rules.InfoTitle)
	meta.Year = result.GetSingle(rules.InfoYear)
	meta.Cover = result.GetSingle(rules.InfoCover)
	meta.Description = result.GetSingle(rules.InfoDesc)
	if rules.InfoTag.Type == deepcolor.ItemTypeSingle {
		meta.Tags = strings.Split(result.GetSingle(rules.InfoTag), "-")
	}
	if rules.InfoTag.Type == deepcolor.ItemTypeList {
		meta.Tags = result.GetList(rules.InfoTag)
	}
	return nil
}

func commonUpdateAnime(url string, rules rule.CommonPlaylistRules, animee *anime.Anime) error {
	result, err := deepcolor.Fetch(
		deepcolor.TentacleHTML(url, "utf-8"),
		deepcolor.GetCORS, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	ids := result.GetList(rules.Videos)
	playlistNames := result.GetList(rules.PlaylistNames)
	videoNames := result.GetList(rules.VideoNames)
	animee.Playlists = make([]anime.Playlist, 0)
	current_playlist_id := "-123ysdh5sgsdf54"
	current_playlist_index := -1
	for index, id := range ids {
		tmp := strings.Split(id, "-")
		if len(tmp) < 3 {
			continue
		}
		animeId, playlistId, epId := tmp[0], tmp[1], tmp[2]
		if playlistId != current_playlist_id {
			current_playlist_id = playlistId
			current_playlist_index = current_playlist_index + 1
			animee.Playlists = append(animee.Playlists, anime.Playlist{
				Name:   playlistNames[current_playlist_index],
				Videos: make([]anime.AnimeVideo, 0),
			})
		}
		animee.Playlists[current_playlist_index].Videos = append(animee.Playlists[current_playlist_index].Videos,
			anime.AnimeVideo{
				Title: videoNames[index],
				Url:   "",
				Provider: core.ProviderMeta{
					Name: "",
					Url:  strings.Join([]string{animeId, playlistId, epId}, "-"),
				},
			})
	}
	return nil
}
