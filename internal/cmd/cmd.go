package cmd

import (
	"context"
	actuator "github.com/go-mogu/mogu-picture/internal/app/actuator/controller"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/core/config"
	_ "github.com/go-mogu/mogu-picture/internal/core/config"
	"github.com/go-mogu/mogu-picture/internal/core/middle"
	_ "github.com/go-mogu/mogu-picture/internal/logic"
	"github.com/go-mogu/mogu-picture/internal/router"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "mogu-picture",
		Usage: "mogu-picture",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetName(g.Cfg().MustGet(ctx, "app.name").String())
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middle.MiddlewareHandlerResponse)
				group.Bind(actuator.Actuator)
				group.Middleware(middle.TokenMiddle)
				router.BindController(group)
			})
			enhanceOpenAPIDoc(s)
			err = s.Start()
			utils.ErrIsNil(ctx, err)
			err = config.RegisterInstance(ctx, s)
			if err != nil {
				return err
			}
			g.Wait()
			return err
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
}
