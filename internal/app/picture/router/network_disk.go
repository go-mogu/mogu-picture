package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitNetworkDisk(group *ghttp.RouterGroup) {
	group.Group("/networkDisk", func(group *ghttp.RouterGroup) {
		group.Bind(controller.NetworkDisk)
	})
}
