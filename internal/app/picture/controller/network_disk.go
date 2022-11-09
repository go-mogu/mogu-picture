package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	NetworkDisk = cNetworkDisk{}
)

type cNetworkDisk struct{}

// PageList 分页查询网盘文件表
func (c *cNetworkDisk) PageList(ctx context.Context, req *v1.NetworkDiskPageListReq) (res *v1.NetworkDiskPageListRes, err error) {
	total, list, err := service.NetworkDisk().PageList(ctx, req.NetworkDisk)
	if err != nil {
		return nil, err
	}
	res = &v1.NetworkDiskPageListRes{
		Total: total,
		Rows:  list,
	}
	return
}

// List 列表查询网盘文件表
func (c *cNetworkDisk) List(ctx context.Context, req *v1.NetworkDiskListReq) (res *v1.NetworkDiskListRes, err error) {
	list, err := service.NetworkDisk().List(ctx, req.NetworkDisk)
	if err != nil {
		return nil, err
	}
	res = &v1.NetworkDiskListRes{
		Rows: list,
	}
	return
}

// Get 查询网盘文件表详情
func (c *cNetworkDisk) Get(ctx context.Context, req *v1.NetworkDiskGetReq) (res *v1.NetworkDiskGetRes, err error) {
	entity, err := service.NetworkDisk().Get(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res = &v1.NetworkDiskGetRes{
		NetworkDisk: entity,
	}
	return
}

// Add 添加网盘文件表
func (c *cNetworkDisk) Add(ctx context.Context, req *v1.NetworkDiskAddReq) (res *v1.NetworkDiskAddRes, err error) {
	in := new(model.NetworkDisk)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.NetworkDisk().Add(ctx, *in)
	return
}

// Edit 编辑网盘文件表
func (c *cNetworkDisk) Edit(ctx context.Context, req *v1.NetworkDiskEditReq) (res *v1.NetworkDiskEditRes, err error) {
	in := new(model.NetworkDisk)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.NetworkDisk().Edit(ctx, *in)
	return
}

// EditState 编辑网盘文件表状态
func (c *cNetworkDisk) EditState(ctx context.Context, req *v1.NetworkDiskEditStateReq) (res *v1.NetworkDiskEditRes, err error) {
	err = service.NetworkDisk().EditState(ctx, req.Ids, req.State)
	return
}

// Delete 删除网盘文件表
func (c *cNetworkDisk) Delete(ctx context.Context, req *v1.NetworkDiskDelReq) (res *v1.NetworkDiskDelRes, err error) {
	err = service.NetworkDisk().Delete(ctx, req.Ids)
	return
}
