package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitStorage(group *ghttp.RouterGroup) {
	group.Group("/storage", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Storage)
	})
}
