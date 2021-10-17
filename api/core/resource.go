package core

type ApiResource struct {
	Url    string            `json:"url"`
	Header map[string]string `json:"header"`
}
