// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         int         `json:"id"         orm:"id"          ` // 用户id，自增
	Name       string      `json:"name"       orm:"name"        ` // 用户名
	Password   string      `json:"password"   orm:"password"    ` // 密码的md5
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 用户注册时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 最后修改时间
}
