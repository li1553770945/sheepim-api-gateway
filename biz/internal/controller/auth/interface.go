package auth

import (
	"context"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/auth"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
)

type AuthControllerImpl struct {
	AuthRpcClient authservice.Client
}

type IAuthController interface {
	Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp)
}

func NewAuthController(authRpcClient authservice.Client) IAuthController {
	return &AuthControllerImpl{
		AuthRpcClient: authRpcClient,
	}
}
