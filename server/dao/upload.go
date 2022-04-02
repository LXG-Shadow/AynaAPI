package dao

import "AynaAPI/server/model"

func (d *Dao) GetUploadFile(hashString, contentType string) (*model.UploadFile, error) {
	var file model.UploadFile
	err := d.engine.Where(model.UploadFile{FileHash: hashString, ContentType: contentType}).First(&file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (d *Dao) CreateUploadFile(file *model.UploadFile) error {
	return d.engine.Create(file).Error
}
