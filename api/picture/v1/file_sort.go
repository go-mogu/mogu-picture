package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileSortPageListReq 分页查询文件分类表Req
type FileSortPageListReq struct {
	g.Meta `path:"/pageList" tags:"FileSort" method:"post" summary:"分页查询文件分类表"`
	model.FileSort
}

// FileSortPageListRes 分页查询文件分类表Res
type FileSortPageListRes struct {
	Total int                `json:"total" dc:"行数"`
	Rows  []*entity.FileSort `json:"rows" dc:"文件分类表数组"`
}

// FileSortListReq 列表查询文件分类表Req
type FileSortListReq struct {
	g.Meta `path:"/list" tags:"FileSort" method:"post" summary:"列表查询文件分类表"`
	entity.FileSort
}

// FileSortListRes 列表查询Res
type FileSortListRes struct {
	Rows []*entity.FileSort `json:"rows" dc:"文件分类表数组"`
}

// FileSortGetReq 查询文件分类表详情Req
type FileSortGetReq struct {
	g.Meta `path:"/:id" tags:"FileSort" method:"get" summary:"列表查询文件分类表"`
	Uid    string `json:"uid" dc:"唯一uid"`
}

// FileSortGetRes 查询文件分类表详情Res
type FileSortGetRes struct {
	*entity.FileSort
}

// FileSortAddReq 添加文件分类表Req
type FileSortAddReq struct {
	g.Meta      `path:"/" tags:"FileSort" method:"post" summary:"添加文件分类表"`
	Uid         string      `json:"uid"   dc:"唯一uid"`                            //唯一uid
	ProjectName string      `json:"projectName" v:"required#项目名不能为空"  dc:"项目名"`  //项目名
	SortName    string      `json:"sortName" v:"required#分类名不能为空"  dc:"分类名"`     //分类名
	Url         string      `json:"url"   dc:"分类路径"`                             //分类路径
	Status      int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`         //状态
	CreateTime  *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"` //创建时间
	UpdateTime  *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"` //更新时间
}

// FileSortAddRes 添加文件分类表Req
type FileSortAddRes struct {
	Msg string `json:"msg" dc:"添加提示"`
}

// FileSortEditReq 编辑文件分类表Req
type FileSortEditReq struct {
	g.Meta        `path:"/" tags:"FileSort" method:"put" summary:"编辑文件分类表"`
	Uid           string      `json:"uid" v:"required#唯一uid不能为空"  dc:"唯一uid"`      //主键
	VersionNumber string      `json:"versionNumber" v:"required#未识别到版本号" dc:"版本号"` //版本号
	ProjectName   string      `json:"projectName" v:"required#项目名不能为空"  dc:"项目名"`  //项目名
	SortName      string      `json:"sortName" v:"required#分类名不能为空"  dc:"分类名"`     //分类名
	Url           string      `json:"url"   dc:"分类路径"`                             //分类路径
	Status        int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`         //状态
	CreateTime    *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"` //创建时间
	UpdateTime    *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"` //更新时间
}

// FileSortEditRes 编辑文件分类表Res
type FileSortEditRes struct {
	Msg string `json:"msg" dc:"编辑提示"`
}

// FileSortEditStateReq 编辑文件分类表状态Req
type FileSortEditStateReq struct {
	g.Meta `path:"/state" tags:"FileSort" method:"put" summary:"批量编辑文件分类表状态"`
	Ids    []string `json:"ids" v:"required#字典主键不能为空"  dc:"id集合"`   //主键
	State  int8     `json:"state" v:"required#字典状态不能为空"  dc:"字典状态"` //状态
}

// FileSortDelReq 删除文件分类表Req
type FileSortDelReq struct {
	g.Meta `path:"/" tags:"FileSort" method:"delete" summary:"删除文件分类表"`
	Ids    []string `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}

// FileSortDelRes 删除文件分类表Res
type FileSortDelRes struct {
	Msg string `json:"msg" dc:"删除提示"`
}
