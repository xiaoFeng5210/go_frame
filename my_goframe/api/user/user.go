// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"my_goframe/api/user/v1"
)

type IUserV1 interface {
	CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (res *v1.CreateAccountRes, err error)
	DeleteAccount(ctx context.Context, req *v1.DeleteAccountReq) (res *v1.DeleteAccountRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
