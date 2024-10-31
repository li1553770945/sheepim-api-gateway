namespace go user

struct GetUserInfoReq {
    1: required i64 userId (api.query="userId"); // 添加 api 注解为方便进行参数绑定
}

struct GetUserInfoRespData {
    1: string username;
    2: string nickname;
    3: string role;
}

struct GetUserInfoResp{
    1: required i32 code;
    2: optional string message;
    3: optional GetUserInfoRespData data;
}

service UserService {
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq request) (api.get="/api/user/user-info");
}
