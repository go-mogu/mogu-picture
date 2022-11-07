package model

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosProperties struct {
	Client    constant.ClientConfig
	Server    []constant.ServerConfig
	Config    vo.ConfigParam
	Discovery vo.RegisterInstanceParam
}
