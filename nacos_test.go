package main

import (
	"fmt"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	ctx := gctx.New()
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "test", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            "nacos",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "120.48.7.239",
			ContextPath: "/nacos",
			Port:        30848,
			Scheme:      "http",
			GrpcPort:    32102,
		},
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "mogu-picture-test.yaml",
		Group:  "test",
		Type:   "YAML"})
	if err != nil {
		panic(err)
	}

	//success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        9602,
	//	ServiceName: "mogu-picture",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   false,
	//	Metadata:    map[string]string{"appName": "mogu-picture"},
	//	GroupName:   "test", // 默认值DEFAULT_GROUP
	//})
	err = namingClient.Unsubscribe(&vo.SubscribeParam{
		ServiceName: "mogu-picture",
		GroupName:   "test", // 默认值DEFAULT_GROUP
		SubscribeCallback: func(services []model.Instance, err error) {
			log.Printf("\n\n callback return services:%s \n\n", gjson.New(services).MustToJsonString())
		},
	})
	utils.ErrIsNil(ctx, err)
	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "169.254.48.216",
		Port:        52623,
		ServiceName: "mogu-picture",
		GroupName:   "test",
	})
	if !success {
		panic("注册失败")
	}
	data, err := g.Cfg().Data(ctx)
	if err != nil {
		return
	}
	yaml, err := gjson.LoadYaml(content)
	if err != nil {
		return
	}
	m := yaml.Map()
	for k, v := range m {
		data[k] = v
	}
	get, err := g.Cfg().Get(ctx, "app")
	get, err = g.Cfg().Get(ctx, "logger")
	if err != nil {
		return
	}
	fmt.Println(get)
	if err != nil {
		panic(err)
	}

}
