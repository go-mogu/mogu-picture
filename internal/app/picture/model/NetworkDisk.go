package model

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
)

type NetworkDisk struct {
	entity.NetworkDisk
	OldFilePath string `json:"oldFilePath" dc:"旧文件名" q:"-"`
	NewFilePath string `json:"newFilePath" dc:"新文件目录" q:"-"`
	Files       string `json:"files" dc:"文件" q:"-"`
	FileType    int    `json:"fileType" dc:"文件类型" q:"-"`
	FileUrl     string `json:"fileUrl" dc:"文件URL" q:"-"`
}

type NetworkDiskList []NetworkDisk
