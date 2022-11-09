package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Storage struct {
	Uid            string      `json:"uid" q:"EQ" dc:"唯一uid"`             //唯一uid
	AdminUid       string      `json:"adminUid" q:"EQ" dc:"管理员uid"`       //管理员uid
	StorageSize    int64       `json:"storageSize" q:"EQ" dc:"网盘容量大小"`    //网盘容量大小
	Status         int8        `json:"status" q:"EQ" dc:"状态"`             //状态
	CreateTime     *gtime.Time `json:"createTime" q:"BETWEEN" dc:"创建时间"`  //创建时间
	UpdateTime     *gtime.Time `json:"updateTime" q:"BETWEEN" dc:"更新时间"`  //更新时间
	MaxStorageSize int64       `json:"maxStorageSize" q:"EQ" dc:"最大容量大小"` //最大容量大小
}
