package middle

import (
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func TokenMiddle(r *ghttp.Request) {
	//得到请求头信息authorization信息
	authHeader := ""
	authorization := r.GetHeader(consts.Authorization)
	if authorization != "" {
		authHeader = authorization
	} else if token := r.GetParam(consts.Token); token != nil {
		authHeader = token.String()
	}
	if authHeader != "" && gstr.HasPrefix(authHeader, consts.Bearer) {
		onlineAdmin, err := g.Redis().Do(r.Context(), "GET", consts.LoginTokenKey+consts.SymbolColon+authHeader)
		utils.ErrIsNil(r.Context(), err)
		if onlineAdmin != nil {
			admin := new(model.OnlineAdmin)
			err = onlineAdmin.Struct(&admin)
			utils.ErrIsNil(r.Context(), err)
			r.SetParam(consts.AdminUid, admin.AdminUid)
			r.SetParam(consts.Name, admin.UserName)
			r.SetParam(consts.Token, authHeader)
		}
	}
	r.Middleware.Next()

}
