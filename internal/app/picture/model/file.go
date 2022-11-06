package model

import (
	"mogu-picture/internal/app/picture/model/entity"
	"mogu-picture/internal/model"
)

type File struct {
	model.PageReq
	entity.File
}
