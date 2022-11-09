package lb

import (
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/core/config"
	"github.com/gogf/gf/v2/errors/gerror"
	nacosModel "github.com/nacos-group/nacos-sdk-go/v2/model"
)

var serviceMap = map[string]*serviceCW{}

type serviceCW struct {
	I   int
	Cw  float64
	Gcd float64
}

var i = -1          //表示上一次选择的服务器
var cw float64 = 0  //表示当前调度的权值
var gcd float64 = 2 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

func GetServiceUrl(serviceName string) (url string, err error) {

	if _, ok := serviceMap[serviceName]; !ok {
		serviceMap[serviceName] = &serviceCW{
			I:   i,
			Cw:  cw,
			Gcd: gcd,
		}
	}
	instances, err := config.GetInstances(serviceName)
	if err != nil {
		return
	}
	if len(instances) < 1 {
		return url, gerror.Newf("%s service is not found", serviceName)
	}
	s := serviceMap[serviceName]
	for {
		s.I = (s.I + 1) % len(instances)
		if s.I == 0 {
			s.Cw = s.Cw - s.Gcd
			if s.Cw <= 0 {
				s.Cw = getMaxWeight(instances)
				if s.Cw == 0 {
					return
				}
			}
		}

		if instances[s.I].Weight >= cw {
			return fmt.Sprintf("%s:%d", instances[s.I].Ip, instances[s.I].Port), err
		}
	}
}

func getMaxWeight(instances []nacosModel.Instance) float64 {
	max := float64(0)
	for _, v := range instances {
		if v.Weight >= max {
			max = v.Weight
		}
	}
	return max
}
