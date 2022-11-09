package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Storage = cStorage{}
)

type cStorage struct{}

// PageList 分页查询存储信息表
func (c *cStorage) PageList(ctx context.Context, req *v1.StoragePageListReq) (res *v1.StoragePageListRes, err error) {
	total, list, err := service.Storage().PageList(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	res = &v1.StoragePageListRes{
		Total: total,
		Rows:  list,
	}
	return
}

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
