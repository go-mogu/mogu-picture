package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Storage struct {
	g.Meta         `orm:"table:t_storage, do:true"`
	Uid            interface{} //唯一uid
	AdminUid       interface{} //管理员uid
	StorageSize    interface{} //网盘容量大小
	Status         interface{} //状态
	CreateTime     *gtime.Time //创建时间
	UpdateTime     *gtime.Time //更新时间
	MaxStorageSize interface{} //最大容量大小
}
