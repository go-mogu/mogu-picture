package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	File = cFile{}
)

type cFile struct{}

// PageList 分页查询文件表
func (c *cFile) PageList(ctx context.Context, req *v1.FilePageListReq) (res *v1.FilePageListRes, err error) {
	total, list, err := service.File().PageList(ctx, req.File)
	if err != nil {
		return nil, err
	}
	res = &v1.FilePageListRes{
		Total: total,
		Rows:  list,
	}
	return
}

// List 列表查询文件表
func (c *cFile) List(ctx context.Context, req *v1.FileListReq) (res *v1.FileListRes, err error) {
	list, err := service.File().List(ctx, req.File)
	if err != nil {
		return nil, err
	}
	res = &v1.FileListRes{
		Rows: list,
	}
	return
}

// Get 查询文件表详情
func (c *cFile) Get(ctx context.Context, req *v1.FileGetReq) (res *v1.FileGetRes, err error) {
	entity, err := service.File().Get(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res = &v1.FileGetRes{
		File: entity,
	}
	return
}

// Add 添加文件表
func (c *cFile) Add(ctx context.Context, req *v1.FileAddReq) (res *v1.FileAddRes, err error) {
	in := new(model.File)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.File().Add(ctx, *in)
	return
}

// Edit 编辑文件表
func (c *cFile) Edit(ctx context.Context, req *v1.FileEditReq) (res *v1.FileEditRes, err error) {
	in := new(model.File)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.File().Edit(ctx, *in)
	return
}

// EditState 编辑文件表状态
func (c *cFile) EditState(ctx context.Context, req *v1.FileEditStateReq) (res *v1.FileEditRes, err error) {
	err = service.File().EditState(ctx, req.Ids, req.State)
	return
}

// Delete 删除文件表
func (c *cFile) Delete(ctx context.Context, req *v1.FileDelReq) (res *v1.FileDelRes, err error) {
	err = service.File().Delete(ctx, req.Ids)
	return
}
