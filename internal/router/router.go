package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"mogu-picture/internal/app/picture/router"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		router.InitFile(group)
	})

}
