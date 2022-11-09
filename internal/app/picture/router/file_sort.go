package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitFileSort(group *ghttp.RouterGroup) {
	group.Group("/fileSort", func(group *ghttp.RouterGroup) {
		group.Bind(controller.FileSort)
	})
}
