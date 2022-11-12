package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/go-mogu/mogu-picture/utility/RequestHolder"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	NetworkDisk = cNetworkDisk{}
)

type cNetworkDisk struct{}

// GetFileList 列表查询网盘文件表
func (c *cNetworkDisk) GetFileList(ctx context.Context, req *v1.GetFileListReq) (res *v1.GetFileListRes, err error) {
	RequestHolder.CheckLogin(ctx)
	utils.ErrIsNil(ctx, err)
	networkDisk := req.NetworkDisk
	filePath, err := gurl.Decode(networkDisk.FilePath)
	utils.ErrIsNil(ctx, err)
	networkDisk.FilePath = filePath
	list, err := service.NetworkDisk().GetFileList(ctx, networkDisk)
	utils.ErrIsNil(ctx, err)
	res = (*v1.GetFileListRes)(&list)
	return
}

// CreateFile 添加网盘文件表
func (c *cNetworkDisk) CreateFile(ctx context.Context, req *v1.CreateFileReq) (res v1.CreateFileRes, err error) {
	adminUid := RequestHolder.CheckLogin(ctx)
	in := new(entity.NetworkDisk)
	err = gconv.Struct(req, in)
	utils.ErrIsNil(ctx, err)
	in.AdminUid = adminUid
	err = service.NetworkDisk().CreateFile(ctx, *in)
	res = MessageConf.INSERT_SUCCESS
	return
}

// Edit 重命名文件
func (c *cNetworkDisk) Edit(ctx context.Context, req *v1.NetworkDiskEditReq) (res v1.NetworkDiskEditRes, err error) {
	RequestHolder.CheckLogin(ctx)
	err = service.NetworkDisk().Edit(ctx, *req.NetworkDisk)
	utils.ErrIsNil(ctx, err)
	res = MessageConf.UPDATE_SUCCESS
	return
}

// BatchDeleteFile 删除网盘文件表
func (c *cNetworkDisk) BatchDeleteFile(ctx context.Context, req *v1.BatchDeleteFileReq) (res v1.BatchDeleteFileRes, err error) {
	request := g.RequestFromCtx(ctx)
	list := make(model.NetworkDiskList, 0)
	err = request.Parse(&list)
	if err != nil {
		return "", err
	}
	err = service.NetworkDisk().BatchDeleteFile(ctx, list)
	utils.ErrIsNil(ctx, err, MessageConf.DELETE_FAIL)
	res = MessageConf.BATCH_DELETE_SUCCESS
	return
}

// DeleteFile 删除文件
func (c *cNetworkDisk) DeleteFile(ctx context.Context, req *v1.DeleteFileReq) (res v1.DeleteFileRes, err error) {
	RequestHolder.CheckLogin(ctx)
	config, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	err = service.NetworkDisk().DeleteFile(ctx, req.NetworkDisk, config)
	utils.ErrIsNil(ctx, err, MessageConf.DELETE_FAIL)
	res = MessageConf.DELETE_SUCCESS
	return
}

// UnzipFile 解压文件
func (c *cNetworkDisk) UnzipFile(ctx context.Context, req *v1.UnzipFileReq) (res v1.UnzipFileRes, err error) {
	RequestHolder.CheckLogin(ctx)
	err = service.NetworkDisk().UnzipFile(ctx, req.NetworkDisk)
	utils.ErrIsNil(ctx, err, "解压失败")
	res = MessageConf.OPERATION_SUCCESS
	return
}

// MoveFile 移动文件
func (c *cNetworkDisk) MoveFile(ctx context.Context, req *v1.MoveFileReq) (res v1.MoveFileRes, err error) {
	RequestHolder.CheckLogin(ctx)
	err = service.NetworkDisk().MoveFile(ctx, req.NetworkDisk)
	utils.ErrIsNil(ctx, err)
	res = MessageConf.OPERATION_SUCCESS
	return
}

// BatchMoveFile 批量移动文件
func (c *cNetworkDisk) BatchMoveFile(ctx context.Context, req *v1.BatchMoveFileReq) (res v1.BatchMoveFileRes, err error) {
	RequestHolder.CheckLogin(ctx)
	err = service.NetworkDisk().BatchMoveFile(ctx, req.NetworkDisk)
	utils.ErrIsNil(ctx, err)
	res = MessageConf.OPERATION_SUCCESS
	return
}

// SelectFileByFileType 通过文件类型查询文件
func (c *cNetworkDisk) SelectFileByFileType(ctx context.Context, req *v1.SelectFileByFileTypeReq) (res *v1.SelectFileByFileTypeRes, err error) {
	RequestHolder.CheckLogin(ctx)
	networkDiskList, err := service.NetworkDisk().SelectFileByFileType(ctx, req.NetworkDisk)
	utils.ErrIsNil(ctx, err)
	res = (*v1.SelectFileByFileTypeRes)(&networkDiskList)
	return
}

// GetFileTree 获取文件数
func (c *cNetworkDisk) GetFileTree(ctx context.Context, req *v1.GetFileTreeReq) (res *v1.GetFileTreeRes, err error) {
	RequestHolder.CheckLogin(ctx)
	treeNode, err := service.NetworkDisk().GetFileTree(ctx)
	utils.ErrIsNil(ctx, err)
	res = (*v1.GetFileTreeRes)(&treeNode)
	return
}
