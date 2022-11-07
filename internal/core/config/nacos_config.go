package config

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"time"
)

var (
	nacosConfig model.NacosProperties
)

func init() {
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "nacos").Scan(&nacosConfig)
	utils.ErrIsNil(ctx, err)

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &nacosConfig.Client,
			ServerConfigs: nacosConfig.Server,
		},
	)
	utils.ErrIsNil(ctx, err)
	content, err := configClient.GetConfig(getConfigParam(ctx, nacosConfig.Config))
	utils.ErrIsNil(ctx, err)
	data := g.Cfg().MustData(ctx)
	yaml, err := gjson.LoadYaml(content)
	utils.ErrIsNil(ctx, err)
	m := yaml.Map()
	for k, v := range m {
		data[k] = v
	}

	//success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        9602,
	//	ServiceName: "mogu-picture",
	//	Ephemeral:   true,
	//	GroupName:   "test",
	//})
	//if !success {
	//	panic("注册失败")
	//}

}

func getConfigParam(ctx context.Context, config vo.ConfigParam) vo.ConfigParam {
	appName := g.Cfg().MustGet(ctx, "app.name").String()

	config.DataId = fmt.Sprintf("%s-%s.%s", appName, config.Group, gstr.ToLower(config.Type))
	config.Type = gstr.ToUpper(config.Type)
	return config
}

func RegisterInstance(ctx context.Context, s *ghttp.Server) (err error) {
	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &nacosConfig.Client,
			ServerConfigs: nacosConfig.Server,
		},
	)
	if err != nil {
		return err
	}

	instanceParam := vo.RegisterInstanceParam{
		Port:        uint64(s.GetListenedPort()),
		ServiceName: s.GetName(),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{},
		GroupName:   "test", // 默认值DEFAULT_GROUP
	}
	if instanceParam.Ip == "" {
		ip, err := utils.GetLocalIP()
		if err != nil {
			return err
		}
		instanceParam.Ip = ip
	}
	success, err := namingClient.RegisterInstance(instanceParam)
	if err != nil {
		return
	}
	if !success {
		return gerror.New("register instance failed!")
	}
	gtimer.SetInterval(ctx, 10*time.Second, func(ctx context.Context) {
		actuator(ctx, namingClient)
	})
	return
}

func actuator(ctx context.Context, namingClient naming_client.INamingClient) {
	services, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: g.Server().GetName(),
		GroupName:   nacosConfig.Config.Group,
	})
	utils.ErrIsNil(ctx, err)
	ip, err := utils.GetLocalIP()
	var port = uint64(0)
	utils.ErrIsNil(ctx, err)
	hosts := services.Hosts
	for _, instance := range hosts {
		utils.ErrIsNil(ctx, err)
		if instance.Ip == ip {
			port = instance.Port
			break
		}
	}
	if port != 0 {
		r, err := g.Client().Get(ctx, fmt.Sprintf("http://%s:%d/actuator/health", ip, port))
		utils.ErrIsNil(ctx, err)
		defer func(r *gclient.Response) {
			err = r.Close()
			utils.ErrIsNil(ctx, err)
		}(r)
	} else {
		panic(gerror.New("register instance failed!"))
	}

}
