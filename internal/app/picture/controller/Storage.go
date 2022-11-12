package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	"github.com/go-mogu/mogu-picture/utility/RequestHolder"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Storage = cStorage{}
)

type cStorage struct{}

// List 列表查询存储信息表
func (c *cStorage) List(ctx context.Context, req *v1.StorageListReq) (res *v1.StorageListRes, err error) {
	list, err := service.Storage().List(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	res = &v1.StorageListRes{
		Rows: list,
	}
	return
}

// Get 查询存储信息表详情
func (c *cStorage) Get(ctx context.Context, req *v1.StorageGetReq) (res *v1.StorageGetRes, err error) {
	entity, err := service.Storage().Get(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res = &v1.StorageGetRes{
		Storage: entity,
	}
	return
}

// Add 添加存储信息表
func (c *cStorage) Add(ctx context.Context, req *v1.StorageAddReq) (res *v1.StorageAddRes, err error) {
	in := new(model.Storage)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Storage().Add(ctx, *in)
	return
}

// Edit 编辑存储信息表
func (c *cStorage) Edit(ctx context.Context, req *v1.StorageEditReq) (res *v1.StorageEditRes, err error) {
	in := new(model.Storage)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Storage().Edit(ctx, *in)
	return
}

// EditState 编辑存储信息表状态
func (c *cStorage) EditState(ctx context.Context, req *v1.StorageEditStateReq) (res *v1.StorageEditRes, err error) {
	err = service.Storage().EditState(ctx, req.Ids, req.State)
	return
}

// Delete 删除存储信息表
func (c *cStorage) Delete(ctx context.Context, req *v1.StorageDelReq) (res *v1.StorageDelRes, err error) {
	err = service.Storage().Delete(ctx, req.Ids)
	return
}

// InitStorageSize 初始化容量大小
func (c *cStorage) InitStorageSize(ctx context.Context, req *v1.InitStorageSizeReq) (res v1.InitStorageSizeRes, err error) {
	err = service.Storage().InitStorageSize(ctx, req.AdminUid, req.MaxStorageSize)
	res = MessageConf.OPERATION_SUCCESS
	return
}

// EditStorageSize 编辑容量大小
func (c *cStorage) EditStorageSize(ctx context.Context, req *v1.EditStorageSizeReq) (res v1.EditStorageSizeRes, err error) {
	err = service.Storage().EditStorageSize(ctx, req.AdminUid, req.MaxStorageSize)
	res = MessageConf.OPERATION_SUCCESS
	return
}

// GetStorageByAdminUid 通过管理员uid，获取存储信息
func (c *cStorage) GetStorageByAdminUid(ctx context.Context, req *v1.GetStorageByAdminUidReq) (res v1.GetStorageByAdminUidRes, err error) {
	result, err := service.Storage().GetStorageByAdminUid(ctx, req.AdminUidList)
	res = result
	return
}

// GetStorage 通过管理员uid，获取存储信息
func (c *cStorage) GetStorage(ctx context.Context, req *v1.GetStorageReq) (res v1.GetStorageRes, err error) {
	request := RequestHolder.GetRequest(ctx)
	adminUid := request.Get(SysConf.ADMIN_UID).String()
	result, err := service.Storage().GetStorageByAdmin(ctx, adminUid)
	res = result
	return
}

// UploadFile 文件上传
func (c *cStorage) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res v1.UploadFileRes, err error) {
	RequestHolder.CheckLogin(ctx)
	result, err := service.Storage().UploadFile(ctx, req.NetworkDisk)
	res = v1.UploadFileRes(result)
	return
}
