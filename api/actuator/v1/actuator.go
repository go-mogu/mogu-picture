package v1

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/gogf/gf/v2/frame/g"
)

// HealthReq 健康检测Req
type HealthReq struct {
	g.Meta `path:"/actuator/health" tags:"actuator" method:"get" summary:"健康检测接口"`
	model.File
}

// HealthRes 健康检测Res
type HealthRes struct {
	Status int
}
