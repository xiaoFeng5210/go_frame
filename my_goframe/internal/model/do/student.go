// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Student is the golang structure of table student for DAO operations like Where/Data.
type Student struct {
	g.Meta     `orm:"table:student, do:true"`
	Id         interface{} // 主键自增id
	Name       interface{} // 姓名
	Province   interface{} // 省
	City       interface{} // 城市
	Addr       interface{} // 地址
	Score      interface{} // 考试成绩
	Enrollment *gtime.Time // 入学时间
}
