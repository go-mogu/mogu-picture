package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type FileSort struct {
	g.Meta      `orm:"table:t_file_sort, do:true"`
	Uid         interface{} //唯一uid
	ProjectName interface{} //项目名
	SortName    interface{} //分类名
	Url         interface{} //分类路径
	Status      interface{} //状态
	CreateTime  *gtime.Time //创建时间
	UpdateTime  *gtime.Time //更新时间
}
