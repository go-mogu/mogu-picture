package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type PageReq struct {
	BeginTime *gtime.Time `json:"beginTime" dc:"开始时间" q:"-"`
	EndTime   *gtime.Time `json:"endTime" dc:"结束时间" q:"-"`
	PageNum   int         `json:"pageNum" d:"1" dc:"当前页码" q:"-"`
	PageSize  int         `json:"pageSize" d:"10" dc:"每页数" q:"-"`
	OrderBy   string      `json:"orderBy" dc:"排序方式" q:"-"`
}
