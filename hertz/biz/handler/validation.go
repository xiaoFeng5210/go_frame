package handler

import (
	"errors"
	"fmt"
	"time"
)

// 自定义校验器
func BeforeToday(args ...interface{}) error {
	if len(args) != 1 {
		return fmt.Errorf("the args must be one")
	}
	if date, ok := args[0].(float64); ok { //通过反射获得结构体Field的值
		today := time.Now().Unix()
		if int64(date) < today {
			return nil
		} else {
			return errors.New("after today")
		}
	} else {
		return errors.New("data type must be int64")
	}
}

// 自定义 bind 和 validate 的 Error
type ValidateError struct {
	FailField, Msg string
}

func (e *ValidateError) Error() string {
	if e.Msg != "" {
		return "字段[" + e.FailField + "]校验失败[" + e.Msg + "]"
	}
	return "字段[" + e.FailField + "]校验失败"
}
