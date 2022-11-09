package networkDisk

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
	service.RegisterNetworkDisk(New())
}

type sNetworkDisk struct{}

var insNetworkDisk = sNetworkDisk{}

// New returns the interface of Table service.
func New() *sNetworkDisk {
	return &sNetworkDisk{}
}

// PageList 分页查询网盘文件表
func (s *sNetworkDisk) PageList(ctx context.Context, param model.NetworkDisk) (total int, result []*entity.NetworkDisk, err error) {
	result = make([]*entity.NetworkDisk, 0)
	daoModel := dao.NetworkDisk.Ctx(ctx)
	columnMap := dao.NetworkDisk.ColumnMap()
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

// List 列表查询网盘文件表
func (s *sNetworkDisk) List(ctx context.Context, param entity.NetworkDisk) (result []*entity.NetworkDisk, err error) {
	result = make([]*entity.NetworkDisk, 0)
	daoModel := dao.NetworkDisk.Ctx(ctx)
	columnMap := dao.NetworkDisk.ColumnMap()
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

// Get 查询网盘文件表详情
func (s *sNetworkDisk) Get(ctx context.Context, uid string) (result *entity.NetworkDisk, err error) {
	result = new(entity.NetworkDisk)
	daoModel := dao.NetworkDisk.Ctx(ctx)
	err = daoModel.Where(dao.NetworkDisk.Columns().Uid, uid).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

// Add 添加网盘文件表
func (s *sNetworkDisk) Add(ctx context.Context, in model.NetworkDisk) (err error) {
	err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).OmitEmpty().Data(in).Insert()
		return err
	})
	return
}

// Edit 编辑网盘文件表
func (s *sNetworkDisk) Edit(ctx context.Context, in model.NetworkDisk) (err error) {
	err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).OmitEmpty().Data(in).Where(g.Map{
			dao.NetworkDisk.Columns().Uid: in.Uid,
		}).Update()
		return err
	})
	return
}

// EditState 编辑网盘文件表状态
func (s *sNetworkDisk) EditState(ctx context.Context, ids []string, state int8) (err error) {
	err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).Data(g.Map{consts.StateColumn: state}).Where(g.Map{
			dao.NetworkDisk.Columns().Uid: ids,
		}).Update()
		return err
	})
	return
}

// Delete 删除网盘文件表
func (s *sNetworkDisk) Delete(ctx context.Context, ids []string) (err error) {
	err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).Where(g.Map{
			dao.NetworkDisk.Columns().Uid: ids,
		}).Delete()
		return err
	})
	return
}
