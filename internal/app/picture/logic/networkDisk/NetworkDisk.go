package networkDisk

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/go-mogu/mogu-picture/internal/consts/EFilePriority"
	"github.com/go-mogu/mogu-picture/internal/consts/EOpenStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/EStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/SQLConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	"github.com/go-mogu/mogu-picture/internal/model/queue"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/go-mogu/mogu-picture/utility/RequestHolder"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	service.RegisterNetworkDisk(New())
}

type sNetworkDisk struct{}

var treeId int64

// New returns the interface of Table service.
func New() *sNetworkDisk {
	return &sNetworkDisk{}
}

// GetFileList 列表查询网盘文件表
func (s *sNetworkDisk) GetFileList(ctx context.Context, networkDisk model.NetworkDisk) (result []*model.NetworkDisk, err error) {
	// 获取配置文件
	config, err := feign.GetSystemConfig(ctx)
	if err != nil {
		return
	}
	picturePriority := config.PicturePriority
	result = make([]*model.NetworkDisk, 0)
	columns := dao.NetworkDisk.Columns()
	if err != nil {
		return
	}
	daoModel := dao.NetworkDisk.Ctx(ctx).Where(columns.Status, EStatus.ENABLE).OrderAsc(columns.CreateTime)
	// 根据扩展名查找
	if networkDisk.FileType != 0 {
		// 判断是否是其它文件
		if utils.OTHER_TYPE == networkDisk.FileType {
			daoModel = daoModel.WhereNotIn(SQLConf.EXTEND_NAME, utils.GetFileExtendsByType(networkDisk.FileType))
		} else {
			daoModel = daoModel.WhereIn(SQLConf.EXTEND_NAME, utils.GetFileExtendsByType(networkDisk.FileType))
		}
	} else if networkDisk.FilePath != "" {
		daoModel = daoModel.Where(SQLConf.FILE_PATH, networkDisk.FilePath)
	}
	err = daoModel.Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	for _, item := range result {
		if EFilePriority.QI_NIU == picturePriority {
			item.FileUrl = config.QiNiuPictureBaseUrl + item.QiNiuUrl
		} else if EFilePriority.MINIO == picturePriority {
			item.FileUrl = config.MinioPictureBaseUrl + item.MinioUrl
		} else {
			item.FileUrl = config.LocalPictureBaseUrl + item.LocalUrl
		}
	}
	return
}

// CreateFile 添加网盘文件表
func (s *sNetworkDisk) CreateFile(ctx context.Context, in entity.NetworkDisk) (err error) {
	in.Uid = guid.S()
	if in.IsDir == 1 {
		in.FileSize = 0
	}
	now := gtime.Now()
	in.CreateTime = now
	in.UpdateTime = now
	in.Status = EStatus.ENABLE
	err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		_, err := tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).Data(in).Insert()
		return err
	})
	return
}

