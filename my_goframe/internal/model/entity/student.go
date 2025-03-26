// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Student is the golang structure for table student.
type Student struct {
	Id         int         `json:"id"         orm:"id"         description:"主键自增id"` // 主键自增id
	Name       string      `json:"name"       orm:"name"       description:"姓名"`     // 姓名
	Province   string      `json:"province"   orm:"province"   description:"省"`      // 省
	City       string      `json:"city"       orm:"city"       description:"城市"`     // 城市
	Addr       string      `json:"addr"       orm:"addr"       description:"地址"`     // 地址
	Score      float64     `json:"score"      orm:"score"      description:"考试成绩"`   // 考试成绩
	Enrollment *gtime.Time `json:"enrollment" orm:"enrollment" description:"入学时间"`   // 入学时间
}
