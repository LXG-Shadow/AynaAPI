package model

type ApiResource struct {
	Url    string            `json:"url"`
	Header map[string]string `json:"header"`
}
