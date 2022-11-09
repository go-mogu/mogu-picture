package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type NetworkDisk struct {
	g.Meta      `orm:"table:t_network_disk, do:true"`
	Uid         interface{} //唯一uid
	AdminUid    interface{} //管理员uid
	ExtendName  interface{} //扩展名
	FileName    interface{} //文件名
	FilePath    interface{} //文件路径
	FileSize    interface{} //文件大小
	IsDir       interface{} //是否目录
	Status      interface{} //状态
	CreateTime  *gtime.Time //创建时间
	UpdateTime  *gtime.Time //更新时间
	LocalUrl    interface{} //本地文件URL
	QiNiuUrl    interface{} //七牛云URL
	FileOldName interface{} //上传前文件名
	MinioUrl    interface{} //Minio文件URL
}
