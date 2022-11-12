package fileSort

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
	service.RegisterFileSort(New())
}

type sFileSort struct{}

var insFileSort = sFileSort{}

// New returns the interface of Table service.
func New() *sFileSort {
	return &sFileSort{}
}

// PageList 分页查询文件分类表
func (s *sFileSort) PageList(ctx context.Context, param model.FileSort) (total int, result []*entity.FileSort, err error) {
	result = make([]*entity.FileSort, 0)
	daoModel := dao.FileSort.Ctx(ctx)
	columnMap := dao.FileSort.ColumnMap()
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

// List 列表查询文件分类表
func (s *sFileSort) List(ctx context.Context, param entity.FileSort) (result []*entity.FileSort, err error) {
	result = make([]*entity.FileSort, 0)
	daoModel := dao.FileSort.Ctx(ctx)
	columnMap := dao.FileSort.ColumnMap()
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

// Get 查询文件分类表详情
func (s *sFileSort) Get(ctx context.Context, uid string) (result *entity.FileSort, err error) {
	result = new(entity.FileSort)
	daoModel := dao.FileSort.Ctx(ctx)
	err = daoModel.Where(dao.FileSort.Columns().Uid, uid).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// Add 添加文件分类表
func (s *sFileSort) Add(ctx context.Context, in model.FileSort) (err error) {
	err = dao.FileSort.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.FileSort.Table()).Ctx(ctx).OmitEmpty().Data(in).Insert()
		return err
	})
	return
}

// Edit 编辑文件分类表
func (s *sFileSort) Edit(ctx context.Context, in model.FileSort) (err error) {
	err = dao.FileSort.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.FileSort.Table()).Ctx(ctx).OmitEmpty().Data(in).Where(g.Map{
			dao.FileSort.Columns().Uid: in.Uid,
		}).Update()
		return err
	})
	return
}

// EditState 编辑文件分类表状态
func (s *sFileSort) EditState(ctx context.Context, ids []string, state int8) (err error) {
	err = dao.FileSort.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.FileSort.Table()).Ctx(ctx).Data(g.Map{consts.StateColumn: state}).Where(g.Map{
			dao.FileSort.Columns().Uid: ids,
		}).Update()
		return err
	})
	return
}

// Delete 删除文件分类表
func (s *sFileSort) Delete(ctx context.Context, ids []string) (err error) {
	err = dao.FileSort.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.FileSort.Table()).Ctx(ctx).Where(g.Map{
			dao.FileSort.Columns().Uid: ids,
		}).Delete()
		return err
	})
	return
}