// Edit 编辑网盘文件表
func (s *sNetworkDisk) Edit(ctx context.Context, networkDiskVO model.NetworkDisk) (err error) {
	var oldFilePath = networkDiskVO.OldFilePath
	var newFilePath = networkDiskVO.NewFilePath
	var fileName = networkDiskVO.FileName
	var fileOldName = networkDiskVO.FileOldName
	var extendName = networkDiskVO.ExtendName

	if "null" == networkDiskVO.ExtendName {
		extendName = ""
	}
	// 判断移动的路径是否相同【拼接出原始目录】
	var fileOldPath = oldFilePath + fileOldName + "/"
	if fileOldPath == newFilePath {
		return gerror.New("不能选择自己")
	}
	queryWrapper := dao.NetworkDisk.Ctx(ctx).Where(g.Map{
		dao.NetworkDisk.Columns().FilePath: networkDiskVO.FilePath,
		dao.NetworkDisk.Columns().FileName: networkDiskVO.FileName,
	})
	if extendName != "" {
		queryWrapper = queryWrapper.Where(SQLConf.EXTEND_NAME, extendName)
	} else {
		queryWrapper = queryWrapper.WhereNull(SQLConf.EXTEND_NAME)
	}
	list := make([]*model.NetworkDisk, 0)
	err = queryWrapper.Scan(&list)
	if err != nil {
		return err
	}
	for _, networkDisk := range list {
		// 修改新的路径
		networkDisk.FilePath = newFilePath
		// 修改旧文件名
		networkDisk.FileOldName = networkDiskVO.FileOldName
		// 如果扩展名为空，代表是文件夹，还需要修改文件名
		if extendName == "" {
			networkDisk.FileName = networkDiskVO.FileOldName
		}
	}
	if len(list) > 0 {
		err = dao.NetworkDisk.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
			_, err = tx.Model(dao.NetworkDisk.Table()).Ctx(ctx).OmitEmpty().Data(list).Save()
			utils.ErrIsNil(ctx, err)
			//移动子目录
			oldFilePath = oldFilePath + fileName + "/"
			newFilePath = newFilePath + fileOldName + "/"
			oldFilePath = gstr.Replace(oldFilePath, "\\", "\\\\\\\\")
			oldFilePath = gstr.Replace(oldFilePath, "'", "\\'")
			oldFilePath = gstr.Replace(oldFilePath, "%", "\\%")
			oldFilePath = gstr.Replace(oldFilePath, "_", "\\_")
			//为null说明是目录，则需要移动子目录
			if extendName == "" {
				//移动根目录
				childList := make([]*model.NetworkDisk, 0)
				err = dao.NetworkDisk.Ctx(ctx).WhereLike(SQLConf.FILE_PATH, oldFilePath+"%").Scan(&childList)
				utils.ErrIsNil(ctx, err)
				if len(childList) > 0 {
					for _, networkDisk := range childList {
						networkDisk.FilePath = gstr.Replace(networkDisk.FilePath, oldFilePath, newFilePath)

					}
				}
			}
			return err
		})
	}

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

func (s *sNetworkDisk) BatchDeleteFile(ctx context.Context, networkDiskList model.NetworkDiskList) (err error) {
	RequestHolder.CheckLogin(ctx)
	config, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	for _, networkDisk := range networkDiskList {
		_ = s.DeleteFile(ctx, networkDisk, config)
	}
	return
}

func (s *sNetworkDisk) DeleteFile(ctx context.Context, networkDiskVO model.NetworkDisk, systemConfig baseModel.SystemConfig) (err error) {
	uid := networkDiskVO.Uid
	if uid == "" {
		return gerror.New("删除的文件不能为空")
	}
	networkDisk := new(model.NetworkDisk)
	err = dao.NetworkDisk.Ctx(ctx).Where(SysConf.UID, uid).Scan(&networkDisk)
	if err != nil {
		return err
	}
	// 修改为删除状态
	_, err = dao.NetworkDisk.Ctx(ctx).Data(SysConf.STATUS, EStatus.DISABLED).Where(SysConf.UID, uid).Update()
	if err != nil {
		return err
	}
	// 判断删除的是文件 or 文件夹
	if SysConf.ONE == networkDisk.IsDir {
		return s.deleteFileInDir(ctx, networkDisk, systemConfig)
	} else {
		return s.deleteFile(ctx, networkDisk, systemConfig)
	}
	return
}

