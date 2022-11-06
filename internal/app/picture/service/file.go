package service

import (
	"context"
	"mogu-picture/internal/app/picture/model"
	"mogu-picture/internal/app/picture/model/entity"
)

type IFile interface {
	PageList(ctx context.Context, param model.File) (total int, result []*entity.File, err error)
	List(ctx context.Context, param entity.File) (result []*entity.File, err error)
	Get(ctx context.Context, uid string) (result *entity.File, err error)
	Add(ctx context.Context, in model.File) (err error)
	Edit(ctx context.Context, in model.File) (err error)
	EditState(ctx context.Context, ids []string, state int8) (err error)
	Delete(ctx context.Context, ids []string) (err error)
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
