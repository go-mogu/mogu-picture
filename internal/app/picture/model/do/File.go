package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type File struct {
	g.Meta          `orm:"table:t_file, do:true"`
	Uid             interface{} //唯一uid
	FileOldName     interface{} //旧文件名
	PicName         interface{} //文件名
	PicUrl          interface{} //文件地址
	PicExpandedName interface{} //文件扩展名
	FileSize        interface{} //文件大小
	FileSortUid     interface{} //文件分类uid
	AdminUid        interface{} //管理员uid
	UserUid         interface{} //用户uid
	Status          interface{} //状态
	CreateTime      *gtime.Time //创建时间
	UpdateTime      *gtime.Time //更新时间
	QiNiuUrl        interface{} //七牛云地址
	MinioUrl        interface{} //Minio文件URL
	FileMd5         interface{} //文件md5值
}
