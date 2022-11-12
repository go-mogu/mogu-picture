package service

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
)

type ICloudStorage interface {
	UploadFile(param model.UploadFileParam) (url string, err error)
	DeleteFile(ctx context.Context, fileName string, systemConfig baseModel.SystemConfig) (err error)
}

var abstractCloudStorage = map[string]ICloudStorage{}

func CloudStorage(key string) ICloudStorage {
	if abstractCloudStorage[key] == nil {
		panic(fmt.Sprintf("implement not found for interface ICloudStorage[%s], forgot register?", key))
	}
	return abstractCloudStorage[key]
}
func RegisterCloudStorage(key string, i ICloudStorage) {
	abstractCloudStorage[key] = i
}
