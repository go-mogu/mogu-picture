package dao

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao/internal"
)

type fileSortDao struct {
	*internal.FileSortDao
}

var (
	FileSort = fileSortDao{
		internal.NewFileSortDao(),
	}
)
