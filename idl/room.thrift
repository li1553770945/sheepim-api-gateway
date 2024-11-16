namespace go room

struct CreateRoomRespData{
    1: required string roomId
    2: required string clientId
    3: required string clientToken
}
struct CreateRoomResp{
    1: required i32 code
    2: required string message
    3: optional CreateRoomRespData data
}

struct JoinRoomReq{
    1: required string roomId (api.query="roomId");
}

struct JoinRoomRespData{
    1: required string clientId
    2: required string clientToken
}

struct JoinRoomResp{
    1: required i32 code
    2: required string message
    3: optional JoinRoomRespData data

}
service RoomService {
    CreateRoomResp CreateRoom()(api.post="/api/rooms")
    JoinRoomResp JoinRoom(JoinRoomReq req)  (api.post="/api/rooms/join")
}
