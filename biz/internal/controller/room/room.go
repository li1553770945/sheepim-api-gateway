package room

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/room"
)

func (c *RoomControllerImpl) CreateRoom(ctx context.Context) *room.CreateRoomResp {
	hlog.CtxInfof(ctx, "调用房间服务创建房间")

	// 构造 RPC 请求（这里可能不需要额外参数，如果有可调整）

	// 调用 RPC 服务
	rpcResp, err := c.RoomRpcClient.CreateRoom(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用房间服务失败: %s", err.Error())
		return &room.CreateRoomResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用房间服务失败",
		}
	}

	hlog.CtxInfof(ctx, "调用房间服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	// 转换 RPC 响应为 HTTP 响应
	return assembler.CreateRoomRespRpcToHttp(rpcResp)
}

func (c *RoomControllerImpl) JoinRoom(ctx context.Context, req *room.JoinRoomReq) *room.JoinRoomResp {
	hlog.CtxInfof(ctx, "调用房间服务加入房间")

	// 转换 HTTP 请求为 RPC 请求
	rpcReq := assembler.JoinRoomReqHttpToRpc(req)

	// 调用 RPC 服务
	rpcResp, err := c.RoomRpcClient.JoinRoom(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用房间服务失败: %s", err.Error())
		return &room.JoinRoomResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用房间服务失败",
		}

	}

	hlog.CtxInfof(ctx, "调用房间服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	// 转换 RPC 响应为 HTTP 响应
	return assembler.JoinRoomRespRpcToHttp(rpcResp)
}
