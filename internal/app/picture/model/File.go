package model

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/model"
)

type File struct {
	entity.File
	PictureUrl string `json:"pictureUrl" q:"-" dc:"当前的url地址"` //当前的url地址
}

type FileVO struct {
	model.BaseVO
	AdminUid     string            `json:"adminUid" q:"EQ" dc:"管理员uid"`   //管理员uid
	UserUid      string            `json:"userUid" q:"EQ" dc:"用户uid"`     //用户uid
	ProjectName  string            `json:"projectName" q:"LIKE" dc:"项目名"` //项目名
	SortName     string            `json:"sortName" q:"EQ" dc:"模块名"`      //模块名
	UrlList      []string          `json:"urlList" q:"-" dc:"图片Url集合"`    //图片Url集合
	SystemConfig map[string]string `json:"systemConfig" q:"-" dc:"系统配置"`  //系统配置
	Token        string            `json:"token" dc:"上传图片时携带的token令牌"`    //token
}

type UploadFileParam struct {
	Ctx          context.Context    `json:"-" dc:"上下文"`             //上下文
	SystemConfig model.SystemConfig `json:"systemConfig" dc:"系统设置"` //系统设置
	OldName      string             `json:"oldName" dc:"旧文件名"`      //旧文件名
	NewFileName  string             `json:"newFileName" dc:"新文件名"`  //新文件名
	Data         []byte             `json:"data" dc:"上传文件字节数组"`     //上传文件字节数组
	FileSort     *entity.FileSort   `json:"fileSort" dc:"文件分类信息"`   //文件分类信息
	File         *File              `json:"file" dc:"文件信息"`         //文件信息
}

type UploadFileInfo struct {
	Data []byte `json:"data" dc:"上传文件字节数组"` //上传文件字节数组
	Size int64  `json:"size" dc:"上传文件大小"`   //上传文件大小
	Name string `json:"name" dc:"上传文件名称"`   //上传文件名称
}
