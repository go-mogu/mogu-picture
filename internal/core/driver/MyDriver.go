package driver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-mogu/mogu-picture/internal/consts/EStatus"
	"github.com/go-mogu/mogu-picture/internal/consts/SQLConf"
	"github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

// MyDriver is a custom database driver, which is used for testing only.
// For simplifying the unit testing case purpose, MyDriver struct inherits the mysql driver
// gdb.Driver and overwrites its functions DoQuery and DoExec.
// So if there's any sql execution, it goes through MyDriver.DoQuery/MyDriver.DoExec firstly
// and then gdb.Driver.DoQuery/gdb.Driver.DoExec.
// You can call it sql "HOOK" or "HiJack" as your will.
type MyDriver struct {
	*mysql.Driver
}

var (
	// customDriverName is my driver name, which is used for registering.
	customDriverName = "mysql"
)

func init() {
	// It here registers my custom driver in package initialization function "init".
	// You can later use this type in the database configuration.
	if err := gdb.Register(customDriverName, &MyDriver{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *MyDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &MyDriver{
		&mysql.Driver{
			Core: core,
		},
	}, nil
}

func (d *MyDriver) DoInsert(ctx context.Context, link gdb.Link, table string, data gdb.List, option gdb.DoInsertOption) (result sql.Result, err error) {
	now := gtime.Now()
	for _, item := range data {
		item[SQLConf.UID] = idgen.NextId()
		item[SQLConf.CREATE_TIME] = now
		item[SQLConf.UPDATE_TIME] = now
		item[SQLConf.STATUS] = EStatus.ENABLE
	}
	return d.Driver.DoInsert(ctx, link, table, data, option)
}

func (d *MyDriver) DoUpdate(ctx context.Context, link gdb.Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) {
	fmt.Println(data)
	return d.Driver.DoUpdate(ctx, link, table, data, condition, args)
}
