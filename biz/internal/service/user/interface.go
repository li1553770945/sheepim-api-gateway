package user

import (
	"sheepim-api-gateway/biz/model/user"
	"sheepim-user-service/kitex_gen/user/userservice"
)

type UserServiceImpl struct {
	UserRpcClient userservice.Client
}

type IUserService interface {
	GetUserInfo(req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error)
}

func NewUserService(userRpcClient userservice.Client) IUserService {
	return &UserServiceImpl{
		UserRpcClient: userRpcClient,
	}
}
