package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
)

type IFileSort interface {
	PageList(ctx context.Context, param model.FileSort) (total int, result []*entity.FileSort, err error)
	List(ctx context.Context, param entity.FileSort) (result []*entity.FileSort, err error)
	Get(ctx context.Context, uid string) (result *entity.FileSort, err error)
	Add(ctx context.Context, in model.FileSort) (err error)
	Edit(ctx context.Context, in model.FileSort) (err error)
	EditState(ctx context.Context, ids []string, state int8) (err error)
	Delete(ctx context.Context, ids []string) (err error)
}

var localFileSort IFileSort

func FileSort() IFileSort {
	if localFileSort == nil {
		panic("implement not found for interface IFileSort, forgot register?")
	}
	return localFileSort
}
func RegisterFileSort(i IFileSort) {
	localFileSort = i
}
