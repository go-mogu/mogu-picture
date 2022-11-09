package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type StorageDao struct {
	table     string            // 表名称
	group     string            // 数据源分组，默认default
	columns   StorageColumns    // 表字段
	columnMap map[string]string //表字段map
}

// StorageColumns defines and stores column names for table gen_database.
type StorageColumns struct {
	Uid            string //唯一uid
	AdminUid       string //管理员uid
	StorageSize    string //网盘容量大小
	Status         string //状态
	CreateTime     string //创建时间
	UpdateTime     string //更新时间
	MaxStorageSize string //最大容量大小
}

//  storageColumns holds the columns for table mes_gf_mst_yield_line.
var storageColumns = StorageColumns{
	Uid:            "uid",              //唯一uid
	AdminUid:       "admin_uid",        //管理员uid
	StorageSize:    "storage_size",     //网盘容量大小
	Status:         "status",           //状态
	CreateTime:     "create_time",      //创建时间
	UpdateTime:     "update_time",      //更新时间
	MaxStorageSize: "max_storage_size", //最大容量大小
}
var storageColumnMap = map[string]string{
	"Uid":            "uid",              //唯一uid
	"AdminUid":       "admin_uid",        //管理员uid
	"StorageSize":    "storage_size",     //网盘容量大小
	"Status":         "status",           //状态
	"CreateTime":     "create_time",      //创建时间
	"UpdateTime":     "update_time",      //更新时间
	"MaxStorageSize": "max_storage_size", //最大容量大小
}

// NewStorageDao creates and returns a new DAO object for table data access.
func NewStorageDao() *StorageDao {
	return &StorageDao{
		group:     "default",
		table:     "t_storage",
		columns:   storageColumns,
		columnMap: storageColumnMap,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StorageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StorageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StorageDao) Columns() StorageColumns {
	return dao.columns
}

// ColumnMap returns all column map of current dao.
func (dao *StorageDao) ColumnMap() map[string]string {
	return dao.columnMap
}

// Group returns the configuration group name of database of current dao.
func (dao *StorageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StorageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StorageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
