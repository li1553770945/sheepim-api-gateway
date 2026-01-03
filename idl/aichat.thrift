namespace go aichat

struct AIChatReq {
     1: required string message;
     2: optional string conversation_id
}

struct AIChatRespData {
    1: string message;
}

struct AIChatResp{
    1: required i32 code;
    2: required string message;
    3: optional AIChatRespData data;
}

struct AIChatSSEData {
    1: required string event_type;
    2: required string data;
}

// 如果正常只会使用SSE返回SSEData类型数据，异常情况才会使用AIChatResp
service AIChatController {
    AIChatResp SendMessage(1:AIChatReq request) (api.post="/api/aichat");
}
