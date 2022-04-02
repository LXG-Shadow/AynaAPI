package httpc

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func Get(url string, header map[string]string) (*resty.Response, error) {
	resp, err := resty.New().
		SetTimeout(time.Second * 3).R().
		SetHeaders(header).
		Get(url)
	return resp, err

}

func GetBodyString(url string, header map[string]string) string {
	resp, err := Get(url, header)
	if err != nil {
		return ""
	}
	return resp.String()
}

func Head(url string, header map[string]string) (*resty.Response, error) {
	resp, err := resty.New().
		SetTimeout(time.Second * 3).R().
		SetHeaders(header).
		Head(url)
	return resp, err
}

//func Post(url string, header map[string]string, data map[string]string) *grequests.Response {
//	resp, err := grequests.Post(url, &grequests.RequestOptions{
//		Headers: header,
//	})
//	if err != nil {
//		log.Println("Unable to make request: ", err)
//	}
//	return resp
//}
