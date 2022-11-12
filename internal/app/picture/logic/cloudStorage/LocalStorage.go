package cloudStorage

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterCloudStorage(consts.LocalKey, NewLocal())
}

var (
	path = g.Cfg().MustGet(gctx.New(), "file.upload.path").String()
)

type sLocalStorage struct{}

func (s *sLocalStorage) UploadFile(param model.UploadFileParam) (url string, err error) {
	//判断url是否为空，如果为空，使用默认
	sortUrl := param.FileSort.Url
	if sortUrl == "" {
		sortUrl = "base/common"
	}
	//获取新文件名
	now := gtime.Now()
	picUrl := sortUrl + Constants.PATH_SEPARATOR + param.File.PicExpandedName + Constants.PATH_SEPARATOR + gconv.String(now.Year()) + Constants.PATH_SEPARATOR + gconv.String(now.Month()) + Constants.PATH_SEPARATOR + gconv.String(now.Day()) + Constants.PATH_SEPARATOR + param.NewFileName
	newPath := path + Constants.PATH_SEPARATOR + picUrl
	// 序列化文件到本地
	err = gfile.PutBytes(newPath, param.Data)
	if err != nil {
		return "", err
	}
	return picUrl, nil
}

func (s *sLocalStorage) DeleteFile(ctx context.Context, fileName string, systemConfig baseModel.SystemConfig) (err error) {
	err = gfile.Remove(fileName)
	return
}

// NewLocal returns the interface service.
func NewLocal() *sLocalStorage {
	return &sLocalStorage{}
}
