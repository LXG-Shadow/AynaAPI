package imghost

import (
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/config"
	"AynaAPI/config/cookie"
	"encoding/base64"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

const BilibiliUploadApi = "https://api.bilibili.com/x/article/creative/article/upcover"

func UploadBilibili(imgBS64 string) model.ApiResponse {
	resp, err := grequests.Post(BilibiliUploadApi, &grequests.RequestOptions{
		Cookies: cookie.GetCookie("bilibili"),
		Data: map[string]string{
			"csrf":  config.APIConfig.Bilibili_JCT,
			"cover": imgBS64,
		}})
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}

	code := gjson.Get(resp.String(), "code").Int()
	if code != 0 {
		return model.CreateApiResponseByStatus(e.EXTERNAL_API_ERROR, map[string]interface{}{
			"response": resp,
		})
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"url": gjson.Get(resp.String(), "data.url").String(),
	})
}

func UploadFileBilibili(imgPath string) model.ApiResponse {
	ff, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.ERROR_READFILE)
	}
	return UploadBilibili("data:image/png;base64," + base64.StdEncoding.EncodeToString(ff))
}
