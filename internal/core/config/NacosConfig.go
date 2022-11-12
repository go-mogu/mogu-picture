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
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	nacosModel "github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"time"
)

const (
	appNameKey = "app.name"
	nacosKey   = "nacos"
)

var (
	nacosConfig model.NacosProperties
	//服务发现客户端
	namingClient naming_client.INamingClient
	//配置客户端
	configClient config_client.IConfigClient
)

func init() {
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, nacosKey).Scan(&nacosConfig)
	utils.ErrIsNil(ctx, err)
	// 创建动态配置客户端
	configClient, err = clients.NewConfigClient(
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
}

func getConfigParam(ctx context.Context, config vo.ConfigParam) vo.ConfigParam {
	appName := g.Cfg().MustGet(ctx, appNameKey).String()

	config.DataId = fmt.Sprintf("%s-%s.%s", appName, config.Group, gstr.ToLower(config.Type))
	config.Type = gstr.ToUpper(config.Type)
	return config
}

func RegisterInstance(ctx context.Context, s *ghttp.Server) (err error) {
	// 创建服务发现客户端
	namingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &nacosConfig.Client,
			ServerConfigs: nacosConfig.Server,
		},
	)
	if err != nil {
		return err
	}
	if nacosConfig.Discovery.Port == 0 {
		nacosConfig.Discovery.Port = uint64(s.GetListenedPort())
	}
	if nacosConfig.Discovery.ServiceName == "" {
		nacosConfig.Discovery.ServiceName = s.GetName()
	}
	if nacosConfig.Discovery.Ip == "" {
		ip, err := utils.GetLocalIP()
		if err != nil {
			return err
		}
		nacosConfig.Discovery.Ip = ip
	}
	success, err := namingClient.RegisterInstance(nacosConfig.Discovery)
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
	services, err := GetService(g.Server().GetName())
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
	}

}

func GetService(serviceName string) (services nacosModel.Service, err error) {
	services, err = namingClient.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		GroupName:   nacosConfig.Config.Group,
	})
	return
}

func GetInstances(serviceName string) ([]nacosModel.Instance, error) {
	// SelectInstances 只返回满足这些条件的实例列表：healthy=${HealthyOnly},enable=true 和weight>0
	instances, err := namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   nacosConfig.Discovery.GroupName, // 默认值DEFAULT_GROUP
		HealthyOnly: true,
	})
	return instances, err
}
