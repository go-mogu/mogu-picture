package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
)

type INetworkDisk interface {
	PageList(ctx context.Context, param model.NetworkDisk) (total int, result []*entity.NetworkDisk, err error)
	List(ctx context.Context, param entity.NetworkDisk) (result []*entity.NetworkDisk, err error)
	Get(ctx context.Context, uid string) (result *entity.NetworkDisk, err error)
	Add(ctx context.Context, in model.NetworkDisk) (err error)
	Edit(ctx context.Context, in model.NetworkDisk) (err error)
	EditState(ctx context.Context, ids []string, state int8) (err error)
	Delete(ctx context.Context, ids []string) (err error)
}

var localNetworkDisk INetworkDisk

func NetworkDisk() INetworkDisk {
	if localNetworkDisk == nil {
		panic("implement not found for interface INetworkDisk, forgot register?")
	}
	return localNetworkDisk
}
func RegisterNetworkDisk(i INetworkDisk) {
	localNetworkDisk = i
}
