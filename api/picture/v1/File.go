package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/gogf/gf/v2/frame/g"
)

// CropperPictureReq 截图上传 Req
type CropperPictureReq struct {
	g.Meta `path:"/cropperPicture" tags:"File" method:"post" summary:"截图上传"`
}

// CropperPictureRes 截图上传 Res
type CropperPictureRes []map[string]interface{}

// GetPictureReq 查询文件表详情Req
type GetPictureReq struct {
	g.Meta  `path:"/getPicture" tags:"File" method:"get" summary:"通过fileIds获取图片信息接口"`
	FileIds string `json:"fileIds" v:"required#图片UID不能为空" dc:"文件ids"`
	Code    string `json:"code" d:"," dc:"切割符"`
}

// GetPictureRes 查询文件表详情Res
type GetPictureRes []map[string]interface{}

// UploadPicsReq 多文件上传 Req
type UploadPicsReq struct {
	g.Meta `path:"/pictures" tags:"File" method:"post" summary:"多文件上传"`
}

// UploadPicsRes 多文件上传 Res
type UploadPicsRes []*model.File

// UploadPicsByUrlReq 通过URL上传图片 Req
type UploadPicsByUrlReq struct {
	g.Meta `path:"/uploadPicsByUrl" tags:"File" method:"post" summary:"通过URL上传图片"`
	model.FileVO
}

// UploadPicsByUrlRes Ckeditor图像中的图片上传 Res
type UploadPicsByUrlRes []*model.File

// CkeditorUploadFileReq 通过URL上传图片 Req
type CkeditorUploadFileReq struct {
	g.Meta `path:"/ckeditorUploadFile" tags:"File" method:"post" summary:"Ckeditor图像中的图片上传"`
}

// CkeditorUploadFileRes Ckeditor图像中的图片上传 Res
type CkeditorUploadFileRes map[string]interface{}

// CkeditorUploadCopyFileReq Ckeditor复制的图片上传 Req
type CkeditorUploadCopyFileReq struct {
	g.Meta `path:"/ckeditorUploadCopyFile" tags:"File" method:"post" summary:"Ckeditor复制的图片上传"`
}

// CkeditorUploadCopyFileRes Ckeditor复制的图片上传 Res
type CkeditorUploadCopyFileRes map[string]interface{}

// CkeditorUploadToolFileReq 工具栏的文件上传 Req
type CkeditorUploadToolFileReq struct {
	g.Meta `path:"/ckeditorUploadToolFile" tags:"File" method:"post" summary:"工具栏的文件上传"`
}

// CkeditorUploadToolFileRes 工具栏的文件上传 Res
type CkeditorUploadToolFileRes map[string]interface{}
