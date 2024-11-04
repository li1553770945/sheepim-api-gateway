package auth

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"sheepim-api-gateway/biz/constant"
	"sheepim-api-gateway/biz/internal/assembler"
	"sheepim-api-gateway/biz/model/auth"
)

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	rpcReq := assembler.LoginReqHttpToRpc(req)
	rpcResp, err := s.AuthRpcClient.Login(context.Background(), rpcReq)
	if err != nil {
		klog.CtxErrorf(ctx, "调用认证服务失败:"+err.Error())
		return &auth.LoginResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用认证服务失败",
		}, nil

	}
	return assembler.LoginRespRpcToHttp(rpcResp), nil
}
