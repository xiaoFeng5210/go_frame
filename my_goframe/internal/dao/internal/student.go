// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StudentDao is the data access object for the table student.
type StudentDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns StudentColumns // columns contains all the column names of Table for convenient usage.
}

// StudentColumns defines and stores column names for the table student.
type StudentColumns struct {
	Id         string // 主键自增id
	Name       string // 姓名
	Province   string // 省
	City       string // 城市
	Addr       string // 地址
	Score      string // 考试成绩
	Enrollment string // 入学时间
}

// studentColumns holds the columns for the table student.
var studentColumns = StudentColumns{
	Id:         "id",
	Name:       "name",
	Province:   "province",
	City:       "city",
	Addr:       "addr",
	Score:      "score",
	Enrollment: "enrollment",
}

// NewStudentDao creates and returns a new DAO object for table data access.
func NewStudentDao() *StudentDao {
	return &StudentDao{
		group:   "default",
		table:   "student",
		columns: studentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StudentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StudentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StudentDao) Columns() StudentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StudentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StudentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StudentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
