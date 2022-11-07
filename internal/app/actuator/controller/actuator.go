package controller

import (
	"context"
	v1 "github.com/go-mogu/mogu-picture/api/actuator/v1"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Actuator = cActuator{}
)

type cActuator struct{}

// Health 健康检测
func (c *cActuator) Health(ctx context.Context, req *v1.HealthReq) (res *v1.HealthRes, err error) {
	if g.Server().Status() == ghttp.ServerStatusRunning {
		res = &v1.HealthRes{Status: ghttp.ServerStatusRunning}
	} else {
		return nil, gerror.NewCode(gcode.CodeNil, "server is already stopped")
	}
	return
}
