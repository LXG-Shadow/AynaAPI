package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/anime/rule"
	"AynaAPI/api/core"
	"AynaAPI/api/e"
	"AynaAPI/config"
	"AynaAPI/utils/vfile"
	"AynaAPI/utils/vhttp"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"regexp"
	"strings"
)

type Dldm struct {
	CommonProvider
	PlayUrlAPI1 string
	PlayUrlAPI2 string
	AesKey      []byte
	Rules       rule.DldmRules
}

func _newDldm(baseurl string) *Dldm {
	rules := rule.InitializeDldmRules()
	return &Dldm{
		CommonProvider: CommonProvider{
			Name:          "dldm",
			BaseUrl:       baseurl,
			InfoAPI:       baseurl + "/index.php/vod/detail/id/%s.html",
			SearchAPI:     baseurl + "/index.php/vod/search.html?wd=%s",
			SearchRules:   rules.CommonSearchRules,
			InfoRules:     rules.CommonInfoRules,
			PlaylistRules: rules.CommonPlaylistRules,
		},
		PlayUrlAPI1: baseurl + "/index.php/vod/play/id/%s/sid/%s/nid/%s.html",
		PlayUrlAPI2: baseurl + "/v/?url=%s",
		AesKey:      []byte("A42EAC0C2B40847B"),
		Rules:       rules,
	}
}

var DldmAPI *Dldm

func init() {
	if config.APIConfig == nil {
		DldmAPI = _newDldm("https://www.dqsj.cc")
	} else {
		DldmAPI = _newDldm(config.APIConfig.AnimeDldmBaseUrl)
	}
	anime.Providers.Add(DldmAPI.GetName(), DldmAPI)
}

func (p *Dldm) Validate(meta core.ProviderMeta) bool {
	return meta.Name == p.GetName() && regexp.MustCompile("^[0-9]+$").FindString(meta.Url) != ""
}

func (p *Dldm) GetAnimeMeta(meta core.ProviderMeta) (anime.AnimeMeta, error) {
	aMeta := anime.AnimeMeta{Provider: meta}
	if !p.Validate(meta) {
		return aMeta, e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	err := p.UpdateAnimeMeta(&aMeta)
	return aMeta, err
}

func (p *Dldm) GetAnime(meta anime.AnimeMeta) (anime.Anime, error) {
	animee := anime.Anime{AnimeMeta: meta}
	err := p.UpdateAnime(&animee)
	return animee, err
}

func (p *Dldm) UpdateAnimeVideo(video *anime.AnimeVideo) error {
	tmp := strings.Split(video.Provider.Url, "-")
	if len(tmp) < 3 {
		return e.NewError(e.PROVIDER_META_NOT_VALIED)
	}
	url := fmt.Sprintf(p.PlayUrlAPI1, tmp[0], tmp[1], tmp[2])
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         url,
		Charset:     "utf-8",
		ContentType: deepcolor.ResultTypeText,
	}, deepcolor.GetCORS, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	videoid := result.GetSingle(p.Rules.VideoId)
	if videoid == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	videoid = vhttp.QueryUnescapeWithEncoding(videoid, "utf-8")
	if vfile.GetFileExt(videoid) == ".m3u8" {
		video.Url = videoid
		return nil
	}
	url = fmt.Sprintf(p.PlayUrlAPI2, videoid)
	result, err = deepcolor.Fetch(deepcolor.Tentacle{
		Url:         url,
		Charset:     "utf-8",
		ContentType: deepcolor.ResultTypeText,
	}, deepcolor.GetCORS, nil, nil)
	aesiv := result.GetSingle(p.Rules.VideoAesIv)
	if aesiv == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	iv := []byte(aesiv)
	block, err := aes.NewCipher(p.AesKey)
	if err != nil {
		return e.NewError(e.INTERNAL_ERROR)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	raw_text := result.GetSingle(p.Rules.VideoEncUrl)
	if aesiv == "" {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(raw_text)
	if err != nil {
		return e.NewError(e.INTERNAL_ERROR)
	}
	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("ciphertext is not a multiple of the block size")
	}
	mode.CryptBlocks(ciphertext, ciphertext)
	// unpad and decode
	length := len(ciphertext)
	video.Url = string(ciphertext[:length-int(ciphertext[length-1])])
	video.Url = regexp.MustCompile("\\\\/").ReplaceAllString(video.Url, "/")
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	return nil
}

//func (p *Dldm) UpdateAnimeVideo(video *anime.AnimeVideo) error {
//	tmp := strings.Split(video.Provider.Url, "-")
//	if len(tmp) < 3 {
//		return e.NewError(e.PROVIDER_META_NOT_VALIED)
//	}
//	url := fmt.Sprintf(p.PlayUrlAPI1, tmp[0], tmp[1], tmp[2])
//	result, err := deepcolor.Fetch(deepcolor.Tentacle{
//		Url:         url,
//		Charset:     "utf-8",
//		ContentType: deepcolor.ResultTypeText,
//	}, deepcolor.GetCORS, nil, nil)
//	if err != nil {
//		return e.NewError(e.EXTERNAL_API_ERROR)
//	}
//	videoid := result.GetSingle(p.Rules.VideoId)
//	if videoid == "" {
//		return e.NewError(e.EXTERNAL_API_ERROR)
//	}
//	url = fmt.Sprintf(p.PlayUrlAPI2, videoid)
//	result, err = deepcolor.Fetch(deepcolor.Tentacle{
//		Url:         url,
//		Charset:     "utf-8",
//		ContentType: deepcolor.ResultTypeText | deepcolor.ResultTypeJson,
//	}, deepcolor.GetCORS, nil, nil)
//	realUrl := result.GetSingle(p.Rules.VideoUrl)
//	if realUrl == "" {
//		return e.NewError(e.EXTERNAL_API_ERROR)
//	}
//	video.Url = realUrl
//	return nil
//}
