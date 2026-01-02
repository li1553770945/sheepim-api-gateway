namespace go aichat

struct AIChatReq {
     1: required string message;
}

struct AIChatRespData {
    1: string message;
}

struct AIChatResp{
    1: required i32 code;
    2: required string message;
    3: optional AIChatRespData data;
}

service AIChatController {
    AIChatResp SendMessage(1:AIChatReq request) (api.post="/api/aichat");
}
