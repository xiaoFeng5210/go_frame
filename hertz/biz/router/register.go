// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	math_service "dqq/go/math/biz/router/math_service"
	student_service "dqq/go/math/biz/router/student_service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	student_service.Register(r)

	math_service.Register(r)
}
