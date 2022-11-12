package lb

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/core/config"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	nacosModel "github.com/nacos-group/nacos-sdk-go/v2/model"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
)

const (
	GrayKey = "Gray-Version"
)

// GetServiceUrl 获取服务地址(灰色路由+权重随机)
func GetServiceUrl(ctx context.Context, serviceName string) (url string) {
	instances, err := config.GetInstances(serviceName)
	utils.ErrIsNil(ctx, err)
	if len(instances) < 1 {
		utils.ErrIsNil(ctx, gerror.Newf("%s service is not found", serviceName))
	}
	instance := getInstance(ctx, instances)
	if reflect.DeepEqual(instance, nacosModel.Instance{}) {
		utils.ErrIsNil(ctx, gerror.Newf("%s service is not found", serviceName))
	}
	return fmt.Sprintf("%s:%d", instance.Ip, instance.Port)
}

func getInstance(ctx context.Context, instances []nacosModel.Instance) nacosModel.Instance {
	request := g.RequestFromCtx(ctx)
	grayVersion := request.GetHeader(GrayKey)
	// 请求头设置了灰度版本号，则匹配nacos元数据中版本号相同的服务
	if grayVersion != "" {
		grayInstances := make([]nacosModel.Instance, 0)
		for _, instance := range instances {
			if grayVersion == instance.Metadata[GrayKey] {
				grayInstances = append(grayInstances, instance)
			}
		}
		if len(grayInstances) > 0 {
			if len(grayInstances) == 1 {
				return grayInstances[0]
			}
			instances = grayInstances
		}
	}
	//按权重查找服务
	return getHostByRandomWeight(ctx, instances)

}

func getHostByRandomWeight(ctx context.Context, hosts []nacosModel.Instance) nacosModel.Instance {
	g.Log().Debug(ctx, "entry randomWithWeight")
	if len(hosts) < 1 {
		g.Log().Debug(ctx, "hosts is nil")
		return nacosModel.Instance{}
	}
	balance := newWeightRandomBalance(hosts...)
	return balance.Next()
}

type weightRandomBalance struct {
	instances []nacosModel.Instance
	weights   []float64
	max       float64
}

func newWeightRandomBalance(wn ...nacosModel.Instance) weightRandomBalance {
	sort.Slice(wn, func(i, j int) bool {
		return wn[i].Weight < wn[j].Weight
	})
	totals := make([]float64, len(wn))
	runningTotal := 0.0
	for i, w := range wn {
		runningTotal += w.Weight
		totals[i] = runningTotal
	}
	return weightRandomBalance{instances: wn, weights: totals, max: runningTotal}
}

func (w *weightRandomBalance) Next() nacosModel.Instance {
	random := RandFloat64(0, 1)
	index := sort.SearchFloat64s(w.weights, random)
	if index < 0 {
		index = -index - 1
	} else {
		return w.instances[index]
	}
	if index < len(w.weights) {
		if random < w.weights[index] {
			return w.instances[index]
		}
	}
	return w.instances[len(w.instances)-1]
}

// RandInt64 随机整数
func RandInt64(min, max int64) int64 {
	if min >= max {
		return min
	}
	if min >= 0 {
		return rand.Int63n(max-min+1) + min
	}
	return rand.Int63n(max+(0-min)+1) - (0 - min)
}

// RandFloat64 生成指定范围的随机小数
func RandFloat64(min, max float64) float64 {
	//转为字符串
	minStr := strconv.FormatFloat(min, 'f', -1, 64)
	maxStr := strconv.FormatFloat(max, 'f', -1, 64)

	if len(minStr) != len(maxStr) {
		return 0
	} else {
		var prefix string
		var startSuffix string
		var endSuffix string
		sw := 0
		for i := 0; i < len(minStr); i++ {
			startStr := string(minStr[i])
			endStr := string(maxStr[i])
			if sw == 0 {
				if minStr[i] == maxStr[i] {
					prefix = prefix + startStr
				} else {
					sw = 1
					startSuffix = startSuffix + startStr
					endSuffix = endSuffix + endStr
				}
			} else {
				startSuffix = startSuffix + startStr
				endSuffix = endSuffix + endStr
			}
		}
		//前缀字符串转整数
		start, _ := strconv.ParseInt(startSuffix, 10, 64)
		end, _ := strconv.ParseInt(endSuffix, 10, 64)
		//生成随机整数
		r := RandInt64(start, end)
		//随机整数转字符串
		//前缀拼接随机字符串
		allStr := prefix + strconv.FormatInt(r, 10)
		//全字符串转浮点数
		result, _ := strconv.ParseFloat(allStr, 64)
		return result
	}
}
