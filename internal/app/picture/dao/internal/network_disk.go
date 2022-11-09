package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type NetworkDiskDao struct {
	table     string             // 表名称
	group     string             // 数据源分组，默认default
	columns   NetworkDiskColumns // 表字段
	columnMap map[string]string  //表字段map
}

// NetworkDiskColumns defines and stores column names for table gen_database.
type NetworkDiskColumns struct {
	Uid         string //唯一uid
	AdminUid    string //管理员uid
	ExtendName  string //扩展名
	FileName    string //文件名
	FilePath    string //文件路径
	FileSize    string //文件大小
	IsDir       string //是否目录
	Status      string //状态
	CreateTime  string //创建时间
	UpdateTime  string //更新时间
	LocalUrl    string //本地文件URL
	QiNiuUrl    string //七牛云URL
	FileOldName string //上传前文件名
	MinioUrl    string //Minio文件URL
}

//  networkDiskColumns holds the columns for table mes_gf_mst_yield_line.
var networkDiskColumns = NetworkDiskColumns{
	Uid:         "uid",           //唯一uid
	AdminUid:    "admin_uid",     //管理员uid
	ExtendName:  "extend_name",   //扩展名
	FileName:    "file_name",     //文件名
	FilePath:    "file_path",     //文件路径
	FileSize:    "file_size",     //文件大小
	IsDir:       "is_dir",        //是否目录
	Status:      "status",        //状态
	CreateTime:  "create_time",   //创建时间
	UpdateTime:  "update_time",   //更新时间
	LocalUrl:    "local_url",     //本地文件URL
	QiNiuUrl:    "qi_niu_url",    //七牛云URL
	FileOldName: "file_old_name", //上传前文件名
	MinioUrl:    "minio_url",     //Minio文件URL
}
var networkDiskColumnMap = map[string]string{
	"Uid":         "uid",           //唯一uid
	"AdminUid":    "admin_uid",     //管理员uid
	"ExtendName":  "extend_name",   //扩展名
	"FileName":    "file_name",     //文件名
	"FilePath":    "file_path",     //文件路径
	"FileSize":    "file_size",     //文件大小
	"IsDir":       "is_dir",        //是否目录
	"Status":      "status",        //状态
	"CreateTime":  "create_time",   //创建时间
	"UpdateTime":  "update_time",   //更新时间
	"LocalUrl":    "local_url",     //本地文件URL
	"QiNiuUrl":    "qi_niu_url",    //七牛云URL
	"FileOldName": "file_old_name", //上传前文件名
	"MinioUrl":    "minio_url",     //Minio文件URL
}

// NewNetworkDiskDao creates and returns a new DAO object for table data access.
func NewNetworkDiskDao() *NetworkDiskDao {
	return &NetworkDiskDao{
		group:     "default",
		table:     "t_network_disk",
		columns:   networkDiskColumns,
		columnMap: networkDiskColumnMap,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetworkDiskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetworkDiskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetworkDiskDao) Columns() NetworkDiskColumns {
	return dao.columns
}

// ColumnMap returns all column map of current dao.
func (dao *NetworkDiskDao) ColumnMap() map[string]string {
	return dao.columnMap
}

// Group returns the configuration group name of database of current dao.
func (dao *NetworkDiskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetworkDiskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetworkDiskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
