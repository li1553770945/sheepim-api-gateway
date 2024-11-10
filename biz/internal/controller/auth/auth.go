package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/auth"
)

func (c *AuthControllerImpl) Login(ctx context.Context, req *auth.LoginReq) *auth.LoginResp {
	hlog.CtxInfof(ctx, "收到用户 %c 的登录请求", req.Username)
	rpcReq := assembler.LoginReqHttpToRpc(req)

	hlog.CtxInfof(ctx, "用户 %c 调用认证服务进行登录验证", req.Username)
	rpcResp, err := c.AuthRpcClient.Login(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用认证服务失败:"+err.Error())
		return &auth.LoginResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用认证服务失败",
		}

	}
	hlog.CtxInfof(ctx, "用户 %c 登录认证服务返回状态：%d", req.Username, rpcResp.BaseResp.Code)
	return assembler.LoginRespRpcToHttp(rpcResp)
}
