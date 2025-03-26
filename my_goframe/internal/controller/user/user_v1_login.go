package user

import (
	"context"

	v1 "my_goframe/api/user/v1"
	"my_goframe/internal/dao"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	ok := dao.CheckLogin(req.Name, req.Password)
	res = &v1.LoginRes{Ok: ok}
	return
}
