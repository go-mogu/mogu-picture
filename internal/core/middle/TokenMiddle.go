package middle

import (
	"github.com/go-mogu/mogu-picture/internal/consts/RedisConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	"github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func TokenMiddle(r *ghttp.Request) {
	//得到请求头信息authorization信息
	authHeader := ""
	authorization := r.GetHeader(SysConf.AUTHORIZATION)
	if authorization != "" {
		authHeader = authorization
	} else if token := r.Get(SysConf.TOKEN); token != nil {
		authHeader = token.String()
	}
	if authHeader != "" && gstr.HasPrefix(authHeader, SysConf.BEARER) {
		onlineAdmin, err := g.Redis().Do(r.Context(), "GET", RedisConf.LOGIN_TOKEN_KEY+RedisConf.SEGMENTATION+authHeader)
		utils.ErrIsNil(r.Context(), err)
		if onlineAdmin != nil {
			admin := new(model.OnlineAdmin)
			err = onlineAdmin.Struct(&admin)
			utils.ErrIsNil(r.Context(), err)
			r.SetParam(SysConf.ADMIN_UID, admin.AdminUid)
			r.SetParam(SysConf.NAME, admin.UserName)
			r.SetParam(SysConf.USER_UID, admin.AdminUid)
			r.SetParam(SysConf.TOKEN, authHeader)
		}
	}
	r.Middleware.Next()

}
