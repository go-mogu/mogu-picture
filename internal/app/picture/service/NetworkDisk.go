package service

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
)

type INetworkDisk interface {
	GetFileList(ctx context.Context, param model.NetworkDisk) (result []*model.NetworkDisk, err error)
	CreateFile(ctx context.Context, in entity.NetworkDisk) (err error)
	Edit(ctx context.Context, in model.NetworkDisk) (err error)
	EditState(ctx context.Context, ids []string, state int8) (err error)
	BatchDeleteFile(ctx context.Context, networkDiskList model.NetworkDiskList) (err error)
	DeleteFile(ctx context.Context, networkDisk model.NetworkDisk, config baseModel.SystemConfig) (err error)
	UnzipFile(ctx context.Context, disk model.NetworkDisk) (err error)
	MoveFile(ctx context.Context, networkDisk model.NetworkDisk) (err error)
	BatchMoveFile(ctx context.Context, networkDisk model.NetworkDisk) (err error)
	SelectFileByFileType(ctx context.Context, networkDisk model.NetworkDisk) (result model.NetworkDiskList, err error)
	GetFileTree(ctx context.Context) (result baseModel.TreeNode, err error)
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
