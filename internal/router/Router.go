package router

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/router"
	"github.com/go-mogu/mogu-picture/internal/core/middle"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Middleware(middle.TokenMiddle)
	router.InitFile(group)
	router.InitNetworkDisk(group)
	router.InitStorage(group)

}