func (s *sNetworkDisk) deleteFileInDir(ctx context.Context, networkDisk *model.NetworkDisk, systemConfig baseModel.SystemConfig) (err error) {
	// 删除的是文件夹，那么需要把文件下所有的文件获得，进行删除
	// 获取文件的路径，查询出该路径下所有的文件
	var path = networkDisk.FilePath + networkDisk.FileName
	list := make([]*model.NetworkDisk, 0)
	// 查询以  path%  开头的
	err = dao.NetworkDisk.Ctx(ctx).Where(SQLConf.STATUS, EStatus.ENABLE).WhereLike(SQLConf.FILE_PATH, path+"%").Scan(&list)
	if err != nil {
		return err
	}
	if len(list) > 0 {
		// 将所有的状态设置成失效
		ids := make([]string, 0)
		for _, item := range list {
			ids = append(ids, item.Uid)
		}
		err = s.EditState(ctx, ids, EStatus.DISABLED)
		if err != nil {
			return err
		}
		basePath := g.Cfg().MustGet(ctx, "file.upload.path").String()
		for _, removeFile := range list {
			if EOpenStatus.OPEN == systemConfig.UploadLocal {
				// 删除本地文件
				_ = service.CloudStorage(consts.LocalKey).DeleteFile(ctx, gfile.Join(basePath, removeFile.LocalUrl), systemConfig)
			}
			// 删除七牛云上文件
			if EOpenStatus.OPEN == systemConfig.UploadQiNiu {
				_ = service.CloudStorage(consts.QiNiuKey).DeleteFile(ctx, removeFile.QiNiuUrl, systemConfig)
			}
			// 删除Minio中的文件
			if EOpenStatus.OPEN == systemConfig.UploadMinio {
				_ = service.CloudStorage(consts.MinioKey).DeleteFile(ctx, removeFile.MinioUrl, systemConfig)
			}
		}

	}
	return
}

func (s *sNetworkDisk) deleteFile(ctx context.Context, networkDisk *model.NetworkDisk, systemConfig baseModel.SystemConfig) (err error) {
	// 删除本地文件，同时移除本地文件
	if EOpenStatus.OPEN == systemConfig.UploadLocal {
		// 删除本地文件
		basePath := g.Cfg().MustGet(ctx, "file.upload.path").String()
		_ = service.CloudStorage(consts.LocalKey).DeleteFile(ctx, gfile.Join(basePath, networkDisk.LocalUrl), systemConfig)
	}
	// 删除七牛云上文件
	if EOpenStatus.OPEN == systemConfig.UploadQiNiu {
		_ = service.CloudStorage(consts.QiNiuKey).DeleteFile(ctx, networkDisk.QiNiuUrl, systemConfig)
	}
	// 删除Minio中的文件
	if EOpenStatus.OPEN == systemConfig.UploadMinio {
		_ = service.CloudStorage(consts.MinioKey).DeleteFile(ctx, networkDisk.MinioUrl, systemConfig)
	}
	request := RequestHolder.GetRequest(ctx)
	adminUid := request.Get(SysConf.ADMIN_UID).String()
	storage, err := service.Storage().GetStorageByAdmin(ctx, adminUid)
	if err != nil {
		return gerror.New("本地文件空间不存在")
	}
	storageSize := storage.StorageSize - networkDisk.FileSize
	if storageSize > 0 {
		storage.StorageSize = storageSize
	} else {
		storage.StorageSize = 0
	}
	err = service.Storage().Edit(ctx, *storage)
	return
}

func (s *sNetworkDisk) UnzipFile(ctx context.Context, networkDisk model.NetworkDisk) (err error) {
	zipFileUrl := gfile.Join(gfile.MainPkgPath(), networkDisk.FileUrl)
	unzipUrl := gfile.Dir(zipFileUrl)
	fileEntryNameList, err := utils.Unzip(zipFileUrl, unzipUrl)
	if err != nil {
		return err
	}
	fileBeanList := make([]entity.NetworkDisk, 0)
	for _, entryName := range fileEntryNameList {
		totalFileUrl := unzipUrl + entryName
		tempFileBean := entity.NetworkDisk{
			CreateTime: gtime.Now(),
			AdminUid:   SysConf.DEFAULT_UID,
			FilePath:   utils.PathSplitFormat(networkDisk.FilePath + gfile.Dir(entryName)),
		}
		if gfile.IsDir(totalFileUrl) {
			tempFileBean.IsDir = 1
			tempFileBean.FileName = gfile.Basename(totalFileUrl)
		} else {
			tempFileBean.IsDir = 0
			tempFileBean.ExtendName = gfile.ExtName(totalFileUrl)
			tempFileBean.FileName = gfile.Name(totalFileUrl)
			tempFileBean.FileSize = gfile.Size(totalFileUrl)
		}
		fileBeanList = append(fileBeanList, tempFileBean)
	}
	_, err = dao.NetworkDisk.Ctx(ctx).Data(fileBeanList).OmitEmpty().Save()
	return
}

