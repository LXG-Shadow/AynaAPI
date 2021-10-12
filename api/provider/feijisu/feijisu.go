package susudm

import (
	"AynaAPI/api/core"
	e2 "AynaAPI/api/core/e"
	"AynaAPI/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"regexp"
	"strings"
)

const (
	Host = "http://feijisu7.com"
)

type FeijisuDmVideo struct {
	Id         string   `json:"id"`
	Category   string   `json:"category"`
	EpId       string   `json:"ep_id"`
	Title      string   `json:"title"`
	PictureUrl string   `json:"pic"`
	Urls       []string `json:"urls"`
	Episodes   []string `json:"episodes"`
}

func InitDefault() *FeijisuDmVideo {
	return &FeijisuDmVideo{EpId: "1"}
}
func InitWithUid(uid string) (*FeijisuDmVideo, bool) {
	idString := regexp.MustCompile("feijisu-[0-9]+-[a-zA-z]+-[0-9]+").FindString(uid)
	if idString == "" {
		return nil, false
	}
	idstringL := strings.Split(idString, "-")
	return &FeijisuDmVideo{
		Id:       idstringL[1],
		Category: idstringL[2],
		EpId:     idstringL[3],
	}, true
}

func InitWithUrl(url string) (*FeijisuDmVideo, bool) {
	if v, b := InitWithUid(url); b {
		return v, b
	}
	if urlString := regexp.MustCompile("feijisu[0-9]\\.com/[a-zA-z]+/[0-9]+/").FindString(url); urlString != "" {
		if idString, b := utils.SliceString(urlString, 13, -1); b {
			idstringL := strings.Split(idString, "/")
			v := InitDefault()
			v.Id = idstringL[1]
			v.Category = idstringL[0]
			return v, true
		}
	}
	if urlString := regexp.MustCompile("feijisu[0-9]\\.com/[a-zA-z]+/[0-9]+/[0-9]+.html").FindString(url); urlString != "" {
		if idString, b := utils.SliceString(urlString, 13, -5); b {
			idstringL := strings.Split(idString, "/")
			return &FeijisuDmVideo{
				Id:       idstringL[1],
				Category: idstringL[0],
				EpId:     idstringL[2],
			}, true
		}
	}
	return nil, false
}

func GenerateUniqueId(id string, category string, epId string) string {
	return fmt.Sprintf("feijisu-%s-%s-%s", id, category, epId)
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *FeijisuDmVideo) MarshalJSON() ([]byte, error) {
	type FakeV FeijisuDmVideo
	return json.Marshal(&struct {
		Uid string `json:"uid"`
		*FakeV
	}{
		Uid:   self.GetUniqueId(),
		FakeV: (*FakeV)(self),
	})
}

func (self *FeijisuDmVideo) GetUniqueId() string {
	return fmt.Sprintf("feijisu-%s-%s-%s", self.Id, self.Category, self.EpId)
}

func (self *FeijisuDmVideo) Initialize() bool {
	resp := GetInfo(self.Id, self.Category, self.EpId)
	if resp.Status != e2.SUCCESS {
		return false
	}
	self.PictureUrl = cast.ToString(resp.Data["pic"])
	self.Title = cast.ToString(resp.Data["title"])
	self.Episodes = cast.ToStringSlice(resp.Data["episodes"])
	resp = GetPlayData(self.Id, self.EpId)
	if resp.Status != e2.SUCCESS {
		return false
	}
	self.Urls = cast.ToStringSlice(resp.Data["urls"])
	return true
}

func (self *FeijisuDmVideo) GetPlayUrls() []core.ApiResource {
	pUrls := make([]core.ApiResource, 0)
	for _, url := range self.Urls {
		pUrls = append(pUrls, core.ApiResource{
			Url:    url,
			Header: nil,
		})
	}
	return pUrls
}
