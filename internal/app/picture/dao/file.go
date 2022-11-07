package dao

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao/internal"
)

type fileDao struct {
	*internal.FileDao
}

var (
	File = fileDao{
		internal.NewFileDao(),
	}
)
