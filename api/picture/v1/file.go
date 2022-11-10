package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FilePageListReq 分页查询文件表Req
type FilePageListReq struct {
	g.Meta `path:"/pageList" tags:"File" method:"post" summary:"分页查询文件表"`
	model.File
}

// FilePageListRes 分页查询文件表Res
type FilePageListRes struct {
	Total int            `json:"total" dc:"行数"`
	Rows  []*entity.File `json:"rows" dc:"文件表数组"`
}

// FileListReq 列表查询文件表Req
type FileListReq struct {
	g.Meta `path:"/list" tags:"File" method:"post" summary:"列表查询文件表"`
	entity.File
}

// FileListRes 列表查询Res
type FileListRes struct {
	Rows []*entity.File `json:"rows" dc:"文件表数组"`
}

// FileAddReq 添加文件表Req
type FileAddReq struct {
	g.Meta          `path:"/" tags:"File" method:"post" summary:"添加文件表"`
	Uid             string      `json:"uid"   dc:"唯一uid"`                                   //唯一uid
	FileOldName     string      `json:"fileOldName" v:"required#旧文件名不能为空"  dc:"旧文件名"`       //旧文件名
	PicName         string      `json:"picName" v:"required#文件名不能为空"  dc:"文件名"`             //文件名
	PicUrl          string      `json:"picUrl"   dc:"文件地址"`                                 //文件地址
	PicExpandedName string      `json:"picExpandedName" v:"required#文件扩展名不能为空"  dc:"文件扩展名"` //文件扩展名
	FileSize        int         `json:"fileSize"   dc:"文件大小"`                               //文件大小
	FileSortUid     string      `json:"fileSortUid"   dc:"文件分类uid"`                         //文件分类uid
	AdminUid        string      `json:"adminUid"   dc:"管理员uid"`                             //管理员uid
	UserUid         string      `json:"userUid"   dc:"用户uid"`                               //用户uid
	Status          int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`                //状态
	CreateTime      *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`        //创建时间
	UpdateTime      *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`        //更新时间
	QiNiuUrl        string      `json:"qiNiuUrl"   dc:"七牛云地址"`                              //七牛云地址
	MinioUrl        string      `json:"minioUrl"   dc:"Minio文件URL"`                         //Minio文件URL
}

// FileAddRes 添加文件表Req
type FileAddRes struct {
	Msg string `json:"msg" dc:"添加提示"`
}

// FileEditReq 编辑文件表Req
type FileEditReq struct {
	g.Meta          `path:"/" tags:"File" method:"put" summary:"编辑文件表"`
	Uid             string      `json:"uid" v:"required#唯一uid不能为空"  dc:"唯一uid"`             //主键
	VersionNumber   string      `json:"versionNumber" v:"required#未识别到版本号" dc:"版本号"`        //版本号
	FileOldName     string      `json:"fileOldName" v:"required#旧文件名不能为空"  dc:"旧文件名"`       //旧文件名
	PicName         string      `json:"picName" v:"required#文件名不能为空"  dc:"文件名"`             //文件名
	PicUrl          string      `json:"picUrl"   dc:"文件地址"`                                 //文件地址
	PicExpandedName string      `json:"picExpandedName" v:"required#文件扩展名不能为空"  dc:"文件扩展名"` //文件扩展名
	FileSize        int         `json:"fileSize"   dc:"文件大小"`                               //文件大小
	FileSortUid     string      `json:"fileSortUid"   dc:"文件分类uid"`                         //文件分类uid
	AdminUid        string      `json:"adminUid"   dc:"管理员uid"`                             //管理员uid
	UserUid         string      `json:"userUid"   dc:"用户uid"`                               //用户uid
	Status          int8        `json:"status" v:"required#状态不能为空"  dc:"状态"`                //状态
	CreateTime      *gtime.Time `json:"createTime" v:"required#创建时间不能为空"  dc:"创建时间"`        //创建时间
	UpdateTime      *gtime.Time `json:"updateTime" v:"required#更新时间不能为空"  dc:"更新时间"`        //更新时间
	QiNiuUrl        string      `json:"qiNiuUrl"   dc:"七牛云地址"`                              //七牛云地址
	MinioUrl        string      `json:"minioUrl"   dc:"Minio文件URL"`                         //Minio文件URL
}

// FileEditRes 编辑文件表Res
type FileEditRes struct {
	Msg string `json:"msg" dc:"编辑提示"`
}

// FileEditStateReq 编辑文件表状态Req
type FileEditStateReq struct {
	g.Meta `path:"/state" tags:"File" method:"put" summary:"批量编辑文件表状态"`
	Ids    []string `json:"ids" v:"required#字典主键不能为空"  dc:"id集合"`   //主键
	State  int8     `json:"state" v:"required#字典状态不能为空"  dc:"字典状态"` //状态
}

// FileDelReq 删除文件表Req
type FileDelReq struct {
	g.Meta `path:"/" tags:"File" method:"delete" summary:"删除文件表"`
	Ids    []string `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}

// FileDelRes 删除文件表Res
type FileDelRes struct {
	Msg string `json:"msg" dc:"删除提示"`
}

// CropperPictureReq 截图上传 Req
type CropperPictureReq struct {
	g.Meta `path:"/cropperPicture" tags:"File" method:"post" summary:"截图上传"`
}

// CropperPictureRes 截图上传 Res
type CropperPictureRes []map[string]interface{}

// GetPictureReq 查询文件表详情Req
type GetPictureReq struct {
	g.Meta  `path:"/getPicture" tags:"File" method:"get" summary:"通过fileIds获取图片信息接口"`
	FileIds string `json:"fileIds" dc:"文件ids"`
	Code    string `json:"code" dc:"切割符"`
}

// GetPictureRes 查询文件表详情Res
type GetPictureRes []map[string]interface{}

// UploadPicsReq 多文件上传 Req
type UploadPicsReq struct {
	g.Meta `path:"/pictures" tags:"File" method:"post" summary:"多文件上传"`
}

// UploadPicsRes 多文件上传 Res
type UploadPicsRes []*entity.File

// UploadPicsByUrlReq 通过URL上传图片 Req
type UploadPicsByUrlReq struct {
	g.Meta `path:"/uploadPicsByUrl" tags:"File" method:"post" summary:"通过URL上传图片"`
	model.FileVO
}

// UploadPicsByUrlRes 通过URL上传图片 Res
type UploadPicsByUrlRes []*entity.File
