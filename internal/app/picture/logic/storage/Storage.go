package storage

import (
	"context"
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/EStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SQLConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterStorage(New())
}

type sStorage struct{}

// New returns the interface of Table service.
func New() *sStorage {
	return &sStorage{}
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

func (s *sStorage) GetStorageByAdmin(ctx context.Context, adminUid string) (result *model.Storage, err error) {
	result = new(model.Storage)
	err = dao.Storage.Ctx(ctx).Where(g.Map{
		SysConf.STATUS:    EStatus.ENABLE,
		SQLConf.ADMIN_UID: adminUid,
	}).Scan(&result)
	return
}

func (s *sStorage) InitStorageSize(ctx context.Context, adminUid string, maxStorageSize int64) (err error) {
	storage, err := s.GetStorageByAdmin(ctx, adminUid)
	utils.ErrIsNil(ctx, err, MessageConf.ENTITY_EXIST)
	if storage != nil {
		return gerror.New(MessageConf.ENTITY_EXIST)
	} else {
		saveStorage := entity.Storage{
			AdminUid:       adminUid,
			StorageSize:    0,
			MaxStorageSize: maxStorageSize,
		}
		_, err = dao.Storage.Ctx(ctx).Data(saveStorage).OmitEmpty().Insert()
	}
	return
}

func (s *sStorage) EditStorageSize(ctx context.Context, adminUid string, maxStorageSize int64) (err error) {
	storage, err := s.GetStorageByAdmin(ctx, adminUid)
	utils.ErrIsNil(ctx, err, "未分配存储空间，重新初始化网盘空间！")
	if storage == nil {
		return gerror.New("未分配存储空间，重新初始化网盘空间！")
	} else {
		if maxStorageSize < storage.StorageSize {
			return gerror.New("网盘容量不能小于当前已用空间")
		}
		storage.MaxStorageSize = maxStorageSize
		_, err = dao.Storage.Ctx(ctx).Data(storage).OmitEmpty().Where(dao.Storage.Columns().Uid, storage.Uid).Update()
	}
	return
}

func (s *sStorage) GetStorageByAdminUid(ctx context.Context, adminUidList []string) (result []*model.Storage, err error) {
	result = make([]*model.Storage, 0)
	err = dao.Storage.Ctx(ctx).Where(SysConf.STATUS, EStatus.ENABLE).WhereIn(SQLConf.ADMIN_UID, adminUidList).Scan(&result)
	return
}

func (s *sStorage) UploadFile(ctx context.Context, networkDisk model.NetworkDisk) (result string, err error) {
	request := g.RequestFromCtx(ctx)
	systemConfig, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	// 计算文件大小
	var newStorageSize int64 = 0
	var storageSize int64
	fileList := make([]*model.UploadFileInfo, 0)
	for fileName := range request.MultipartForm.File {
		file := request.GetUploadFile(fileName)
		newStorageSize = newStorageSize + file.Size
		fileReader, err := file.Open()
		if err != nil {
			return "", err
		}
		bytes := make([]byte, file.Size)
		_, err = fileReader.Read(bytes)
		if err != nil {
			return "", err
		}
		fileReader.Close()
		fileList = append(fileList, &model.UploadFileInfo{
			Data: bytes,
			Size: file.Size,
			Name: fileName,
		})
	}
	adminUid := request.Get(SysConf.ADMIN_UID).String()
	storage, err := s.GetStorageByAdmin(ctx, adminUid)
	utils.ErrIsNil(ctx, err, "上传失败，您没有分配可用的上传空间！")
	storageSize = storage.StorageSize + newStorageSize
	// 判断上传的文件是否超过了剩余空间
	if storage.MaxStorageSize < storageSize {
		return "", gerror.New("上传失败，您没有分配可用的上传空间！")
	}

	// 上传文件
	files, err := service.File().BatchUploadFile(ctx, fileList, systemConfig)
	utils.ErrIsNil(ctx, err)
	networkDiskList := make([]entity.NetworkDisk, 0)
	for _, file := range files {
		saveNetworkDisk := entity.NetworkDisk{
			AdminUid:    adminUid,
			ExtendName:  file.PicExpandedName,
			FileName:    file.PicName,
			FilePath:    networkDisk.FilePath,
			FileSize:    file.FileSize,
			LocalUrl:    file.PicUrl,
			QiNiuUrl:    file.QiNiuUrl,
			FileOldName: file.FileOldName,
			MinioUrl:    file.MinioUrl,
		}
		networkDiskList = append(networkDiskList, saveNetworkDisk)
	}
	// 上传文件
	_, err = dao.NetworkDisk.Ctx(ctx).Data(networkDiskList).OmitEmpty().Save()
	utils.ErrIsNil(ctx, err)
	// 更新容量大小
	_, err = dao.Storage.Ctx(ctx).Where(SysConf.UID, storage.Uid).Data(g.Map{dao.Storage.Columns().StorageSize: storageSize}).Update()
	result = MessageConf.OPERATION_SUCCESS
	return
}
