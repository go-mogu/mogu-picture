package model

type PageInfo struct {
	CurrentPage int    `json:"currentPage" d:"1" dc:"当前页码" q:"-"`
	PageSize    int    `json:"pageSize" d:"10" dc:"每页数" q:"-"`
	Keyword     string `json:"keyword" dc:"排序方式" q:"-"`
}

type BaseVO struct {
	PageInfo
	Uid    string `json:"uid" d:"10" dc:"uid" q:"EQ"`
	Status int    `json:"status" dc:"状态" q:"-"`
}
