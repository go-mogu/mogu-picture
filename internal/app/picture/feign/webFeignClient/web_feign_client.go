package webFeignClient

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/util"
	"net/http"
)

const (
	ServiceName     = "mogu-web"
	getSystemConfig = "/oauth/getSystemConfig"
)

func GetSystemConfig(ctx context.Context, token string) (result string, err error) {
	resp, err := util.SendRequest(ctx, ServiceName, http.MethodGet, getSystemConfig, token)
	if err != nil {
		return
	}
	result = resp.ReadAllString()
	return
}
