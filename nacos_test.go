package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestClient(t *testing.T) {
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
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
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

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "mogu-picture-test.yaml",
		Group:  "DEFAULT_GROUP"})

	fmt.Println(content)

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        9602,
		ServiceName: "mogu-picture",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   false,
		Metadata:    map[string]string{"appName": "mogu-picture"},
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetContent(content)
	get, err := g.Cfg().Get(gctx.New(), "logger")
	if err != nil {
		return
	}
	fmt.Println(get)
	if err != nil {
		panic(err)
	}

	if !success {
		panic("注册失败")
	}

}
