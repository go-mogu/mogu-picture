package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type FileSort struct {
	Uid         string      `json:"uid" q:"EQ" dc:"唯一uid"`            //唯一uid
	ProjectName string      `json:"projectName" q:"LIKE" dc:"项目名"`    //项目名
	SortName    string      `json:"sortName" q:"LIKE" dc:"分类名"`       //分类名
	Url         string      `json:"url" q:"EQ" dc:"分类路径"`             //分类路径
	Status      int8        `json:"status" q:"EQ" dc:"状态"`            //状态
	CreateTime  *gtime.Time `json:"createTime" q:"BETWEEN" dc:"创建时间"` //创建时间
	UpdateTime  *gtime.Time `json:"updateTime" q:"BETWEEN" dc:"更新时间"` //更新时间
}
