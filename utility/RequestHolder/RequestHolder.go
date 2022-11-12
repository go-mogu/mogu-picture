package RequestHolder

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// GetRequest 获取request
func GetRequest(ctx context.Context) *ghttp.Request {
	request := g.RequestFromCtx(ctx)
	return request
}

// GetAdminUid 获取AdminUid
func GetAdminUid(ctx context.Context) string {
	return GetRequest(ctx).Get(SysConf.ADMIN_UID).String()
}

// CheckLogin 检查当前用户是否登录【未登录操作将抛出QueryException异常】
func CheckLogin(ctx context.Context) (adminUid string) {
	adminUid = GetAdminUid(ctx)
	if g.IsEmpty(adminUid) {
		g.Log().Error(ctx, MessageConf.INVALID_TOKEN)
		panic(MessageConf.INVALID_TOKEN)
	}
	return
}
