package user

import "sheepim-api-gateway/biz/model/user"

type UserServiceImpl struct {
}

type IUserService interface {
	GetUserInfo(req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error)
}

func NewUserService() IUserService {
	return &UserServiceImpl{}
}
