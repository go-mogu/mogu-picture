package cloudStorage

import (
	"bytes"
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func init() {
	service.RegisterCloudStorage(consts.QiNiuKey, QiNiu())
}

type sQiNiuStorage struct{}

func (s *sQiNiuStorage) UploadFile(param model.UploadFileParam) (url string, err error) {
	urlInfo, err := gurl.ParseURL(param.SystemConfig.QiNiuPictureBaseUrl, -1)
	if err != nil {
		return "", err
	}
	putPolicy := storage.PutPolicy{
		Scope: param.SystemConfig.QiNiuBucket,
	}
	mac := qbox.NewMac(param.SystemConfig.QiNiuAccessKey, param.SystemConfig.QiNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          setQiNiuArea(param.SystemConfig.QiNiuArea),
		UseHTTPS:      urlInfo["fragment"] == "https",
		UseCdnDomains: false,
	}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}
	err = resumeUploader.Put(param.Ctx, &ret, upToken, gfile.Name(param.NewFileName), bytes.NewReader(param.Data), int64(len(param.Data)), &putExtra)
	utils.ErrIsNil(param.Ctx, err, "七牛云上传文件失败")
	return ret.Key, nil
}

func (s *sQiNiuStorage) DeleteFile(ctx context.Context, fileName string, systemConfig baseModel.SystemConfig) (err error) {
	mac := qbox.NewMac(systemConfig.QiNiuAccessKey, systemConfig.QiNiuSecretKey)
	urlInfo, err := gurl.ParseURL(systemConfig.QiNiuPictureBaseUrl, -1)
	if err != nil {
		return err
	}
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: urlInfo["fragment"] == "https",
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = setQiNiuArea(systemConfig.QiNiuArea)
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err = bucketManager.Delete(systemConfig.QiNiuBucket, fileName)
	return
}

// 设置七牛云上传区域（内部方法）
func setQiNiuArea(area string) *storage.Region {
	switch area {
	case "z0":
		return &storage.ZoneHuadong
	case "z1":
		return &storage.ZoneHuabei
	case "z2":
		return &storage.ZoneHuanan
	case "na0":
		return &storage.ZoneBeimei
	case "as0":
		return &storage.ZoneXinjiapo
	}
	return nil
}

// QiNiu returns the interface service.
func QiNiu() *sQiNiuStorage {
	return &sQiNiuStorage{}
}
