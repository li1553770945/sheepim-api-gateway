package user

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"sheepim-api-gateway/biz/internal/assembler"
	"sheepim-api-gateway/biz/model/user"
)

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error) {
	rpcReq := assembler.UserReqHttpToRpc(req)
	rpcResp, err := s.UserRpcClient.GetUserInfo(context.Background(), rpcReq)
	if err != nil {
		klog.CtxErrorf(ctx, "调用用户服务失败:"+err.Error())
		return
	}
	return assembler.UserRespRpcToHttp(rpcResp), nil
}
