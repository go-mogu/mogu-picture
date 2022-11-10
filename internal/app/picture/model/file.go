package model

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/model"
)

type File struct {
	model.PageInfo
	entity.File
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
