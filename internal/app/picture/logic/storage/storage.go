package storage

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterStorage(New())
}

type sStorage struct{}

var insStorage = sStorage{}

// New returns the interface of Table service.
func New() *sStorage {
	return &sStorage{}
}

// PageList 分页查询存储信息表
func (s *sStorage) PageList(ctx context.Context, param model.Storage) (total int, result []*entity.Storage, err error) {
	result = make([]*entity.Storage, 0)
	daoModel := dao.Storage.Ctx(ctx)
	columnMap := dao.Storage.ColumnMap()
	daoModel, err = utils.GetWrapper(param, daoModel, columnMap)
	if err != nil {
		return
	}
	total, err = daoModel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取总行数失败")
		return
	}
	if total == 0 {
		return
	}
	err = daoModel.Page(param.CurrentPage, param.PageSize).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// List 列表查询存储信息表
func (s *sStorage) List(ctx context.Context, param entity.Storage) (result []*entity.Storage, err error) {
	result = make([]*entity.Storage, 0)
	daoModel := dao.Storage.Ctx(ctx)
	columnMap := dao.Storage.ColumnMap()
	daoModel, err = utils.GetWrapper(param, daoModel, columnMap)
	if err != nil {
		return
	}
	err = daoModel.Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// Get 查询存储信息表详情
func (s *sStorage) Get(ctx context.Context, uid string) (result *entity.Storage, err error) {
	result = new(entity.Storage)
	daoModel := dao.Storage.Ctx(ctx)
	err = daoModel.Where(dao.Storage.Columns().Uid, uid).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// Add 添加存储信息表
func (s *sStorage) Add(ctx context.Context, in model.Storage) (err error) {
	err = dao.Storage.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.Storage.Table()).Ctx(ctx).OmitEmpty().Data(in).Insert()
		return err
	})
	return
}

// Edit 编辑存储信息表
func (s *sStorage) Edit(ctx context.Context, in model.Storage) (err error) {
	err = dao.Storage.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.Storage.Table()).Ctx(ctx).OmitEmpty().Data(in).Where(g.Map{
			dao.Storage.Columns().Uid: in.Uid,
		}).Update()
		return err
	})
	return
}

// EditState 编辑存储信息表状态
func (s *sStorage) EditState(ctx context.Context, ids []string, state int8) (err error) {
	err = dao.Storage.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.Storage.Table()).Ctx(ctx).Data(g.Map{consts.StateColumn: state}).Where(g.Map{
			dao.Storage.Columns().Uid: ids,
		}).Update()
		return err
	})
	return
}

// Delete 删除存储信息表
func (s *sStorage) Delete(ctx context.Context, ids []string) (err error) {
	err = dao.Storage.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.Storage.Table()).Ctx(ctx).Where(g.Map{
			dao.Storage.Columns().Uid: ids,
		}).Delete()
		return err
	})
	return
}
