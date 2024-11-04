package assembler

import (
	"sheepim-api-gateway/biz/model/auth"
	authservice "sheepim-auth-service/kitex_gen/auth"
)

func LoginReqHttpToRpc(httpReq *auth.LoginReq) *authservice.LoginReq {
	return &authservice.LoginReq{
		Username: httpReq.Username,
		Password: httpReq.Password,
	}
}

func LoginRespRpcToHttp(rpcResp *authservice.LoginResp) *auth.LoginResp {
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
