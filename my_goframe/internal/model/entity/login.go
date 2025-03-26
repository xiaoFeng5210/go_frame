// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Login is the golang structure for table login.
type Login struct {
	Username string `json:"username" orm:"username" ` // 用户名
	Password string `json:"password" orm:"password" ` // md5之后的密码
}
