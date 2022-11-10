package feign

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign/adminFeignClient"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign/webFeignClient"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/go-mogu/mogu-picture/internal/consts/EOpenStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/ErrorCode"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	"github.com/go-mogu/mogu-picture/internal/consts/RedisConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	"github.com/go-mogu/mogu-picture/internal/model"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"time"
)

func GetSystemConfig(ctx context.Context) (systemConfig model.SystemConfig, err error) {
	request := g.RequestFromCtx(ctx)
	// 后台携带的token
	token := request.Get(SysConf.TOKEN).String()
	// 参数中携带的token
	paramsToken := request.GetQuery(SysConf.TOKEN).String()
	// 获取平台【web：门户，admin：管理端】
	platform := request.Get(SysConf.PLATFORM).String()
	systemConfigMap := map[string]string{}
	// 判断是否是web端发送过来的请求【后端发送过来的token长度为32】
	if SysConf.WEB == platform || (paramsToken != "" && len(paramsToken) == Constants.THIRTY_TWO) {
		// 如果是调用web端获取配置的接口
		systemConfigMap, err = getSystemConfigMap(ctx, paramsToken, platform)
	} else {
		// 调用admin端获取配置接口
		if token != "" {
			// 判断是否是后台过来的请求
			systemConfigMap, err = getSystemConfigMap(ctx, token, SysConf.ADMIN)
		} else {
			// 判断是否是通过params参数传递过来的
			systemConfigMap, err = getSystemConfigMap(ctx, paramsToken, SysConf.ADMIN)
		}
	}
	if reflect.DeepEqual(systemConfigMap, map[string]string{}) {
		g.Log().Error(ctx, MessageConf.PLEASE_SET_QI_NIU)
		return systemConfig, gerror.New(MessageConf.PLEASE_SET_QI_NIU)
	}

	systemConfig = model.SystemConfig{}
	uploadQiNiu := systemConfigMap[SysConf.UPLOAD_QI_NIU]
	uploadLocal := systemConfigMap[SysConf.UPLOAD_LOCAL]
	localPictureBaseUrl := systemConfigMap[SysConf.LOCAL_PICTURE_BASE_URL]
	qiNiuPictureBaseUrl := systemConfigMap[SysConf.QI_NIU_PICTURE_BASE_URL]
	qiNiuAccessKey := systemConfigMap[SysConf.QI_NIU_ACCESS_KEY]
	qiNiuSecretKey := systemConfigMap[SysConf.QI_NIU_SECRET_KEY]
	qiNiuBucket := systemConfigMap[SysConf.QI_NIU_BUCKET]
	qiNiuArea := systemConfigMap[SysConf.QI_NIU_AREA]
	minioEndPoint := systemConfigMap[SysConf.MINIO_END_POINT]
	minioAccessKey := systemConfigMap[SysConf.MINIO_ACCESS_KEY]
	minioSecretKey := systemConfigMap[SysConf.MINIO_SECRET_KEY]
	minioBucket := systemConfigMap[SysConf.MINIO_BUCKET]
	uploadMinio := systemConfigMap[SysConf.UPLOAD_MINIO]
	minioPictureBaseUrl := systemConfigMap[SysConf.MINIO_PICTURE_BASE_URL]

	// 判断七牛云参数是否存在异常

	if EOpenStatus.OPEN == uploadQiNiu && (qiNiuPictureBaseUrl == "" || qiNiuAccessKey == "" || qiNiuSecretKey == "" || qiNiuBucket == "" || qiNiuArea == "") {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_QI_NIU, "", nil), MessageConf.PLEASE_SET_QI_NIU)
	}
	// 判断本地服务参数是否存在异常
	if EOpenStatus.OPEN == uploadLocal && localPictureBaseUrl == "" {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_LOCAL, "", nil), MessageConf.PLEASE_SET_LOCAL)
	}
	// 判断Minio服务是否存在异常
	if EOpenStatus.OPEN == uploadMinio && (minioEndPoint == "" || minioPictureBaseUrl == "" || minioAccessKey == "" || minioSecretKey == "" || minioBucket == "") {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_MINIO, "", nil), MessageConf.PLEASE_SET_MINIO)
	}

	systemConfig.QiNiuAccessKey = qiNiuAccessKey
	systemConfig.QiNiuSecretKey = qiNiuSecretKey
	systemConfig.QiNiuBucket = qiNiuBucket
	systemConfig.QiNiuArea = qiNiuArea
	systemConfig.UploadQiNiu = uploadQiNiu
	systemConfig.UploadLocal = uploadLocal
	systemConfig.MinioEndPoint = minioEndPoint
	systemConfig.MinioAccessKey = minioAccessKey
	systemConfig.MinioSecretKey = minioSecretKey
	systemConfig.MinioBucket = minioBucket
	systemConfig.MinioPictureBaseUrl = minioPictureBaseUrl
	systemConfig.UploadMinio = uploadMinio
	systemConfig.PicturePriority = systemConfigMap[SysConf.PICTURE_PRIORITY]
	systemConfig.LocalPictureBaseUrl = systemConfigMap[SysConf.LOCAL_PICTURE_BASE_URL]
	systemConfig.QiNiuPictureBaseUrl = systemConfigMap[SysConf.QI_NIU_PICTURE_BASE_URL]

	return systemConfig, nil
}

