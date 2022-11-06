package logic

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"mogu-picture/internal/app/picture/dao"
	"mogu-picture/internal/app/picture/model"
	"mogu-picture/internal/app/picture/model/entity"
	"mogu-picture/internal/app/picture/service"
	"mogu-picture/internal/consts"
	utils "mogu-picture/utility"
)

func init() {
	service.RegisterFile(New())
}

type sFile struct{}

var insFile = sFile{}

// New returns the interface of Table service.
func New() *sFile {
	return &sFile{}
}

// PageList 分页查询文件表
func (s *sFile) PageList(ctx context.Context, param model.File) (total int, result []*entity.File, err error) {
	result = make([]*entity.File, 0)
	daoModel := dao.File.Ctx(ctx)
	columnMap := dao.File.ColumnMap()
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
	err = daoModel.Page(param.PageNum, param.PageSize).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// List 列表查询文件表
func (s *sFile) List(ctx context.Context, param entity.File) (result []*entity.File, err error) {
	result = make([]*entity.File, 0)
	daoModel := dao.File.Ctx(ctx)
	columnMap := dao.File.ColumnMap()
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

// Get 查询文件表详情
func (s *sFile) Get(ctx context.Context, uid string) (result *entity.File, err error) {
	result = new(entity.File)
	daoModel := dao.File.Ctx(ctx)
	err = daoModel.Where(dao.File.Columns().Uid, uid).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// Add 添加文件表
func (s *sFile) Add(ctx context.Context, in model.File) (err error) {
	err = dao.File.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.File.Table()).Ctx(ctx).OmitEmpty().Data(in).Insert()
		return err
	})
	return
}

// Edit 编辑文件表
func (s *sFile) Edit(ctx context.Context, in model.File) (err error) {
	err = dao.File.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.File.Table()).Ctx(ctx).OmitEmpty().Data(in).Where(g.Map{
			dao.File.Columns().Uid: in.Uid,
		}).Update()
		return err
	})
	return
}

// EditState 编辑文件表状态
func (s *sFile) EditState(ctx context.Context, ids []string, state int8) (err error) {
	err = dao.File.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.File.Table()).Ctx(ctx).Data(g.Map{consts.StateColumn: state}).Where(g.Map{
			dao.File.Columns().Uid: ids,
		}).Update()
		return err
	})
	return
}

// Delete 删除文件表
func (s *sFile) Delete(ctx context.Context, ids []string) (err error) {
	err = dao.File.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.File.Table()).Ctx(ctx).Where(g.Map{
			dao.File.Columns().Uid: ids,
		}).Delete()
		return err
	})
	return
}
