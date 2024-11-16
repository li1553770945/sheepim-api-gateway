package room

import (
	"context"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/room"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/room/roomservice"
)

// RoomControllerImpl 实现了 IRoomController 接口
type RoomControllerImpl struct {
	RoomRpcClient roomservice.Client
}

// IRoomController 定义了 Room 的接口
type IRoomController interface {
	CreateRoom(ctx context.Context) *room.CreateRoomResp
	JoinRoom(ctx context.Context, req *room.JoinRoomReq) *room.JoinRoomResp
}

// NewRoomController 构造函数，用于创建 IRoomController 的实现
func NewRoomController(roomRpcClient roomservice.Client) IRoomController {
	return &RoomControllerImpl{
		RoomRpcClient: roomRpcClient,
	}
}
