package httpc

import (
	"github.com/levigross/grequests"
	"log"
)

func Get(url string, header map[string]string) *grequests.Response {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: header,
	})
	if err != nil {
		log.Println("Unable to make request: ", err)
	}
	return resp

}

func Head(url string, header map[string]string) *grequests.Response {
	resp, err := grequests.Head(url, &grequests.RequestOptions{
		Headers: header,
	})
	if err != nil {
		log.Println("Unable to make request: ", err)
	}
	return resp

}

func Post(url string, header map[string]string, data map[string]string) *grequests.Response {
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: header,
	})
	if err != nil {
		log.Println("Unable to make request: ", err)
	}
	return resp
}
