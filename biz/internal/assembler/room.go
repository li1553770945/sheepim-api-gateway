package assembler

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/room"
	roomRpc "github.com/li1553770945/sheepim-room-service/kitex_gen/room"
)

// HTTP -> RPC 转换
func JoinRoomReqHttpToRpc(req *room.JoinRoomReq) *roomRpc.JoinRoomReq {
	return &roomRpc.JoinRoomReq{
		RoomId: req.RoomId,
	}
}

// RPC -> HTTP 转换
func JoinRoomRespRpcToHttp(resp *roomRpc.JoinRoomResp) *room.JoinRoomResp {
	if resp.BaseResp.Code != constant.Success {
		return &room.JoinRoomResp{
			Code:    resp.BaseResp.Code,
			Message: resp.BaseResp.Message,
		}
	}

	return &room.JoinRoomResp{
		Code: resp.BaseResp.Code,
		Data: &room.JoinRoomRespData{
			ClientId:    *resp.ClientId,
			ClientToken: *resp.ClientToken,
		},
	}
}

// RPC -> HTTP 转换
func CreateRoomRespRpcToHttp(resp *roomRpc.CreateRoomResp) *room.CreateRoomResp {
	return &room.CreateRoomResp{
		Code:    resp.BaseResp.Code,
		Message: resp.BaseResp.Message,
		Data: &room.CreateRoomRespData{
			RoomId:      *resp.RoomId,
			ClientId:    *resp.ClientId,
			ClientToken: *resp.ClientToken,
		},
	}
}
