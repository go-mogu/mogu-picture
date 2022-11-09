package model

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/model"
)

type Storage struct {
	model.PageReq
	entity.Storage
}
