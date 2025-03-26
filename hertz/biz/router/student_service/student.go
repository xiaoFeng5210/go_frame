// Code generated by hertz generator. DO NOT EDIT.

package student_service

import (
	student_service "dqq/go/math/biz/handler/student_service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/student", append(_queryMw(), student_service.Query)...)
	_student := root.Group("/student", _studentMw()...)
	_student.POST("/cookie", append(_cookieMw(), student_service.Cookie)...)
	_student.POST("/form", append(_postformMw(), student_service.PostForm)...)
	_student.POST("/header", append(_headerMw(), student_service.Header)...)
	_student.POST("/json", append(_postjsonMw(), student_service.PostJson)...)
	_student.POST("/pb", append(_postpbMw(), student_service.PostPb)...)
	{
		_name := _student.Group("/:name", _nameMw()...)
		_name.POST("/*addr", append(_restfulMw(), student_service.Restful)...)
	}
}
