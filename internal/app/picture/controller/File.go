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
)

var (
	File = cFile{}
)

type cFile struct{}

// CropperPicture 截图上传
func (c *cFile) CropperPicture(ctx context.Context, req *v1.CropperPictureReq) (res *v1.CropperPictureRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("请选择图片")
		return
	}
	picture, err := service.File().CropperPicture(ctx, file)
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
	fileList := make([]*model.UploadFileInfo, 0)
	for _, uploadFile := range uploadFiles {
		file, err := uploadFile.Open()
		if err != nil {
			return nil, err
		}
		bytes := make([]byte, uploadFile.Size)
		_, err = file.Read(bytes)
		if err != nil {
			return nil, err
		}
		file.Close()
		fileList = append(fileList, &model.UploadFileInfo{
			Data: bytes,
			Size: uploadFile.Size,
			Name: uploadFile.Filename,
		})
	}
	result, err := service.File().BatchUploadFile(ctx, fileList, systemConfig)
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

// CkeditorUploadFile Ckeditor图像中的图片上传
func (c *cFile) CkeditorUploadFile(ctx context.Context, req *v1.CkeditorUploadFileReq) (res *v1.CkeditorUploadFileRes, err error) {
	result, err := service.File().CkeditorUploadFile(ctx)
	if err != nil {
		return nil, err
	}
	res = (*v1.CkeditorUploadFileRes)(&result)
	return
}

// CkeditorUploadCopyFile Ckeditor复制的图片上传
func (c *cFile) CkeditorUploadCopyFile(ctx context.Context, req *v1.CkeditorUploadCopyFileReq) (res *v1.CkeditorUploadCopyFileRes, err error) {
	result, err := service.File().CkeditorUploadCopyFile(ctx)
	if err != nil {
		return nil, err
	}
	res = (*v1.CkeditorUploadCopyFileRes)(&result)
	return
}

// CkeditorUploadToolFile 工具栏的文件上传
func (c *cFile) CkeditorUploadToolFile(ctx context.Context, req *v1.CkeditorUploadToolFileReq) (res *v1.CkeditorUploadToolFileRes, err error) {
	result, err := service.File().CkeditorUploadToolFile(ctx)
	if err != nil {
		return nil, err
	}
	res = (*v1.CkeditorUploadToolFileRes)(&result)
	return
}
