package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type File struct {
	Uid             string      `json:"uid" q:"EQ" dc:"唯一uid"`               //唯一uid
	FileOldName     string      `json:"fileOldName" q:"LIKE" dc:"旧文件名"`      //旧文件名
	PicName         string      `json:"picName" q:"LIKE" dc:"文件名"`           //文件名
	PicUrl          string      `json:"picUrl" q:"EQ" dc:"文件地址"`             //文件地址
	PicExpandedName string      `json:"picExpandedName" q:"LIKE" dc:"文件扩展名"` //文件扩展名
	FileSize        int64       `json:"fileSize" q:"EQ" dc:"文件大小"`           //文件大小
	FileSortUid     string      `json:"fileSortUid" q:"EQ" dc:"文件分类uid"`     //文件分类uid
	AdminUid        string      `json:"adminUid" q:"EQ" dc:"管理员uid"`         //管理员uid
	UserUid         string      `json:"userUid" q:"EQ" dc:"用户uid"`           //用户uid
	Status          int8        `json:"status" q:"EQ" dc:"状态"`               //状态
	CreateTime      *gtime.Time `json:"createTime" q:"BETWEEN" dc:"创建时间"`    //创建时间
	UpdateTime      *gtime.Time `json:"updateTime" q:"BETWEEN" dc:"更新时间"`    //更新时间
	QiNiuUrl        string      `json:"qiNiuUrl" q:"EQ" dc:"七牛云地址"`          //七牛云地址
	MinioUrl        string      `json:"minioUrl" q:"EQ" dc:"Minio文件URL"`     //Minio文件URL
	FileMd5         string      `json:"fileMd5" q:"EQ" dc:"文件md5值"`          //文件md5值
}
