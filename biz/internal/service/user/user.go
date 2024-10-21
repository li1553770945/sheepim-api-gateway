package user

import "sheepim-api-gateway/biz/model/user"

func (s *UserServiceImpl) GetUserInfo(req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error) {
	userinfo = &user.GetUserInfoResp{
		Username: "test",
		Nickname: "test",
		Role:     "test",
	}
	return
}
