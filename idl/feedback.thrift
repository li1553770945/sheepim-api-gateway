namespace go feedback

struct GetFeedbackCategoryRespData {
    1: required i64 id,
    2: required string name,
}

struct GetFeedbackCategoryResp {
    1: required i32 code
    2: required string message
    3: optional list<GetFeedbackCategoryRespData> data
}

struct GetFeedbackRespData {
    1: required i64 id
    2: required string title,
    3: required string content,
    4: required string name,
    5: required string contact
    6: required string category
    7: required i64 createdAt
}

struct GetFeedbackResp {
    1: required i32 code,
    2: required string message,
    3: optional GetFeedbackRespData data
}

struct AddFeedbackRespData {
    1: required string uuid
}

struct AddFeedbackResp {
    1: required i32 code,
    2: required string message,
    3: optional AddFeedbackRespData data
}

struct AddReplyRespData {
    1: required string reply_content
}

struct AddReplyResp {
    1: required i32 code,
    2: required string message,
    3: optional AddReplyRespData data
}

struct GetReplyRespData {
    1: required string content
    2: required i64 createdAt
}

struct GetReplyResp {
    1: required i32 code,
    2: required string message,
    3: optional GetReplyRespData data
}

struct GetFeedbackReq {
    1: required string uuid (api.query="uuid")
}

struct AddFeedbackReq {
    1: required i64 categoryId (api.form="categoryId"),
    2: required string title (api.form="title"),
    3: required string content (api.form="content"),
    4: required string name (api.form="name"),
    5: optional string contact (api.form="contact")
}

struct AddReplyReq {
    1: required i64 feedbackId (api.form="feedbackId"),
    2: required string content (api.form="content")
}

struct GetReplyReq {
    1: required string feedbackUuid (api.query="feedbackUuid")
}

struct UnreadFeedbackData {
    1: required i64 id
    2: required string title
    3: required string name
    4: required i64 createdAt
    5: required string uuid
}

struct GetUnreadFeedbackResp {
    1: required i32 code,
    2: required string message,
    3: optional list<UnreadFeedbackData> data
}

service FeedbackController {
    GetFeedbackCategoryResp GetFeedbackCategories() (api.get="/api/feedback/categories")
    GetFeedbackResp GetFeedback(GetFeedbackReq req) (api.get="/api/feedback")
    AddFeedbackResp AddFeedback(AddFeedbackReq req) (api.post="/api/feedback")
    AddReplyResp AddReply(AddReplyReq req) (api.post="/api/feedback/reply")
    GetReplyResp GetReply(GetReplyReq req) (api.get="/api/feedback/reply")
    GetUnreadFeedbackResp GetUnreadFeedback() (api.get="/api/feedback/unread")
}
