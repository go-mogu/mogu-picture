package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type NetworkDisk struct {
	Uid         string      `json:"uid" q:"EQ" dc:"唯一uid"`            //唯一uid
	AdminUid    string      `json:"adminUid" q:"EQ" dc:"管理员uid"`      //管理员uid
	ExtendName  string      `json:"extendName" q:"LIKE" dc:"扩展名"`     //扩展名
	FileName    string      `json:"fileName" q:"LIKE" dc:"文件名"`       //文件名
	FilePath    string      `json:"filePath" q:"EQ" dc:"文件路径"`        //文件路径
	FileSize    int64       `json:"fileSize" q:"EQ" dc:"文件大小"`        //文件大小
	IsDir       int         `json:"isDir" q:"EQ" dc:"是否目录"`           //是否目录
	Status      int8        `json:"status" q:"EQ" dc:"状态"`            //状态
	CreateTime  *gtime.Time `json:"createTime" q:"BETWEEN" dc:"创建时间"` //创建时间
	UpdateTime  *gtime.Time `json:"updateTime" q:"BETWEEN" dc:"更新时间"` //更新时间
	LocalUrl    string      `json:"localUrl" q:"EQ" dc:"本地文件URL"`     //本地文件URL
	QiNiuUrl    string      `json:"qiNiuUrl" q:"EQ" dc:"七牛云URL"`      //七牛云URL
	FileOldName string      `json:"fileOldName" q:"LIKE" dc:"上传前文件名"` //上传前文件名
	MinioUrl    string      `json:"minioUrl" q:"EQ" dc:"Minio文件URL"`  //Minio文件URL
}
