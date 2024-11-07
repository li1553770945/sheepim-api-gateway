package assembler

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/model/user"
	userservice "github.com/li1553770945/sheepim-user-service/kitex_gen/user"
)

func UserReqHttpToRpc(httpReq *user.GetUserInfoReq) *userservice.UserInfoReq {
	return &userservice.UserInfoReq{
		UserId: httpReq.GetUserId(),
	}
}

func UserRespRpcToHttp(rpcResp *userservice.UserInfoResp) *user.GetUserInfoResp {
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
