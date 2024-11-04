package auth

import (
	"context"
	"sheepim-api-gateway/biz/model/auth"
	"sheepim-auth-service/kitex_gen/auth/authservice"
)

type AuthServiceImpl struct {
	AuthRpcClient authservice.Client
}

type IAuthService interface {
	Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp, err error)
}

func NewUserService(authRpcClient authservice.Client) IAuthService {
	return &AuthServiceImpl{
		AuthRpcClient: authRpcClient,
	}
}
