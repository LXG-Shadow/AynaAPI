package web_service

import (
	"AynaAPI/server/model"
	"AynaAPI/server/pkg/upload"
	"AynaAPI/utils/vfile"
	"errors"
	"mime/multipart"
	"os"
)

func (svc *WebService) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*model.UploadFile, error) {
	md5Str, ext, ok := upload.GetMD5FileInfo(fileHeader)
	if !ok {
		return nil, errors.New("read file fail")
	}
	contentType, _ := vfile.GetFileHeaderContentType(fileHeader)

	if !upload.CheckContainExt(fileType, ext) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	filename := md5Str + ext

	dst := upload.GetFullSavePath(filename)
	accessUrl := upload.GetFullAccessUrl(filename)

	if f, _ := svc.Dao.GetUploadFile(md5Str, ""); f != nil {
		return f, errors.New("file already exists")
	}

	uploadFile := model.UploadFile{
		Filename:    fileHeader.Filename,
		FileHash:    md5Str,
		FileType:    int(fileType),
		ContentType: contentType,
		SavedPath:   dst,
		AccessUrl:   accessUrl,
	}

	if err := svc.Dao.CreateUploadFile(&uploadFile); err != nil {
		return nil, err
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	return &uploadFile, nil
}
