namespace go user

struct GetUserInfoReq {
    1: string username (api.query="username"); // 添加 api 注解为方便进行参数绑定
}

struct GetUserInfoResp {
    1: string username;
    2: string nickname;
    3: string role;
}


service UserService {
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq request) (api.get="/user-info");
}
