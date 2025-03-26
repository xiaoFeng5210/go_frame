package user

import (
	"context"

	v1 "my_goframe/api/user/v1"
	"my_goframe/internal/dao"
	"my_goframe/internal/model/do"
)

func (c *ControllerV1) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (res *v1.CreateAccountRes, err error) {
	_, err = dao.Login.Ctx(ctx).Data(do.Login{
		Username: req.Name,
		Password: req.Password,
	}).Insert()
	return
}