// GetSystemConfigByWebToken 通过Web端的token获取系统配置文件 【传入Admin端的token】
func GetSystemConfigByWebToken(ctx context.Context, token string) (resultMap map[string]string, err error) {
	// 判断该token的有效性
	webUserJsonResult, err := g.Redis().Do(ctx, "GET", RedisConf.USER_TOKEN+Constants.SYMBOL_COLON+token)
	if err != nil {
		return nil, err
	}
	if webUserJsonResult == nil {
		return nil, gerror.New(MessageConf.INVALID_TOKEN)
	}
	// 从Redis中获取的SystemConf 或者 通过feign获取的
	//从Redis中获取内容
	jsonResult, err := g.Redis().Do(ctx, "GET", RedisConf.SYSTEM_CONFIG)
	if err != nil {
		return nil, err
	}
	// 判断Redis中是否有数据
	if jsonResult != nil {
		resultMap = jsonResult.MapStrStr()
	} else {
		// 进行七牛云校验
		resultStr, err := webFeignClient.GetSystemConfig(ctx, token)
		if err != nil {
			return nil, err
		}
		resultVar, err := gjson.LoadContent(resultStr)
		if err != nil {
			return nil, err
		}
		resultTempMap := resultVar.Map()
		if resultTempMap[SysConf.CODE] != nil && SysConf.SUCCESS == gconv.String(resultTempMap[SysConf.CODE]) {
			resultMap = gconv.MapStrStr(resultTempMap[SysConf.DATA])
			//将从token存储到redis中，设置30分钟后过期
			g.Redis().Do(ctx, "SET", RedisConf.SYSTEM_CONFIG, gjson.New(resultMap).MustToTomlString(), 30, time.Minute)
		}

	}
	return
}

// 通过Token获取系统配置【返回Map类型】
func getSystemConfigMap(ctx context.Context, token, platform string) (resultMap map[string]string, err error) {
	// 判断该token的有效性
	webUserJsonResult, err := g.Redis().Do(ctx, "GET", RedisConf.USER_TOKEN+Constants.SYMBOL_COLON+token)
	if err != nil {
		return nil, err
	}
	if webUserJsonResult == nil {
		return nil, gerror.New(MessageConf.INVALID_TOKEN)
	}
	// 从Redis中获取的SystemConf 或者 通过feign获取的
	//从Redis中获取内容
	jsonResult, err := g.Redis().Do(ctx, "GET", RedisConf.SYSTEM_CONFIG)
	if err != nil {
		return nil, err
	}
	// 判断Redis中是否有数据
	if jsonResult != nil {
		resultMap = jsonResult.MapStrStr()
	} else {
		// 通过feign获取系统配置
		resultStr := ""
		if SysConf.WEB == platform {
			// 进行七牛云校验
			resultStr, err = webFeignClient.GetSystemConfig(ctx, token)
		} else {
			resultStr, err = adminFeignClient.GetSystemConfig(ctx)
		}
		if err != nil {
			return nil, err
		}
		resultVar, err := gjson.LoadContent(resultStr)
		if err != nil {
			return nil, err
		}
		resultTempMap := resultVar.Map()
		if resultTempMap[SysConf.CODE] != nil && SysConf.SUCCESS == gconv.String(resultTempMap[SysConf.CODE]) {
			resultMap = gconv.MapStrStr(resultTempMap[SysConf.DATA])
			//将从token存储到redis中，设置30分钟后过期
			g.Redis().Do(ctx, "SET", RedisConf.SYSTEM_CONFIG, gjson.New(resultMap).MustToTomlString(), 30, time.Minute)
		}
	}
	return
}

func GetSystemConfigByMap(systemConfigMap map[string]string) (systemConfig model.SystemConfig, err error) {
	if systemConfigMap == nil {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.SYSTEM_CONFIG_NOT_EXIST, "", nil), MessageConf.SYSTEM_CONFIG_NOT_EXIST)
	}
	err = gconv.Struct(systemConfigMap, &systemConfig)
	if err != nil {
		return systemConfig, err
	}
	// 判断七牛云参数是否存在异常
	if EOpenStatus.OPEN == systemConfig.UploadQiNiu && (systemConfig.QiNiuPictureBaseUrl == "" || systemConfig.QiNiuAccessKey == "" || systemConfig.QiNiuSecretKey == "" || systemConfig.QiNiuBucket == "" || systemConfig.QiNiuArea == "") {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_QI_NIU, "", nil), MessageConf.PLEASE_SET_QI_NIU)
	}
	// 判断本地服务参数是否存在异常
	if EOpenStatus.OPEN == systemConfig.UploadLocal && systemConfig.LocalPictureBaseUrl == "" {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_LOCAL, "", nil), MessageConf.PLEASE_SET_LOCAL)
	}
	// 判断Minio服务是否存在异常
	if EOpenStatus.OPEN == systemConfig.UploadMinio && (systemConfig.MinioEndPoint == "" || systemConfig.MinioPictureBaseUrl == "" || systemConfig.MinioAccessKey == "" || systemConfig.MinioSecretKey == "" || systemConfig.MinioBucket == "") {
		return systemConfig, gerror.NewCode(gcode.New(ErrorCode.PLEASE_SET_MINIO, "", nil), MessageConf.PLEASE_SET_MINIO)
	}
	return
}
