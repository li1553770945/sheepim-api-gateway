package assembler

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/model/user"
	userRpc "github.com/li1553770945/sheepim-user-service/kitex_gen/user"
)

func UserReqHttpToRpc(httpReq *user.GetUserInfoReq) *userRpc.UserInfoReq {
	return &userRpc.UserInfoReq{
		UserId: httpReq.GetUserId(),
	}
}

func UserRespRpcToHttp(rpcResp *userRpc.UserInfoResp) *user.GetUserInfoResp {
	if rpcResp.BaseResp.Code != 0 {
		return &user.GetUserInfoResp{
			Code:    rpcResp.BaseResp.Code,
			Message: rpcResp.BaseResp.Message,
		}
	}
	return &user.GetUserInfoResp{
		Code: rpcResp.BaseResp.Code,
		Data: &user.GetUserInfoRespData{
			Username: *rpcResp.Username,
			Nickname: *rpcResp.Nickname,
			Role:     *rpcResp.Role,
		},
	}

}
