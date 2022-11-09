package dao

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao/internal"
)

type storageDao struct {
	*internal.StorageDao
}

var (
	Storage = storageDao{
		internal.NewStorageDao(),
	}
)
