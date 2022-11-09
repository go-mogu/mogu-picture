package file

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/app/picture/util"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/go-mogu/mogu-picture/internal/consts/EFilePriority"
	"github.com/go-mogu/mogu-picture/internal/consts/EOpenStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/EStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	service.RegisterFile(New())
}

type sFile struct{}

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

// CropperPicture 截图上传
func (s *sFile) CropperPicture(ctx context.Context, multipartFileList []*ghttp.UploadFile) (listMap []map[string]interface{}, err error) {
	systemConfig, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	qiNiuPictureBaseUrl := systemConfig.QiNiuPictureBaseUrl
	localPictureBaseUrl := systemConfig.LocalPictureBaseUrl
	minioPictureBaseUrl := systemConfig.MinioPictureBaseUrl
	result, err := s.BatchUploadFile(ctx, multipartFileList, systemConfig)
	utils.ErrIsNil(ctx, err)
	if len(result) > 0 {
		for _, file := range result {
			item := make(map[string]interface{})
			item[SysConf.UID] = file.Uid
			if EFilePriority.QI_NIU == systemConfig.PicturePriority {
				item[SysConf.URL] = qiNiuPictureBaseUrl + file.QiNiuUrl
			} else if EFilePriority.MINIO == systemConfig.PicturePriority {
				item[SysConf.URL] = minioPictureBaseUrl + file.MinioUrl
			} else {
				item[SysConf.URL] = localPictureBaseUrl + file.PicUrl
			}
			listMap = append(listMap, item)
		}
	}
	return
}

func (s *sFile) BatchUploadFile(ctx context.Context, fileList []*ghttp.UploadFile, systemConfig baseModel.SystemConfig) (result []*entity.File, err error) {
	uploadQiNiu := systemConfig.UploadQiNiu
	uploadLocal := systemConfig.UploadLocal
	uploadMinio := systemConfig.UploadMinio
	request := g.RequestFromCtx(ctx)

	//如果是用户上传，则包含用户uid
	var userUid = ""
	//如果是管理员上传，则包含管理员uid
	var adminUid = ""
	//项目名
	var projectName = ""
	//模块名
	var sortName = ""

	// 判断图片来源
	// 当图片从mogu-admin传递过来的时候
	userUid = request.Get(SysConf.USER_UID).String()
	adminUid = request.Get(SysConf.ADMIN_UID).String()
	projectName = request.Get(SysConf.PROJECT_NAME).String()
	sortName = request.Get(SysConf.SORT_NAME).String()
	//projectName现在默认base
	if projectName == "" {
		projectName = "base"
	}

	//TODO 检测用户上传，如果不是网站的用户就不能调用
	if userUid == "" && adminUid == "" {
		return result, gerror.New("请先注册")
	}
	sortColumns := dao.FileSort.Columns()
	var fileSorts = make([]*entity.FileSort, 0)
	err = dao.FileSort.Ctx(ctx).
		Where(g.Map{
			sortColumns.SortName:    sortName,
			sortColumns.ProjectName: projectName,
			sortColumns.Status:      EStatus.ENABLE,
		}).Scan(&fileSorts)
	utils.ErrIsNil(ctx, err)
	fileSort := new(entity.FileSort)
	if len(fileSorts) >= 1 {
		fileSort = fileSorts[0]
	} else {
		return result, gerror.New("文件不被允许上传")
	}
	//文件上传
	if len(fileList) > 0 {
		result = []*entity.File{}
		for _, multipartFile := range fileList {
			oldName := multipartFile.Filename
			size := multipartFile.Size
			//获取扩展名，默认是jpg
			picExpandedName := util.GetPicExpandedName(oldName)

			// 检查是否是安全的格式
			if !utils.IsPic(picExpandedName) {
				return result, gerror.New("请上传正确格式的文件")
			}
			//获取新文件名
			newFileName := fmt.Sprintf("%d%s%s", gtime.Now().UnixMilli(), Constants.SYMBOL_POINT, picExpandedName)
			localUrl := ""
			qiNiuUrl := ""
			minioUrl := ""
			// 上传七牛云，判断是否能够上传七牛云
			if EOpenStatus.OPEN == uploadQiNiu {
				fmt.Println(uploadQiNiu)
			}

			// 判断是否能够上传Minio文件服务器
			if EOpenStatus.OPEN == uploadMinio {
				fmt.Println(uploadMinio)
			}
			//TODO 先实现本地上传
			if EOpenStatus.OPEN == uploadLocal {
				localUrl, err = service.LocalFileService().UploadFile(ctx, multipartFile, *fileSort)
				utils.ErrIsNil(ctx, err)
			}
			file := entity.File{
				FileOldName:     oldName,
				PicName:         newFileName,
				PicUrl:          localUrl,
				PicExpandedName: picExpandedName,
				FileSize:        gconv.Int(size),
				FileSortUid:     fileSort.Uid,
				AdminUid:        adminUid,
				UserUid:         userUid,
				Status:          EStatus.ENABLE,
				CreateTime:      gtime.Now(),
				UpdateTime:      gtime.Now(),
				QiNiuUrl:        qiNiuUrl,
				MinioUrl:        minioUrl,
				Uid:             guid.S(),
			}
			_, err = dao.File.Ctx(ctx).Save(file)
			utils.ErrIsNil(ctx, err)
			result = append(result, &file)
		}
		//保存成功返回数据
		return
	}
	return result, gerror.New("请上传图片")
}
