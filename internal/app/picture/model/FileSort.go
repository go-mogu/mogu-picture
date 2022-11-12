package model

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/model"
)

type FileSort struct {
	model.PageInfo
	entity.FileSort
}
