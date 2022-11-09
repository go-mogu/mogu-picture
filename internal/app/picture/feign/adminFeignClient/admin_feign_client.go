package adminFeignClient

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/util"
	"net/http"
)

const (
	ServiceName     = "mogu-admin"
	getSystemConfig = "/systemConfig/getSystemConfig"
)

func GetSystemConfig(ctx context.Context) (result string, err error) {
	resp, err := util.SendRequest(ctx, ServiceName, http.MethodGet, getSystemConfig)
	if err != nil {
		return
	}
	result = resp.ReadAllString()
	return
}
