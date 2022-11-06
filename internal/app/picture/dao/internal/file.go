package internal
	
import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)
	
type FileDao struct {
	table   string                   // 表名称
	group   string                   // 数据源分组，默认default
	columns FileColumns // 表字段
	columnMap map[string]string //表字段map
}
	
// FileColumns defines and stores column names for table gen_database.
type FileColumns struct {
    Uid  string //唯一uid
    FileOldName  string //旧文件名
    PicName  string //文件名
    PicUrl  string //文件地址
    PicExpandedName  string //文件扩展名
    FileSize  string //文件大小
    FileSortUid  string //文件分类uid
    AdminUid  string //管理员uid
    UserUid  string //用户uid
    Status  string //状态
    CreateTime  string //创建时间
    UpdateTime  string //更新时间
    QiNiuUrl  string //七牛云地址
    MinioUrl  string //Minio文件URL
}
	
//  fileColumns holds the columns for table mes_gf_mst_yield_line.
var fileColumns = FileColumns{
    Uid:  "uid", //唯一uid
    FileOldName:  "file_old_name", //旧文件名
    PicName:  "pic_name", //文件名
    PicUrl:  "pic_url", //文件地址
    PicExpandedName:  "pic_expanded_name", //文件扩展名
    FileSize:  "file_size", //文件大小
    FileSortUid:  "file_sort_uid", //文件分类uid
    AdminUid:  "admin_uid", //管理员uid
    UserUid:  "user_uid", //用户uid
    Status:  "status", //状态
    CreateTime:  "create_time", //创建时间
    UpdateTime:  "update_time", //更新时间
    QiNiuUrl:  "qi_niu_url", //七牛云地址
    MinioUrl:  "minio_url", //Minio文件URL
}
var fileColumnMap = map[string]string{
    "Uid":  "uid", //唯一uid
    "FileOldName":  "file_old_name", //旧文件名
    "PicName":  "pic_name", //文件名
    "PicUrl":  "pic_url", //文件地址
    "PicExpandedName":  "pic_expanded_name", //文件扩展名
    "FileSize":  "file_size", //文件大小
    "FileSortUid":  "file_sort_uid", //文件分类uid
    "AdminUid":  "admin_uid", //管理员uid
    "UserUid":  "user_uid", //用户uid
    "Status":  "status", //状态
    "CreateTime":  "create_time", //创建时间
    "UpdateTime":  "update_time", //更新时间
    "QiNiuUrl":  "qi_niu_url", //七牛云地址
    "MinioUrl":  "minio_url", //Minio文件URL
}
// NewFileDao creates and returns a new DAO object for table data access.
func NewFileDao() *FileDao {
	return &FileDao{
		group:   "default",
		table:   "t_file",
		columns: fileColumns,
        columnMap: fileColumnMap,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FileDao) DB() gdb.DB {
	return g.DB(dao.group)
}
// Table returns the table name of current dao.
func (dao *FileDao) Table() string {
	return dao.table
}
// Columns returns all column names of current dao.
func (dao *FileDao) Columns() FileColumns {
	return dao.columns
}
// ColumnMap returns all column map of current dao.
func (dao *FileDao) ColumnMap() map[string]string {
	return dao.columnMap
}
// Group returns the configuration group name of database of current dao.
func (dao *FileDao) Group() string {
	return dao.group
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
