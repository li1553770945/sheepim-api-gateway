package assembler

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/model/auth"
	authRpc "github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
)

func LoginReqHttpToRpc(httpReq *auth.LoginReq) *authRpc.LoginReq {
	return &authRpc.LoginReq{
		Username: httpReq.Username,
		Password: httpReq.Password,
	}
}

func LoginRespRpcToHttp(rpcResp *authRpc.LoginResp) *auth.LoginResp {
	if rpcResp.BaseResp.Code != 0 {
		return &auth.LoginResp{
			Code:    rpcResp.BaseResp.Code,
			Message: rpcResp.BaseResp.Message,
		}
	}
	return &auth.LoginResp{
		Code: rpcResp.BaseResp.Code,
		Data: &auth.LoginRespData{
			Token: *rpcResp.Token,
		},
	}

}
