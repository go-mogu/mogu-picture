package file

import (
	"context"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao"
	"github.com/go-mogu/mogu-picture/internal/app/picture/feign"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/do"
	"github.com/go-mogu/mogu-picture/internal/app/picture/model/entity"
	"github.com/go-mogu/mogu-picture/internal/app/picture/service"
	"github.com/go-mogu/mogu-picture/internal/consts"
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/go-mogu/mogu-picture/internal/consts/EFilePriority"
	"github.com/go-mogu/mogu-picture/internal/consts/EOpenStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/EStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/MessageConf"
	"github.com/go-mogu/mogu-picture/internal/consts/SysConf"
	baseModel "github.com/go-mogu/mogu-picture/internal/model"
	utils "github.com/go-mogu/mogu-picture/utility"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

func init() {
	service.RegisterFile(New())
}

type sFile struct{}

// New returns the interface of Table service.
func New() *sFile {
	return &sFile{}
}

// CropperPicture 截图上传
func (s *sFile) CropperPicture(ctx context.Context, uploadFile *ghttp.UploadFile) (result []map[string]interface{}, err error) {
	systemConfig, err := feign.GetSystemConfig(ctx)
	utils.ErrIsNil(ctx, err)
	result = make([]map[string]interface{}, 0)
	file, err := s.uploadFileByMultipartFile(ctx, systemConfig, uploadFile)
	if err != nil {
		return nil, err
	}
	result = append(result, file)
	return
}

func (s *sFile) uploadFileByMultipartFile(ctx context.Context, systemConfig baseModel.SystemConfig, uploadFile *ghttp.UploadFile) (result map[string]interface{}, err error) {
	fileList := make([]*model.UploadFileInfo, 0)
	fileReader, err := uploadFile.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()
	bytes := make([]byte, uploadFile.Size)
	_, err = fileReader.Read(bytes)
	if err != nil {
		return nil, err
	}

	fileList = append(fileList, &model.UploadFileInfo{
		Data: bytes,
		Size: uploadFile.Size,
		Name: uploadFile.Filename,
	})
	files, err := s.BatchUploadFile(ctx, fileList, systemConfig)
	utils.ErrIsNil(ctx, err)
	file := files[0]
	result = make(map[string]interface{})
	result[SysConf.UPLOADED] = 1
	result[SysConf.UID] = file.Uid
	result[SysConf.FILE_NAME] = file.PicName
	result[SysConf.URL] = file.PictureUrl
	return result, nil
}

// GetPicture 查询文件表详情
func (s *sFile) GetPicture(ctx context.Context, fileIds string, code string) (result []map[string]interface{}, err error) {
	if code == "" {
		code = Constants.SYMBOL_COMMA
	}
	if fileIds == "" {
		g.Log().Error(ctx, MessageConf.PICTURE_UID_IS_NULL)
		return nil, gerror.New(MessageConf.PICTURE_UID_IS_NULL)
	}
	result = make([]map[string]interface{}, 0)
	changeStringToString := gstr.Split(fileIds, code)
	fileList := make([]*entity.File, 0)
	daoModel := dao.File.Ctx(ctx)
	err = daoModel.Where(dao.File.Columns().Uid, changeStringToString).Scan(&fileList)
	utils.ErrIsNil(ctx, err)
	if len(fileList) > 0 {
		for _, file := range fileList {
			remap := make(map[string]interface{})
			// 获取七牛云地址
			remap[SysConf.QI_NIU_BUCKET] = file.QiNiuUrl
			// 获取Minio对象存储地址
			remap[SysConf.MINIO_URL] = file.MinioUrl
			// 获取本地地址
			remap[SysConf.URL] = file.PicUrl
			// 后缀名，也就是类型
			remap[SysConf.EXPANDED_NAME] = file.PicExpandedName
			//名称
			remap[SysConf.NAME] = file.PicName
			remap[SysConf.UID] = file.Uid
			remap[SysConf.FILE_OLD_NAME] = file.FileOldName
			result = append(result, remap)
		}
	}
	return
}

func (s *sFile) BatchUploadFile(ctx context.Context, fileList []*model.UploadFileInfo, systemConfig baseModel.SystemConfig) (result []*model.File, err error) {
	request := g.RequestFromCtx(ctx)
	//项目名
	var projectName = request.Get(SysConf.PROJECT_NAME).String()
	//模块名
	var sortName = request.Get(SysConf.SORT_NAME).String()
	//projectName现在默认base
	if projectName == "" {
		projectName = "blog"
	}
	fileSort, err := s.getFileSort(ctx, sortName, projectName)
	if err != nil || fileSort == nil {
		return nil, gerror.New("文件不被允许上传")
	}
	adminUid := request.Get(SysConf.ADMIN_UID).String()
	userUid := request.Get(SysConf.USER_UID).String()
	//文件上传
	if len(fileList) > 0 {
		result = []*model.File{}
		for _, multipartFile := range fileList {
			oldName := multipartFile.Name
			size := multipartFile.Size
			//获取扩展名，默认是jpg
			picExpandedName := utils.GetPicExpandedName(oldName)
			//获取新文件名
			newFileName := fmt.Sprintf("%d%s%s", gtime.Now().UnixMilli(), Constants.SYMBOL_POINT, picExpandedName)
			fileInfo := &model.File{}
			fileInfo.AdminUid = adminUid
			fileInfo.UserUid = userUid
			fileInfo.PicExpandedName = picExpandedName
			fileInfo.PicName = newFileName
			fileInfo.FileOldName = oldName
			fileInfo.FileSize = size
			file, err := s.UploadFile(model.UploadFileParam{
				Ctx:          ctx,
				SystemConfig: systemConfig,
				OldName:      oldName,
				NewFileName:  newFileName,
				Data:         multipartFile.Data,
				FileSort:     fileSort,
				File:         fileInfo,
			})
			if err != nil {
				return result, err
			}
			result = append(result, file)
		}
		//保存成功返回数据
		return
	}
	return result, gerror.New("请上传图片")
}

func (s *sFile) UploadFile(param model.UploadFileParam) (result *model.File, err error) {
	//checkMd5
	md5 := gmd5.MustEncryptBytes(param.Data)
	result = s.getFileMd5(param.Ctx, md5, param.SystemConfig)
	// 文件重复
	if result != nil {
		return
	}
	// 判断上传的格式是否合法
	if !utils.IsPic(gfile.ExtName(param.NewFileName)) {
		return result, gerror.New("请上传正确的图片")
	}

	//对图片大小进行限制
	if len(param.Data) > (10 * 1024 * 1024) {
		return result, gerror.New("图片大小不能超过10M")
	}

	result = param.File
	result.FileSortUid = param.FileSort.Uid
	result.FileMd5 = md5
	// 上传七牛云，判断是否能够上传七牛云
	if EOpenStatus.OPEN == param.SystemConfig.UploadQiNiu {
		url, err := service.CloudStorage(consts.QiNiuKey).UploadFile(param)
		utils.ErrIsNil(param.Ctx, err)
		result.QiNiuUrl = url
	}

	// 判断是否能够上传Minio文件服务器
	if EOpenStatus.OPEN == param.SystemConfig.UploadMinio {
		url, err := service.CloudStorage(consts.MinioKey).UploadFile(param)
		utils.ErrIsNil(param.Ctx, err)
		result.MinioUrl = url
	}
	if EOpenStatus.OPEN == param.SystemConfig.UploadLocal {
		url, err := service.CloudStorage(consts.LocalKey).UploadFile(param)
		utils.ErrIsNil(param.Ctx, err)
		result.PicUrl = url
	}
	s.setFilePriority(result, param.SystemConfig)
	_, err = dao.File.Ctx(param.Ctx).OmitEmpty().Save(result)
	return
}

func (s *sFile) UploadPicsByUrl(ctx context.Context, fileVO model.FileVO) (result []*model.File, err error) {
	fileList := make([]*model.UploadFileInfo, 0)
	for _, itemUrl := range fileVO.UrlList {
		bytes := g.Client().GetBytes(ctx, itemUrl)
		fileList = append(fileList, &model.UploadFileInfo{
			Data: bytes,
			Size: int64(len(bytes)),
			Name: gfile.Basename(itemUrl),
		})
	}
	// 获取配置文件
	var systemConfig baseModel.SystemConfig
	if fileVO.SystemConfig != nil {
		resultMap := fileVO.SystemConfig
		systemConfig, err = feign.GetSystemConfigByMap(resultMap)
	} else {
		// 从Redis中获取七牛云配置文件
		systemConfig, err = feign.GetSystemConfig(ctx)
	}
	request := g.RequestFromCtx(ctx)
	request.SetParam(SysConf.USER_UID, fileVO.UserUid)
	request.SetParam(SysConf.ADMIN_UID, fileVO.AdminUid)
	return s.BatchUploadFile(ctx, fileList, systemConfig)
}

func (s *sFile) CkeditorUploadFile(ctx context.Context) (result map[string]interface{}, err error) {
	request := g.RequestFromCtx(ctx)
	systemConfig, err := feign.GetSystemConfig(ctx)
	if err != nil {
		return nil, err
	}
	fileList := make([]*model.UploadFileInfo, 0)
	for fileName := range request.MultipartForm.File {
		file := request.GetUploadFile(fileName)
		//获取旧名称
		oldName := file.Filename
		fileReader, err := file.Open()
		if err != nil {
			return nil, err
		}
		bytes := make([]byte, file.Size)
		_, err = fileReader.Read(bytes)
		if err != nil {
			return nil, err
		}
		fileReader.Close()
		fileList = append(fileList, &model.UploadFileInfo{
			Data: bytes,
			Size: file.Size,
			Name: oldName,
		})
	}
	if len(fileList) > 0 {
		// 设置图片上传服务必要的信息
		request.SetParam(SysConf.USER_UID, SysConf.DEFAULT_UID)
		request.SetParam(SysConf.ADMIN_UID, SysConf.DEFAULT_UID)
		// 批量上传图片
		files, err := s.BatchUploadFile(ctx, fileList, systemConfig)
		if err != nil {
			return nil, err
		}
		if len(files) > 0 {
			picture := files[0]
			result = make(map[string]interface{})
			result[SysConf.UPLOADED] = 1
			result[SysConf.FILE_NAME] = picture.PicName
			result[SysConf.URL] = picture.PictureUrl
		}

	}

	return
}

func (s *sFile) CkeditorUploadCopyFile(ctx context.Context) (result map[string]interface{}, err error) {
	request := g.RequestFromCtx(ctx)
	token := request.GetParam(SysConf.TOKEN).String()
	if token == "" {
		return nil, gerror.New("未读取到携带token")
	}
	params := gstr.Split(token, "\\?url=")
	// 从Redis中获取系统配置文件
	systemConfig, err := feign.GetSystemConfig(ctx)
	var userUid = SysConf.DEFAULT_UID
	var adminUid = SysConf.DEFAULT_UID
	// 需要上传的URL
	var itemUrl = params[1]
	result = make(map[string]interface{})
	result[SysConf.UPLOADED] = 1
	result[SysConf.FILE_NAME] = itemUrl
	//判断需要上传的域名和本机图片域名是否一致
	if EFilePriority.QI_NIU == systemConfig.ContentPicturePriority {
		// 判断需要上传的域名和七牛域名是否一致，如果一致，那么就不需要重新上传，而是直接返回
		if systemConfig.QiNiuPictureBaseUrl != "" && itemUrl != "" && gstr.Contains(itemUrl, systemConfig.QiNiuPictureBaseUrl) {
			return
		}
	} else if EFilePriority.MINIO == systemConfig.ContentPicturePriority {
		// 判断需要上传的域名和minio域名是否一致，如果一致，那么就不需要重新上传，而是直接返回
		if systemConfig.MinioPictureBaseUrl != "" && itemUrl != "" && gstr.Contains(itemUrl, systemConfig.MinioPictureBaseUrl) {
			return
		}
	} else {
		// 判断需要上传的域名和minio域名是否一致，如果一致，那么就不需要重新上传，而是直接返回
		if systemConfig.LocalPictureBaseUrl != "" && itemUrl != "" && gstr.Contains(itemUrl, systemConfig.LocalPictureBaseUrl) {
			return
		}
	}
	fileSort, err := s.getFileSort(ctx, SysConf.ADMIN, SysConf.PROJECT_NAME)
	utils.ErrIsNil(ctx, err)
	if fileSort == nil {
		return result, gerror.New("文件不被允许上传")
	}
	//获取新文件名(默认为jpg)
	newFileName := fmt.Sprintf("%d.%s", gtime.Now().UnixMilli(), Constants.FILE_SUFFIX_JPG)
	bytes := g.Client().GetBytes(ctx, itemUrl)
	file := &model.File{
		File: entity.File{
			FileOldName:     itemUrl,
			PicName:         newFileName,
			PicExpandedName: Constants.FILE_SUFFIX_JPG,
			FileSize:        int64(len(bytes)),
			AdminUid:        adminUid,
			UserUid:         userUid,
		},
	}
	file, err = s.UploadFile(model.UploadFileParam{
		Ctx:          ctx,
		SystemConfig: systemConfig,
		OldName:      itemUrl,
		NewFileName:  newFileName,
		Data:         bytes,
		FileSort:     fileSort,
		File:         file,
	})
	if err != nil {
		return nil, err
	}
	result[SysConf.FILE_NAME] = newFileName
	result[SysConf.URL] = file.PictureUrl
	return
}

// CkeditorUploadToolFile 工具栏的文件上传
func (s *sFile) CkeditorUploadToolFile(ctx context.Context) (result map[string]interface{}, err error) {
	return s.CkeditorUploadFile(ctx)
}

func (s *sFile) getFileSort(ctx context.Context, sortName, projectName string) (fileSort *entity.FileSort, err error) {
	sortColumns := dao.FileSort.Columns()
	err = dao.FileSort.Ctx(ctx).
		Where(g.Map{
			sortColumns.SortName:    sortName,
			sortColumns.ProjectName: projectName,
			sortColumns.Status:      EStatus.ENABLE,
		}).Scan(&fileSort)
	utils.ErrIsNil(ctx, err)
	return
}

func (s *sFile) getFileMd5(ctx context.Context, md5 string, systemConfig baseModel.SystemConfig) (result *model.File) {
	// 未获取到m5d，直接返回空
	if md5 == "" || &systemConfig == nil {
		return nil
	}
	err := dao.File.Ctx(ctx).Where(do.File{
		FileMd5: md5,
		Status:  EStatus.ENABLE,
	}).Scan(&result)
	if err != nil {
		return nil
	}
	if result == nil {
		return
	}
	// 判断历史的文件，是否都满足上传要求
	// 1. 开启了上传本地，但是本地没有图片
	if EOpenStatus.OPEN == systemConfig.UploadLocal && result.PicUrl == "" {
		return nil
	}
	// 2. 开启了上传七牛云，但是七牛云没有图片
	if EOpenStatus.OPEN == systemConfig.UploadQiNiu && result.QiNiuUrl == "" {
		return nil
	}
	// 3. 开启了上传minio，但是minio没有图片
	if EOpenStatus.OPEN == systemConfig.UploadMinio && result.MinioUrl == "" {
		return nil
	}
	// 根据图片优先级设置图片
	s.setFilePriority(result, systemConfig)
	g.Log().Infof(ctx, "【命中重复的文件】md5: %s,fileUid: %s, url: %s", md5, result.Uid, result.PictureUrl)
	// 都满足条件，那么返回已存在的图片
	return
}

func (s *sFile) setFilePriority(file *model.File, systemConfig baseModel.SystemConfig) {
	if file == nil || &systemConfig == nil {
		return
	}
	picturePriority := systemConfig.PicturePriority
	if EFilePriority.QI_NIU == picturePriority {
		file.PictureUrl = systemConfig.QiNiuPictureBaseUrl + file.QiNiuUrl
	} else if EFilePriority.MINIO == picturePriority {
		file.PictureUrl = systemConfig.MinioPictureBaseUrl + file.MinioUrl
	} else if EFilePriority.LOCAL == picturePriority {
		file.PictureUrl = systemConfig.LocalPictureBaseUrl + file.PicUrl
	}
}
