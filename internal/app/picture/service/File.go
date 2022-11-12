package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IFile interface {
	CropperPicture(ctx context.Context, uploadFile *ghttp.UploadFile) (result []map[string]interface{}, err error)
	GetPicture(ctx context.Context, fileIds string, code string) (result []map[string]interface{}, err error)
	UploadFile(param model.UploadFileParam) (result *model.File, err error)
	BatchUploadFile(ctx context.Context, fileList []*model.UploadFileInfo, systemConfig baseModel.SystemConfig) (result []*model.File, err error)
	UploadPicsByUrl(ctx context.Context, fileVO model.FileVO) (result []*model.File, err error)
	CkeditorUploadFile(ctx context.Context) (result map[string]interface{}, err error)
	CkeditorUploadCopyFile(ctx context.Context) (result map[string]interface{}, err error)
	CkeditorUploadToolFile(ctx context.Context) (result map[string]interface{}, err error)
}

var localFile IFile

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}
func RegisterFile(i IFile) {
	localFile = i
}
