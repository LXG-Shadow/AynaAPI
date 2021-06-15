package yhdm

import (
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"regexp"
	"strings"
)

const Host string = "http://www.yhdm.so"

type YhdmVideo struct {
	Id         string   `json:"id"`
	EpId       string   `json:"ep_id"`
	Title      string   `json:"title"`
	PictureUrl string   `json:"pic"`
	Url        string   `json:"urls"`
	Episodes   []string `json:"episodes"`
}

func InitDefault() *YhdmVideo {
	return &YhdmVideo{
		Id:   "0",
		EpId: "1",
	}
}

func InitWithUid(uid string) (*YhdmVideo, bool) {
	idString := regexp.MustCompile("yhdm-[0-9]+-[^-]+").FindString(uid)
	if idString == "" {
		return nil, false
	}
	idstringL := strings.Split(idString, "-")
	return &YhdmVideo{
		Id:   idstringL[1],
		EpId: idstringL[2],
	}, true
}

func InitWithUrl(url string) (*YhdmVideo, bool) {
	if v, b := InitWithUid(url); b {
		return v, b
	}
	if urlString := regexp.MustCompile("yhdm\\.so/show/[0-9]+\\.html").FindString(url); urlString != "" {
		if id, b := utils.SliceString(urlString, 13, -5); b {
			v := InitDefault()
			v.Id = id
			return v, true
		}
	}
	if urlString := regexp.MustCompile("yhdm\\.so/v/[0-9]+-[^-]+\\.html").FindString(url); urlString != "" {
		if idString, b := utils.SliceString(urlString, 10, -5); b {
			idstringL := strings.Split(idString, "-")
			return &YhdmVideo{
				Id:   idstringL[0],
				EpId: idstringL[1],
			}, true
		}
	}
	return nil, false
}

func GenerateUniqueId(id string, epId string) string {
	return fmt.Sprintf("yhdm-%s-%s", id, epId)
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *YhdmVideo) MarshalJSON() ([]byte, error) {
	type FakeV YhdmVideo
	return json.Marshal(&struct {
		Uid string `json:"uid"`
		*FakeV
	}{
		Uid:   self.GetUniqueId(),
		FakeV: (*FakeV)(self),
	})
}

func (self *YhdmVideo) GetUniqueId() string {
	return fmt.Sprintf("yhdm-%s-%s", self.Id, self.EpId)
}

func (self *YhdmVideo) Initialize() bool {
	resp := GetInfo(self.Id, self.EpId)
	if resp.Status != e.SUCCESS {
		return false
	}
	self.PictureUrl = cast.ToString(resp.Data["pic"])
	self.Title = cast.ToString(resp.Data["title"])
	self.Episodes = cast.ToStringSlice(resp.Data["episodes"])
	self.Url = cast.ToStringMapString(resp.Data)["playUrl"]
	return true
}

func (self *YhdmVideo) GetPlayUrls() []model.ApiResource {
	pUrls := make([]model.ApiResource, 0)
	resp := ResolveVideoUrl(self.Url)
	if resp.Status == e.SUCCESS {
		pUrls = append(pUrls, model.ApiResource{
			Url:    cast.ToString(resp.Data["realUrl"]),
			Header: nil,
		})
	}
	return pUrls
}
