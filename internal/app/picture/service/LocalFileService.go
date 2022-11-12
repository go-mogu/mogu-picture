package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ILocalFileService interface {
	BatchUploadFile(ctx context.Context, multipartFileList []*ghttp.UploadFile, fileSort entity.FileSort) (list []string, err error)
	UploadFile(ctx context.Context, newFileName string, multipartFile *ghttp.UploadFile, fileSort entity.FileSort) (result string, err error)
	UploadPictureByUrl(ctx context.Context, itemUrl string, newFileName string, fileSort entity.FileSort) (result string, err error)
}

var localFileService ILocalFileService

func LocalFileService() ILocalFileService {
	if localFileService == nil {
		panic("implement not found for interface ILocalFileService, forgot register?")
	}
	return localFileService
}
func RegisterLocalFileService(i ILocalFileService) {
	localFileService = i
}
