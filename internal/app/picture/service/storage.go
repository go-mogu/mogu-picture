package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
)

type IStorage interface {
	PageList(ctx context.Context, param model.Storage) (total int, result []*entity.Storage, err error)
	List(ctx context.Context, param entity.Storage) (result []*entity.Storage, err error)
	Get(ctx context.Context, uid string) (result *entity.Storage, err error)
	Add(ctx context.Context, in model.Storage) (err error)
	Edit(ctx context.Context, in model.Storage) (err error)
	EditState(ctx context.Context, ids []string, state int8) (err error)
	Delete(ctx context.Context, ids []string) (err error)
}

var localStorage IStorage

func Storage() IStorage {
	if localStorage == nil {
		panic("implement not found for interface IStorage, forgot register?")
	}
	return localStorage
}
func RegisterStorage(i IStorage) {
	localStorage = i
}
