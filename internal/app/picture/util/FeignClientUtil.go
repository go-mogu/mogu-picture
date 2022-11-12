package util

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	"github.com/go-mogu/mogu-picture/internal/core/lb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

func SendRequest(ctx context.Context, serviceName, method, path string, data ...interface{}) (resp *gclient.Response, err error) {
	url := lb.GetServiceUrl(ctx, serviceName)
	request := g.RequestFromCtx(ctx)
	// 后台携带的token
	c := g.Client()
	token := request.GetHeader(SysConf.AUTHORIZATION)
	c.SetHeader(SysConf.AUTHORIZATION, token)
	resp, err = c.DoRequest(ctx, method, url+path, data)
	return
}
