package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// GetFileListReq 列表查询网盘文件表Req
type GetFileListReq struct {
	g.Meta `path:"/getFileList" tags:"NetworkDisk" method:"post" summary:"列表查询网盘文件表" dc:"列表查询网盘文件表"`
	model.NetworkDisk
}

// GetFileListRes 列表查询Res
type GetFileListRes []*model.NetworkDisk

// CreateFileReq 添加网盘文件表Req
type CreateFileReq struct {
	g.Meta `path:"/createFile" tags:"NetworkDisk" method:"post" summary:"添加网盘文件表" dc:"添加网盘文件表"`
	*entity.NetworkDisk
}

// CreateFileRes 添加网盘文件表Req
type CreateFileRes string

// NetworkDiskEditReq 重命名文件 Req
type NetworkDiskEditReq struct {
	g.Meta `path:"/edit" tags:"NetworkDisk" method:"post" summary:"重命名文件" dc:"重命名文件"`
	*model.NetworkDisk
}

// NetworkDiskEditRes 编辑网盘文件表Res
type NetworkDiskEditRes string

// BatchDeleteFileReq 批量删除文件 Req
type BatchDeleteFileReq struct {
	g.Meta `path:"/batchDeleteFile" tags:"NetworkDisk" method:"post" summary:"批量删除文件" dc:"批量删除文件"`
}

// BatchDeleteFileRes 删除网盘文件表Res
type BatchDeleteFileRes string

// DeleteFileReq 删除文件 Req
type DeleteFileReq struct {
	g.Meta `path:"/deleteFile" tags:"NetworkDisk" method:"post" summary:"删除文件" dc:"删除文件"`
	model.NetworkDisk
}

// DeleteFileRes 删除网盘文件表Res
type DeleteFileRes string

// UnzipFileReq 解压文件 Req
type UnzipFileReq struct {
	g.Meta `path:"/unzipFile" tags:"NetworkDisk" method:"post" summary:"解压文件" dc:"解压文件"`
	model.NetworkDisk
}

// UnzipFileRes 解压文件 Res
type UnzipFileRes string

// MoveFileReq 文件移动 Req
type MoveFileReq struct {
	g.Meta `path:"/moveFile" tags:"NetworkDisk" method:"post" summary:"文件移动" dc:"文件移动"`
	model.NetworkDisk
}

// MoveFileRes 文件移动 Res
type MoveFileRes string

// BatchMoveFileReq 批量文件移动 Req
type BatchMoveFileReq struct {
	g.Meta `path:"/batchMoveFile" tags:"NetworkDisk" method:"post" summary:"批量文件移动" dc:"批量文件移动"`
	model.NetworkDisk
}

// BatchMoveFileRes 批量文件移动 Res
type BatchMoveFileRes string

// SelectFileByFileTypeReq 通过文件类型查询文件 Req
type SelectFileByFileTypeReq struct {
	g.Meta `path:"/selectFileByFileType" tags:"NetworkDisk" method:"get" summary:"通过文件类型查询文件" dc:"通过文件类型查询文件"`
	model.NetworkDisk
}

// SelectFileByFileTypeRes 通过文件类型查询文件 Res
type SelectFileByFileTypeRes model.NetworkDiskList

// GetFileTreeReq 获取文件树 Req
type GetFileTreeReq struct {
	g.Meta `path:"/getFileTree" tags:"NetworkDisk" method:"post" summary:"获取文件树" dc:"获取文件树"`
}

// GetFileTreeRes 获取文件树 Res
type GetFileTreeRes baseModel.TreeNode
