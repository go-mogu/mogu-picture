package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NetworkDiskPageListReq 分页查询网盘文件表Req
type NetworkDiskPageListReq struct {
	g.Meta `path:"/pageList" tags:"NetworkDisk" method:"post" summary:"分页查询网盘文件表"`
	model.NetworkDisk
}

// NetworkDiskPageListRes 分页查询网盘文件表Res
type NetworkDiskPageListRes struct {
	Total int                   `json:"total" dc:"行数"`
	Rows  []*entity.NetworkDisk `json:"rows" dc:"网盘文件表数组"`
}

// NetworkDiskListReq 列表查询网盘文件表Req
type NetworkDiskListReq struct {
	g.Meta `path:"/list" tags:"NetworkDisk" method:"post" summary:"列表查询网盘文件表"`
	entity.NetworkDisk
}

// NetworkDiskListRes 列表查询Res
type NetworkDiskListRes struct {
	Rows []*entity.NetworkDisk `json:"rows" dc:"网盘文件表数组"`
}

// NetworkDiskGetReq 查询网盘文件表详情Req
type NetworkDiskGetReq struct {
	g.Meta `path:"/:id" tags:"NetworkDisk" method:"get" summary:"列表查询网盘文件表"`
	Uid    string `json:"uid" dc:"唯一uid"`
}

// NetworkDiskGetRes 查询网盘文件表详情Res
type NetworkDiskGetRes struct {
	*entity.NetworkDisk
}

// NetworkDiskAddReq 添加网盘文件表Req
type NetworkDiskAddReq struct {
	g.Meta      `path:"/" tags:"NetworkDisk" method:"post" summary:"添加网盘文件表"`
	Uid         string      `json:"uid"   dc:"唯一uid"`                                 //唯一uid
	AdminUid    string      `json:"adminUid" v:"required#管理员uid不能为空"  dc:"管理员uid"`    //管理员uid
	ExtendName  string      `json:"extendName" v:"required#扩展名不能为空"  dc:"扩展名"`        //扩展名
	FileName    string      `json:"fileName" v:"required#文件名不能为空"  dc:"文件名"`          //文件名
	FilePath    string      `json:"filePath"   dc:"文件路径"`                             //文件路径
	FileSize    int64       `json:"fileSize" v:"required#文件大小不能为空"  dc:"文件大小"`        //文件大小
	IsDir       int         `json:"isDir" v:"required#是否目录不能为空"  dc:"是否目录"`           //是否目录
	Status      int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`              //状态
	CreateTime  *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`      //创建时间
	UpdateTime  *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`      //更新时间
	LocalUrl    string      `json:"localUrl"   dc:"本地文件URL"`                          //本地文件URL
	QiNiuUrl    string      `json:"qiNiuUrl"   dc:"七牛云URL"`                           //七牛云URL
	FileOldName string      `json:"fileOldName" v:"required#上传前文件名不能为空"  dc:"上传前文件名"` //上传前文件名
	MinioUrl    string      `json:"minioUrl"   dc:"Minio文件URL"`                       //Minio文件URL
}

// NetworkDiskAddRes 添加网盘文件表Req
type NetworkDiskAddRes struct {
	Msg string `json:"msg" dc:"添加提示"`
}

// NetworkDiskEditReq 编辑网盘文件表Req
type NetworkDiskEditReq struct {
	g.Meta        `path:"/" tags:"NetworkDisk" method:"put" summary:"编辑网盘文件表"`
	Uid           string      `json:"uid" v:"required#唯一uid不能为空"  dc:"唯一uid"`           //主键
	VersionNumber string      `json:"versionNumber" v:"required#未识别到版本号" dc:"版本号"`      //版本号
	AdminUid      string      `json:"adminUid" v:"required#管理员uid不能为空"  dc:"管理员uid"`    //管理员uid
	ExtendName    string      `json:"extendName" v:"required#扩展名不能为空"  dc:"扩展名"`        //扩展名
	FileName      string      `json:"fileName" v:"required#文件名不能为空"  dc:"文件名"`          //文件名
	FilePath      string      `json:"filePath"   dc:"文件路径"`                             //文件路径
	FileSize      int64       `json:"fileSize" v:"required#文件大小不能为空"  dc:"文件大小"`        //文件大小
	IsDir         int         `json:"isDir" v:"required#是否目录不能为空"  dc:"是否目录"`           //是否目录
	Status        int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`              //状态
	CreateTime    *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`      //创建时间
	UpdateTime    *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`      //更新时间
	LocalUrl      string      `json:"localUrl"   dc:"本地文件URL"`                          //本地文件URL
	QiNiuUrl      string      `json:"qiNiuUrl"   dc:"七牛云URL"`                           //七牛云URL
	FileOldName   string      `json:"fileOldName" v:"required#上传前文件名不能为空"  dc:"上传前文件名"` //上传前文件名
	MinioUrl      string      `json:"minioUrl"   dc:"Minio文件URL"`                       //Minio文件URL
}

// NetworkDiskEditRes 编辑网盘文件表Res
type NetworkDiskEditRes struct {
	Msg string `json:"msg" dc:"编辑提示"`
}

// NetworkDiskEditStateReq 编辑网盘文件表状态Req
type NetworkDiskEditStateReq struct {
	g.Meta `path:"/state" tags:"NetworkDisk" method:"put" summary:"批量编辑网盘文件表状态"`
	Ids    []string `json:"ids" v:"required#字典主键不能为空"  dc:"id集合"`   //主键
	State  int8     `json:"state" v:"required#字典状态不能为空"  dc:"字典状态"` //状态
}

// NetworkDiskDelReq 删除网盘文件表Req
type NetworkDiskDelReq struct {
	g.Meta `path:"/" tags:"NetworkDisk" method:"delete" summary:"删除网盘文件表"`
	Ids    []string `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}

// NetworkDiskDelRes 删除网盘文件表Res
type NetworkDiskDelRes struct {
	Msg string `json:"msg" dc:"删除提示"`
}
