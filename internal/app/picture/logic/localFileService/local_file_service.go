package localFileService

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/app/picture/util"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLocalFileService struct{}

var (
	path = g.Cfg().MustGet(gctx.New(), "file.upload.path").String()
)

var insLocalFileService = sLocalFileService{}

// New returns the interface of Table service.
func New() *sLocalFileService {
	return &sLocalFileService{}
}

func init() {
	service.RegisterLocalFileService(New())
}

func (s *sLocalFileService) BatchUploadFile(ctx context.Context, multipartFileList []*ghttp.UploadFile, fileSort entity.FileSort) (list []string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sLocalFileService) UploadFile(ctx context.Context, multipartFile *ghttp.UploadFile, fileSort entity.FileSort) (result string, err error) {
	return uploadSingleFile(ctx, multipartFile, fileSort)
}

/**
 * 本地文件服务图片上传【上传到本地硬盘】
 *
 * @return
 */
func uploadSingleFile(ctx context.Context, multipartFile *ghttp.UploadFile, fileSort entity.FileSort) (result string, err error) {
	//判断url是否为空，如果为空，使用默认
	sortUrl := fileSort.Url
	if sortUrl == "" {
		sortUrl = "base/common/"
	}
	oldName := multipartFile.Filename
	//获取扩展名，默认是jpg
	picExpandedName := util.GetPicExpandedName(oldName)
	//获取新文件名
	now := gtime.Now()
	newFileName := fmt.Sprintf("%d%s%s", now.UnixMilli(), Constants.SYMBOL_POINT, picExpandedName)
	newPath := path + sortUrl + "/" + picExpandedName + "/" + gconv.String(now.Year()) + "/" + gconv.String(now.Month()) + "/" + gconv.String(now.Day()) + "/"

	picUrl := sortUrl + "/" + picExpandedName + "/" + gconv.String(now.Year()) + "/" + gconv.String(now.Month()) + "/" + gconv.String(now.Day()) + "/" + newFileName
	multipartFile.Filename = newFileName
	// 序列化文件到本地
	_, err = multipartFile.Save(newPath)
	if err != nil {
		return "", err
	}
	return picUrl, nil

}
