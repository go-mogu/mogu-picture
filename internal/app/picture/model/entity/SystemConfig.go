package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type SystemConfig struct {
	Uid                       string      `json:"uid" q:"EQ" dc:"主键"`                                                     //主键
	QiNiuAccessKey            string      `json:"qiNiuAccessKey" q:"EQ" dc:"七牛云公钥"`                                       //七牛云公钥
	QiNiuSecretKey            string      `json:"qiNiuSecretKey" q:"EQ" dc:"七牛云私钥"`                                       //七牛云私钥
	Email                     string      `json:"email" q:"EQ" dc:"邮箱账号"`                                                 //邮箱账号
	EmailUserName             string      `json:"emailUserName" q:"LIKE" dc:"邮箱发件人用户名"`                                   //邮箱发件人用户名
	EmailPassword             string      `json:"emailPassword" q:"EQ" dc:"邮箱密码"`                                         //邮箱密码
	SmtpAddress               string      `json:"smtpAddress" q:"EQ" dc:"SMTP地址"`                                         //SMTP地址
	SmtpPort                  string      `json:"smtpPort" q:"EQ" dc:"SMTP端口"`                                            //SMTP端口
	Status                    int8        `json:"status" q:"EQ" dc:"状态"`                                                  //状态
	CreateTime                *gtime.Time `json:"createTime" q:"BETWEEN" dc:"创建时间"`                                       //创建时间
	UpdateTime                *gtime.Time `json:"updateTime" q:"BETWEEN" dc:"更新时间"`                                       //更新时间
	QiNiuBucket               string      `json:"qiNiuBucket" q:"EQ" dc:"七牛云上传空间"`                                        //七牛云上传空间
	QiNiuArea                 string      `json:"qiNiuArea" q:"EQ" dc:"七牛云存储区域 华东（z0），华北(z1)，华南(z2)，北美(na0)，东南亚(as0)"`    //七牛云存储区域 华东（z0），华北(z1)，华南(z2)，北美(na0)，东南亚(as0)
	UploadQiNiu               string      `json:"uploadQiNiu" q:"EQ" dc:"图片是否上传七牛云 (0:否， 1：是)"`                           //图片是否上传七牛云 (0:否， 1：是)
	UploadLocal               string      `json:"uploadLocal" q:"EQ" dc:"图片是否上传本地存储 (0:否， 1：是)"`                          //图片是否上传本地存储 (0:否， 1：是)
	PicturePriority           string      `json:"picturePriority" q:"EQ" dc:"图片显示优先级（ 1 展示 七牛云,  0 本地）"`                  //图片显示优先级（ 1 展示 七牛云,  0 本地）
	QiNiuPictureBaseUrl       string      `json:"qiNiuPictureBaseUrl" q:"EQ" dc:"七牛云域名前缀：http://images.moguit.cn"`        //七牛云域名前缀：http://images.moguit.cn
	LocalPictureBaseUrl       string      `json:"localPictureBaseUrl" q:"EQ" dc:"本地服务器域名前缀：http://localhost:8600"`        //本地服务器域名前缀：http://localhost:8600
	StartEmailNotification    string      `json:"startEmailNotification" q:"EQ" dc:"是否开启邮件通知(0:否， 1:是)"`                  //是否开启邮件通知(0:否， 1:是)
	EditorModel               int8        `json:"editorModel" q:"EQ" dc:"编辑器模式，(0：富文本编辑器CKEditor，1：markdown编辑器Veditor)"`  //编辑器模式，(0：富文本编辑器CKEditor，1：markdown编辑器Veditor)
	ThemeColor                string      `json:"themeColor" q:"EQ" dc:"主题颜色"`                                            //主题颜色
	MinioEndPoint             string      `json:"minioEndPoint" q:"EQ" dc:"Minio远程连接地址"`                                  //Minio远程连接地址
	MinioAccessKey            string      `json:"minioAccessKey" q:"EQ" dc:"Minio公钥"`                                     //Minio公钥
	MinioSecretKey            string      `json:"minioSecretKey" q:"EQ" dc:"Minio私钥"`                                     //Minio私钥
	MinioBucket               string      `json:"minioBucket" q:"EQ" dc:"Minio桶"`                                         //Minio桶
	UploadMinio               int8        `json:"uploadMinio" q:"EQ" dc:"图片是否上传Minio (0:否， 1：是)"`                         //图片是否上传Minio (0:否， 1：是)
	MinioPictureBaseUrl       string      `json:"minioPictureBaseUrl" q:"EQ" dc:"Minio服务器文件域名前缀"`                         //Minio服务器文件域名前缀
	OpenDashboardNotification int8        `json:"openDashboardNotification" q:"EQ" dc:"是否开启仪表盘通知(0:否， 1:是)"`              //是否开启仪表盘通知(0:否， 1:是)
	DashboardNotification     string      `json:"dashboardNotification" q:"EQ" dc:"仪表盘通知【用于首次登录弹框】"`                      //仪表盘通知【用于首次登录弹框】
	ContentPicturePriority    int8        `json:"contentPicturePriority" q:"EQ" dc:"博客详情图片显示优先级（ 0:本地  1: 七牛云 2: Minio）"` //博客详情图片显示优先级（ 0:本地  1: 七牛云 2: Minio）
	OpenEmailActivate         int8        `json:"openEmailActivate" q:"EQ" dc:"是否开启用户邮件激活功能【0 关闭，1 开启】"`                  //是否开启用户邮件激活功能【0 关闭，1 开启】
	SearchModel               int8        `json:"searchModel" q:"EQ" dc:"搜索模式【0:SQL搜索 、1：全文检索】"`                          //搜索模式【0:SQL搜索 、1：全文检索】
}
