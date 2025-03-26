package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateAccountReq struct {
	g.Meta   `path:"/account" tags:"Login" method:"post" summary:"创建登录账户"`
	Name     string `p:"name" v:"required|length:1,10" dc:"用户名"`
	Password string `p:"password" v:"required|length:32,32" dc:"md5之后的密码"`
}

type CreateAccountRes struct {
}

type DeleteAccountReq struct {
	g.Meta `path:"/account/{name}" tags:"Login" method:"delete" summary:"删除账户"`
	Name   string `p:"name" v:"required|length:1,10" dc:"用户名"`
}

type DeleteAccountRes struct {
}

type LoginReq struct {
	g.Meta   `path:"/account/login" tags:"Login" method:"post" summary:"检查用户登录是否成功"`
	Name     string `p:"name" v:"required|length:1,10" dc:"用户名"`
	Password string `p:"password" v:"required|length:32,32" dc:"md5之后的密码"`
}

type LoginRes struct {
	Ok bool `json:"ok" dc:"是否登录成功"`
}
