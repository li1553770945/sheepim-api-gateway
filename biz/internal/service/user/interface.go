package user

import (
	"context"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/user"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
)

type UserServiceImpl struct {
	UserRpcClient userservice.Client
}

type IUserService interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error)
}

func NewUserService(userRpcClient userservice.Client) IUserService {
	return &UserServiceImpl{
		UserRpcClient: userRpcClient,
	}
}
