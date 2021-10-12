package imghost

import (
	"AynaAPI/api/core"
	e2 "AynaAPI/api/core/e"
	"AynaAPI/config"
	"AynaAPI/config/cookie"
	"encoding/base64"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

const BilibiliUploadApi = "https://api.bilibili.com/x/article/creative/article/upcover"

func UploadBilibili(imgBS64 string) core.ApiResponse {
	resp, err := grequests.Post(BilibiliUploadApi, &grequests.RequestOptions{
		Cookies: cookie.GetCookie("bilibili"),
		Data: map[string]string{
			"csrf":  config.APIConfig.Bilibili_JCT,
			"cover": imgBS64,
		}})
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}

	code := gjson.Get(resp.String(), "code").Int()
	if code != 0 {
		return core.CreateApiResponseByStatus(e2.EXTERNAL_API_ERROR, map[string]interface{}{
			"response": resp,
		})
	}
	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"url": gjson.Get(resp.String(), "data.url").String(),
	})
}

func UploadFileBilibili(imgPath string) core.ApiResponse {
	ff, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.ERROR_READFILE)
	}
	return UploadBilibili("data:image/png;base64," + base64.StdEncoding.EncodeToString(ff))
}
