package model

type UploadFile struct {
	BaseModel
	Filename    string `json:"filename"`
	FileHash    string `json:"file_hash"`
	FileType    int    `json:"file_type"`
	ContentType string `json:"content_type"`
	SavedPath   string `json:"saved_path"`
	AccessUrl   string `json:"access_url"`
}
