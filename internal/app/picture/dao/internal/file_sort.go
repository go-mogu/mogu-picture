package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type FileSortDao struct {
	table     string            // 表名称
	group     string            // 数据源分组，默认default
	columns   FileSortColumns   // 表字段
	columnMap map[string]string //表字段map
}

// FileSortColumns defines and stores column names for table gen_database.
type FileSortColumns struct {
	Uid         string //唯一uid
	ProjectName string //项目名
	SortName    string //分类名
	Url         string //分类路径
	Status      string //状态
	CreateTime  string //创建时间
	UpdateTime  string //更新时间
}

// fileSortColumns holds the columns for table mes_gf_mst_yield_line.
var fileSortColumns = FileSortColumns{
	Uid:         "uid",          //唯一uid
	ProjectName: "project_name", //项目名
	SortName:    "sort_name",    //分类名
	Url:         "url",          //分类路径
	Status:      "status",       //状态
	CreateTime:  "create_time",  //创建时间
	UpdateTime:  "update_time",  //更新时间
}
var fileSortColumnMap = map[string]string{
	"Uid":         "uid",          //唯一uid
	"ProjectName": "project_name", //项目名
	"SortName":    "sort_name",    //分类名
	"Url":         "url",          //分类路径
	"Status":      "status",       //状态
	"CreateTime":  "create_time",  //创建时间
	"UpdateTime":  "update_time",  //更新时间
}

// NewFileSortDao creates and returns a new DAO object for table data access.
func NewFileSortDao() *FileSortDao {
	return &FileSortDao{
		group:     "default",
		table:     "t_file_sort",
		columns:   fileSortColumns,
		columnMap: fileSortColumnMap,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FileSortDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FileSortDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FileSortDao) Columns() FileSortColumns {
	return dao.columns
}

// ColumnMap returns all column map of current dao.
func (dao *FileSortDao) ColumnMap() map[string]string {
	return dao.columnMap
}

// Group returns the configuration group name of database of current dao.
func (dao *FileSortDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FileSortDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FileSortDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
