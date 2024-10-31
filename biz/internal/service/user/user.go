package user

import (
	"context"
	"sheepim-api-gateway/biz/internal/assembler"
	"sheepim-api-gateway/biz/model/user"
)

func (s *UserServiceImpl) GetUserInfo(req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error) {
	rpcReq := assembler.UserReqHttpToRpc(req)
	rpcResp, err := s.UserRpcClient.GetUserInfo(context.Background(), rpcReq)
	return assembler.UserRespRpcToHttp(rpcResp), nil
}
