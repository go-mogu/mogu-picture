package localFileService

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/app/picture/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLocalFileService struct{}

var (
	path = g.Cfg().MustGet(gctx.New(), "file.upload.path").String()
)

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

func (s *sLocalFileService) UploadFile(ctx context.Context, newFileName string, multipartFile *ghttp.UploadFile, fileSort entity.FileSort) (result string, err error) {
	return uploadSingleFile(ctx, newFileName, multipartFile, fileSort)
}

/**
 * 本地文件服务图片上传【上传到本地硬盘】
 *
 * @return
 */
func uploadSingleFile(ctx context.Context, newFileName string, multipartFile *ghttp.UploadFile, fileSort entity.FileSort) (result string, err error) {
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

func (s *sLocalFileService) UploadPictureByUrl(ctx context.Context, itemUrl string, newFileName string, fileSort entity.FileSort) (fileUrl string, err error) {
	//判断url是否为空，如果为空，使用默认
	sortUrl := fileSort.Url
	if sortUrl == "" {
		sortUrl = "base/common/"
	}
	//获取新文件名
	now := gtime.Now()
	newPath := path + sortUrl + "/jpg/" + gconv.String(now.Year()) + "/" + gconv.String(now.Month()) + "/" + gconv.String(now.Day()) + "/"
	fileUrl = sortUrl + "/jpg/" + gconv.String(now.Year()) + "/" + gconv.String(now.Month()) + "/" + gconv.String(now.Day()) + "/" + newFileName
	saveUrl := newPath + newFileName
	if !gfile.Exists(newPath) {
		_ = gfile.Mkdir(newPath)
	}
	bytes := g.Client().GetBytes(ctx, itemUrl)
	err = gfile.PutBytes(saveUrl, bytes)
	return
}
