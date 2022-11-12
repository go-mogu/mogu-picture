package cloudStorage

import (
	"bytes"
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"
)

func init() {
	service.RegisterCloudStorage(consts.MinioKey, NewMinio())
}

type sMinioStorage struct{}

func (s *sMinioStorage) UploadFile(param model.UploadFileParam) (url string, err error) {
	ctx := param.Ctx
	minioClient, err := getClient(param.SystemConfig)
	if err != nil {
		return "", err
	}
	// Make a new bucket
	bucketName := param.SystemConfig.MinioBucket
	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		return "", err
	} else {
		//不存在先创建桶
		if !exists {
			err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
			if err != nil {
				return "", err
			}
		}
	}
	fileType := http.DetectContentType(param.Data)

	info, err := minioClient.PutObject(ctx, bucketName, param.NewFileName, bytes.NewReader(param.Data), int64(len(param.Data)), minio.PutObjectOptions{ContentType: fileType})
	if err != nil {
		return "", err
	}
	//Key是minio存储的文件名(带文件夹)
	return Constants.PATH_SEPARATOR + param.SystemConfig.MinioBucket + Constants.PATH_SEPARATOR + info.Key, err
}

func (s *sMinioStorage) DeleteFile(ctx context.Context, fileName string, systemConfig baseModel.SystemConfig) (err error) {
	minioClient, err := getClient(systemConfig)
	if err != nil {
		return err
	}
	//去除前面的桶名
	gstr.Replace(fileName, Constants.PATH_SEPARATOR+systemConfig.MinioBucket, "")
	err = minioClient.RemoveObject(ctx, systemConfig.MinioBucket, fileName, minio.RemoveObjectOptions{})
	return
}

func getClient(systemConfig baseModel.SystemConfig) (*minio.Client, error) {
	endpoint := systemConfig.MinioEndPoint
	accessKeyID := systemConfig.MinioAccessKey
	secretAccessKey := systemConfig.MinioSecretKey
	urlInfo, err := gurl.ParseURL(endpoint, -1)
	if err != nil {
		return nil, err
	}
	endpoint = urlInfo["host"]
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: urlInfo["fragment"] == "https",
	})
	return minioClient, err
}

// NewMinio returns the interface service.
func NewMinio() *sMinioStorage {
	return &sMinioStorage{}
}
