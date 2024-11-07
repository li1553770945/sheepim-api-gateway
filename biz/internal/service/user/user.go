package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/user"
)

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (userinfo *user.GetUserInfoResp, err error) {
	hlog.CtxInfof(ctx, "收到获取用户 id 为 %d 的用户信息请求", req.UserId)
	rpcReq := assembler.UserReqHttpToRpc(req)

	hlog.CtxInfof(ctx, "调用用户服务获取用户ID 为的%d用户信息", req.GetUserId())
	rpcResp, err := s.UserRpcClient.GetUserInfo(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用用户服务失败:"+err.Error())
		return &user.GetUserInfoResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用用户服务失败",
		}, nil
	}
	hlog.CtxInfof(ctx, "调用用户服务成功，获取用户ID为的%d的用户信息，服务返回状态码:%d", req.GetUserId(), rpcResp.BaseResp.Code)
	return assembler.UserRespRpcToHttp(rpcResp), nil
}
