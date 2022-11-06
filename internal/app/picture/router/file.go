package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"mogu-picture/internal/app/picture/controller"
)

func InitFile(group *ghttp.RouterGroup) {
	group.Group("/file", func(group *ghttp.RouterGroup) {
		group.Bind(controller.File)
	})
}
