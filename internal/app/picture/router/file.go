package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitFile(group *ghttp.RouterGroup) {
	group.Group("/file", func(group *ghttp.RouterGroup) {
		group.Bind(controller.File)
	})
}
