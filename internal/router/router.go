package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/router"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		router.InitFile(group)
	})

}
