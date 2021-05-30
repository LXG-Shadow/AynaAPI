package susudm

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

const (
	Host = "http://www.susudm.com"
)

type SusuDmVideo struct {
	Id         string   `json:"id"`
	Category   string   `json:"category"`
	EpId       string   `json:"ep_id"`
	Title      string   `json:"title"`
	PictureUrl string   `json:"pic"`
	Urls       []string `json:"-"`
	Episodes   []string `json:"episodes"`
}

func InitDefault() *SusuDmVideo {
	return &SusuDmVideo{EpId: "1"}
}
func InitWithUid(uid string) (*SusuDmVideo, bool) {
	idString := regexp.MustCompile("susudm-[0-9]+-[a-zA-z]+-[0-9]+").FindString(uid)
	if idString == "" {
		return nil, false
	}
	idstringL := strings.Split(idString, "-")
	return &SusuDmVideo{
		Id:       idstringL[1],
		Category: idstringL[2],
		EpId:     idstringL[3],
	}, true
}

func InitWithUrl(url string) (*SusuDmVideo, bool) {
	if v, b := InitWithUid(url); b {
		return v, b
	}
	if urlString := regexp.MustCompile("susudm\\.com/[a-zA-z]+/[0-9]+/").FindString(url); urlString != "" {
		if idString, b := utils.SliceString(urlString, 11, -1); b {
			idstringL := strings.Split(idString, "/")
			v := InitDefault()
			v.Id = idstringL[1]
			v.Category = idstringL[0]
			return v, true
		}
	}
	if urlString := regexp.MustCompile("susudm\\.com/[a-zA-z]+/[0-9]+/[0-9]+.html").FindString(url); urlString != "" {
		if idString, b := utils.SliceString(urlString, 11, -5); b {
			idstringL := strings.Split(idString, "/")
			return &SusuDmVideo{
				Id:       idstringL[1],
				Category: idstringL[0],
				EpId:     idstringL[2],
			}, true
		}
	}
	return nil, false
}

func GenerateUniqueId(id string, category string, epId string) string {
	return fmt.Sprintf("susudm-%s-%s-%s", id, category, epId)
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *SusuDmVideo) MarshalJSON() ([]byte, error) {
	type FakeV SusuDmVideo
	return json.Marshal(&struct {
		Uid string `json:"uid"`
		*FakeV
	}{
		Uid:   self.GetUniqueId(),
		FakeV: (*FakeV)(self),
	})
}

func (self *SusuDmVideo) GetUniqueId() string {
	return fmt.Sprintf("susudm-%s-%s-%s", self.Id, self.Category, self.EpId)
}

func (self *SusuDmVideo) Initialize() bool {
	resp := GetInfo(self.Id, self.Category, self.EpId)
	if resp.Status != e.SUCCESS {
		return false
	}
	self.PictureUrl = cast.ToString(resp.Data["pic"])
	self.Title = cast.ToString(resp.Data["title"])
	self.Episodes = cast.ToStringSlice(resp.Data["episodes"])
	resp = GetPlayData(self.Id, self.EpId)
	if resp.Status != e.SUCCESS {
		return false
	}
	self.Urls = cast.ToStringSlice(resp.Data["urls"])
	return true
}

func (self *SusuDmVideo) GetPlayUrls() []model.ApiResource {
	pUrls := make([]model.ApiResource, 0)
	for _, url := range self.Urls {
		pUrls = append(pUrls, model.ApiResource{
			Url:    url,
			Header: nil,
		})
	}
	return pUrls
}