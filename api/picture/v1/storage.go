package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StoragePageListReq 分页查询存储信息表Req
type StoragePageListReq struct {
	g.Meta `path:"/pageList" tags:"Storage" method:"post" summary:"分页查询存储信息表"`
	model.Storage
}

// StoragePageListRes 分页查询存储信息表Res
type StoragePageListRes struct {
	Total int               `json:"total" dc:"行数"`
	Rows  []*entity.Storage `json:"rows" dc:"存储信息表数组"`
}

// StorageListReq 列表查询存储信息表Req
type StorageListReq struct {
	g.Meta `path:"/list" tags:"Storage" method:"post" summary:"列表查询存储信息表"`
	entity.Storage
}

// StorageListRes 列表查询Res
type StorageListRes struct {
	Rows []*entity.Storage `json:"rows" dc:"存储信息表数组"`
}

// StorageGetReq 查询存储信息表详情Req
type StorageGetReq struct {
	g.Meta `path:"/:id" tags:"Storage" method:"get" summary:"列表查询存储信息表"`
	Uid    string `json:"uid" dc:"唯一uid"`
}

// StorageGetRes 查询存储信息表详情Res
type StorageGetRes struct {
	*entity.Storage
}

// StorageAddReq 添加存储信息表Req
type StorageAddReq struct {
	g.Meta         `path:"/" tags:"Storage" method:"post" summary:"添加存储信息表"`
	Uid            string      `json:"uid"   dc:"唯一uid"`                                 //唯一uid
	AdminUid       string      `json:"adminUid" v:"required#管理员uid不能为空"  dc:"管理员uid"`    //管理员uid
	StorageSize    int64       `json:"storageSize" v:"required#网盘容量大小不能为空"  dc:"网盘容量大小"` //网盘容量大小
	Status         int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`              //状态
	CreateTime     *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`      //创建时间
	UpdateTime     *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`      //更新时间
	MaxStorageSize int64       `json:"maxStorageSize"   dc:"最大容量大小"`                     //最大容量大小
}

// StorageAddRes 添加存储信息表Req
type StorageAddRes struct {
	Msg string `json:"msg" dc:"添加提示"`
}

// StorageEditReq 编辑存储信息表Req
type StorageEditReq struct {
	g.Meta         `path:"/" tags:"Storage" method:"put" summary:"编辑存储信息表"`
	Uid            string      `json:"uid" v:"required#唯一uid不能为空"  dc:"唯一uid"`           //主键
	VersionNumber  string      `json:"versionNumber" v:"required#未识别到版本号" dc:"版本号"`      //版本号
	AdminUid       string      `json:"adminUid" v:"required#管理员uid不能为空"  dc:"管理员uid"`    //管理员uid
	StorageSize    int64       `json:"storageSize" v:"required#网盘容量大小不能为空"  dc:"网盘容量大小"` //网盘容量大小
	Status         int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`              //状态
	CreateTime     *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`      //创建时间
	UpdateTime     *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`      //更新时间
	MaxStorageSize int64       `json:"maxStorageSize"   dc:"最大容量大小"`                     //最大容量大小
}

// StorageEditRes 编辑存储信息表Res
type StorageEditRes struct {
	Msg string `json:"msg" dc:"编辑提示"`
}

// StorageEditStateReq 编辑存储信息表状态Req
type StorageEditStateReq struct {
	g.Meta `path:"/state" tags:"Storage" method:"put" summary:"批量编辑存储信息表状态"`
	Ids    []string `json:"ids" v:"required#字典主键不能为空"  dc:"id集合"`   //主键
	State  int8     `json:"state" v:"required#字典状态不能为空"  dc:"字典状态"` //状态
}

// StorageDelReq 删除存储信息表Req
type StorageDelReq struct {
	g.Meta `path:"/" tags:"Storage" method:"delete" summary:"删除存储信息表"`
	Ids    []string `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}

// StorageDelRes 删除存储信息表Res
type StorageDelRes struct {
	Msg string `json:"msg" dc:"删除提示"`
}
