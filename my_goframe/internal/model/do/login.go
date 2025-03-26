// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Login is the golang structure of table login for DAO operations like Where/Data.
type Login struct {
	g.Meta   `orm:"table:login, do:true"`
	Username interface{} // 用户名
	Password interface{} // md5之后的密码
}
