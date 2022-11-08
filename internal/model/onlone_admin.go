package model

type OnlineAdmin struct {
	Token         string `json:"token"`         //用户Token
	TokenId       string `json:"tokenId"`       //会话编号
	AdminUid      string `json:"adminUid"`      //管理员的UID
	UserName      string `json:"userName"`      //用户名称
	RoleName      string `json:"roleName"`      //角色名称
	Ipaddr        string `json:"ipaddr"`        //登录IP地址
	LoginLocation string `json:"loginLocation"` //登录地址
	LoginTime     string `json:"loginTime"`     //登录时间
	ExpireTime    string `json:"expireTime"`    // 过期时间
	Browser       string `json:"browser"`       //浏览器类型
	OS            string `json:"os"`            //操作系统

}