func (s *sNetworkDisk) MoveFile(ctx context.Context, networkDiskVO model.NetworkDisk) (err error) {
	return s.Edit(ctx, networkDiskVO)
}

func (s *sNetworkDisk) BatchMoveFile(ctx context.Context, networkDiskVO model.NetworkDisk) (err error) {
	newFilePath := networkDiskVO.NewFilePath
	json, err := gjson.LoadContent(networkDiskVO.Files)
	utils.ErrIsNil(ctx, err)
	fileList := make([]*model.NetworkDisk, 0)
	err = json.Scan(&fileList)
	utils.ErrIsNil(ctx, err)
	for _, file := range fileList {
		file.NewFilePath = newFilePath
		file.OldFilePath = file.FilePath
		_ = s.MoveFile(ctx, *file)
	}
	return
}

// SelectFileByFileType 通过拓展名查询文件
func (s *sNetworkDisk) SelectFileByFileType(ctx context.Context, networkDisk model.NetworkDisk) (result model.NetworkDiskList, err error) {
	fileExtends := utils.GetFileExtendsByType(networkDisk.FileType)
	fmt.Println(fileExtends)
	return
}

func (s *sNetworkDisk) GetFileTree(ctx context.Context) (result baseModel.TreeNode, err error) {
	filePathList := make([]*model.NetworkDisk, 0)
	err = dao.NetworkDisk.Ctx(ctx).Where(g.Map{
		SQLConf.STATUS: EStatus.ENABLE,
		SQLConf.IS_DIR: SysConf.ONE,
	}).Scan(&filePathList)
	result = baseModel.TreeNode{}
	result.Label = "/"
	for _, networkDisk := range filePathList {
		filePath := networkDisk.FilePath + networkDisk.FileName + Constants.PATH_SEPARATOR
		queue := queue.New()
		strArr := gstr.Split(filePath, Constants.PATH_SEPARATOR)
		for _, s := range strArr {
			queue.Push(s)
		}
		if queue.Len() == 0 {
			break
		}
		result = s.insertTreeNode(result, Constants.PATH_SEPARATOR, queue)
	}
	return
}

func (s *sNetworkDisk) insertTreeNode(treeNode baseModel.TreeNode, filepath string, nodeNameQueue *queue.Queue) baseModel.TreeNode {
	childrenTreeNodes := treeNode.Children
	currentNodeName := nodeNameQueue.Peek()
	if currentNodeName == "" {
		return treeNode
	}
	m := map[string]string{}
	filepath = filepath + currentNodeName.(string) + "/"
	m["filepath"] = filepath
	if !s.isExistPath(childrenTreeNodes, currentNodeName.(string)) {
		var resultTreeNode baseModel.TreeNode
		resultTreeNode.Attributes = m
		resultTreeNode.Label = nodeNameQueue.Pop().(string)
		treeId++
		resultTreeNode.Id = treeId
		childrenTreeNodes = append(childrenTreeNodes, resultTreeNode)
	} else {
		nodeNameQueue.Pop()
	}
	if nodeNameQueue.Len() != 0 {
		for i := 0; i < len(childrenTreeNodes); i++ {
			childrenTreeNode := childrenTreeNodes[i]
			if currentNodeName == childrenTreeNode.Label {
				childrenTreeNode = s.insertTreeNode(childrenTreeNode, filepath, nodeNameQueue)
				childrenTreeNodes = append(childrenTreeNodes[:i], childrenTreeNodes[i+1:]...)
				childrenTreeNodes = append(childrenTreeNodes, childrenTreeNode)
				treeNode.Children = childrenTreeNodes
			}
		}
	} else {
		treeNode.Children = childrenTreeNodes
	}
	return treeNode
}

func (s *sNetworkDisk) isExistPath(childrenTreeNodes []baseModel.TreeNode, path string) bool {
	isExistPath := false
	for i := 0; i < len(childrenTreeNodes); i++ {
		if path == childrenTreeNodes[i].Label {
			isExistPath = true
		}
	}
	return isExistPath
}
