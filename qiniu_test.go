package main

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"testing"
)

var (
	localFile = "D:\\Pictures\\Saved Pictures\\20221015012136.jpg"
	bucket    = "go-mogu"
	key       = "hn-bkt-clouddn-com-idvcz9d.www.qiniudns.com"
	accessKey = "***"
	secretKey = "***"
)

func TestUpload(t *testing.T) {
	putPolicy := storage.PutPolicy{
		Scope: "go-mogu",
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)

}

func TestJson(t *testing.T) {
	jsonStr := `[
    {
        "uid": "a36dc715e23f6f8ac73db7a7a0eb7693",
        "adminUid": "1f01cd1d2f474743b241d74008b12333",
        "extendName": "jpg",
        "fileName": "1668272637531.jpg",
        "filePath": "/test/",
        "fileSize": 36050,
        "isDir": 0,
        "status": 1,
        "createTime": "2022-11-13 01:03:58",
        "updateTime": "2022-11-13 01:03:58",
        "localUrl": "/blog/admin/jpg/2022/11/13/1668272638082.jpg",
        "qiNiuUrl": "d66fc0e81a3f4c339b50e75251008b4c",
        "fileOldName": "20221015012136.jpg",
        "minioUrl": "/go-mogu/1668272637884.jpg",
        "oldFilePath": "",
        "newFilePath": "",
        "files": "",
        "fileType": 0,
        "fileUrl": "http://minio-api.ithhit.cn/go-mogu/1668272637884.jpg"
    },
    {
        "uid": "c498de0dee3f6a69d1b15fcb80a49a5b",
        "adminUid": "1f01cd1d2f474743b241d74008b12333",
        "extendName": "jpg",
        "fileName": "1668272460423.jpg",
        "filePath": "/test/",
        "fileSize": 36050,
        "isDir": 0,
        "status": 1,
        "createTime": "2022-11-13 01:01:01",
        "updateTime": "2022-11-13 01:01:01",
        "localUrl": "/blog/admin/jpg/2022/11/13/1668272460966.jpg",
        "qiNiuUrl": "9e75009393914626af4f0a9abebf4e66",
        "fileOldName": "20221015012136.jpg",
        "minioUrl": "/go-mogu/1668272460782.jpg",
        "oldFilePath": "",
        "newFilePath": "",
        "files": "",
        "fileType": 0,
        "fileUrl": "http://minio-api.ithhit.cn/go-mogu/1668272460782.jpg"
    },
    {
        "uid": "5564facad0398bb38f86c287471b8a09",
        "adminUid": "1f01cd1d2f474743b241d74008b12333",
        "extendName": "jpg",
        "fileName": "1668269747941.jpg",
        "filePath": "/test/",
        "fileSize": 36050,
        "isDir": 0,
        "status": 1,
        "createTime": "2022-11-13 00:15:49",
        "updateTime": "2022-11-13 00:15:49",
        "localUrl": "/blog/admin/jpg/2022/11/13/1668269748499.jpg",
        "qiNiuUrl": "9776b15f349d4c2b9f175d221bb31c00",
        "fileOldName": "20221015012136.jpg",
        "minioUrl": "/go-mogu/1668269748303.jpg",
        "oldFilePath": "",
        "newFilePath": "",
        "files": "",
        "fileType": 0,
        "fileUrl": "http://minio-api.ithhit.cn/go-mogu/1668269748303.jpg"
    }
]`
	list := make(model.NetworkDiskList, 0)
	content, err := gjson.LoadContent(jsonStr)
	err = content.Scan(&list)
	if err != nil {
		panic(err)
	}
	g.Dump(list)
}
