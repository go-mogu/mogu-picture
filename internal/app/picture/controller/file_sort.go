package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	FileSort = cFileSort{}
)

type cFileSort struct{}

// PageList 分页查询文件分类表
func (c *cFileSort) PageList(ctx context.Context, req *v1.FileSortPageListReq) (res *v1.FileSortPageListRes, err error) {
	total, list, err := service.FileSort().PageList(ctx, req.FileSort)
	if err != nil {
		return nil, err
	}
	res = &v1.FileSortPageListRes{
		Total: total,
		Rows:  list,
	}
	return
}

// List 列表查询文件分类表
func (c *cFileSort) List(ctx context.Context, req *v1.FileSortListReq) (res *v1.FileSortListRes, err error) {
	list, err := service.FileSort().List(ctx, req.FileSort)
	if err != nil {
		return nil, err
	}
	res = &v1.FileSortListRes{
		Rows: list,
	}
	return
}

// Get 查询文件分类表详情
func (c *cFileSort) Get(ctx context.Context, req *v1.FileSortGetReq) (res *v1.FileSortGetRes, err error) {
	entity, err := service.FileSort().Get(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res = &v1.FileSortGetRes{
		FileSort: entity,
	}
	return
}

// Add 添加文件分类表
func (c *cFileSort) Add(ctx context.Context, req *v1.FileSortAddReq) (res *v1.FileSortAddRes, err error) {
	in := new(model.FileSort)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.FileSort().Add(ctx, *in)
	return
}

// Edit 编辑文件分类表
func (c *cFileSort) Edit(ctx context.Context, req *v1.FileSortEditReq) (res *v1.FileSortEditRes, err error) {
	in := new(model.FileSort)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.FileSort().Edit(ctx, *in)
	return
}

// EditState 编辑文件分类表状态
func (c *cFileSort) EditState(ctx context.Context, req *v1.FileSortEditStateReq) (res *v1.FileSortEditRes, err error) {
	err = service.FileSort().EditState(ctx, req.Ids, req.State)
	return
}

// Delete 删除文件分类表
func (c *cFileSort) Delete(ctx context.Context, req *v1.FileSortDelReq) (res *v1.FileSortDelRes, err error) {
	err = service.FileSort().Delete(ctx, req.Ids)
	return
}
