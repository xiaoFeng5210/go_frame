package user

import (
	"context"

	v1 "my_goframe/api/user/v1"
	"my_goframe/internal/dao"
)

func (c *ControllerV1) DeleteAccount(ctx context.Context, req *v1.DeleteAccountReq) (res *v1.DeleteAccountRes, err error) {
	_, err = dao.Login.Ctx(ctx).Where("username=?", req.Name).Delete()
	return
}
