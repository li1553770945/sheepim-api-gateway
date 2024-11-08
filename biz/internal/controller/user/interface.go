package user

import (
	"context"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/user"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
)

type UserControllerImpl struct {
	UserRpcClient userservice.Client
}

type IUserController interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp)
}

func NewUserController(userRpcClient userservice.Client) IUserController {
	return &UserControllerImpl{
		UserRpcClient: userRpcClient,
	}
}
