package controller

import (
	"context"
	"github.com/go-mogu/mogu-picture/api/picture/v1"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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

// CropperPicture 截图上传
func (c *cFile) CropperPicture(ctx context.Context, req *v1.CropperPictureReq) (res *v1.CropperPictureRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("请选择图片")
		return
	}
	multipartFileList := []*ghttp.UploadFile{file}
	picture, err := service.File().CropperPicture(ctx, multipartFileList)
	res = (*v1.CropperPictureRes)(&picture)
	return
}

// GetPicture 获取文件的信息接口
// fileIds 获取文件信息的ids
// code ids用什么分割的，默认“,”
func (c *cFile) GetPicture(ctx context.Context, req *v1.GetPictureReq) (res *v1.GetPictureRes, err error) {
	result, err := service.File().GetPicture(ctx, req.FileIds, req.Code)
	if err != nil {
		return nil, err
	}
	res = (*v1.GetPictureRes)(&result)
	return
}

// UploadPics 多文件上传
// 上传图片接口   传入 userId sysUserId ,有那个传哪个，记录是谁传的,
// projectName 传入的项目名称如 base 默认是base
// sortName 传入的模块名， 如 admin，user ,等，不在数据库中记录的是不会上传的
func (c *cFile) UploadPics(ctx context.Context, req *v1.UploadPicsReq) (res *v1.UploadPicsRes, err error) {
	systemConfig, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	r := g.RequestFromCtx(ctx)
	uploadFiles := r.GetUploadFiles("filedatas")
	result, err := service.File().BatchUploadFile(ctx, uploadFiles, systemConfig)
	if err != nil {
		return nil, err
	}
	res = (*v1.UploadPicsRes)(&result)
	return
}

// UploadPicsByUrl 通过URL上传图片
// 通过URL将图片上传到自己服务器中【主要用于Github和Gitee的头像上传】
func (c *cFile) UploadPicsByUrl(ctx context.Context, req *v1.UploadPicsByUrlReq) (res *v1.UploadPicsByUrlRes, err error) {
	result, err := service.File().UploadPicsByUrl(ctx, req.FileVO)
	if err != nil {
		return nil, err
	}
	res = (*v1.UploadPicsByUrlRes)(&result)
	return
}
