namespace go feedback

struct GetFeedbackCategoryRespData {
    1: required i64 id,
    2: required string name,
    3: required string value
}

struct GetFeedbackCategoryResp {
    1: required i32 code,
    2: required string message,
    3: optional list<GetFeedbackCategoryRespData> data
}

struct GetFeedbackRespData {
    1: required i64 id
    2: required string title,
    3: required string content,
    4: required string name,
    5: optional string contact
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
    1: required i64 category_id (api.form="category_id"),
    2: required string title (api.form="title"),
    3: required string content (api.form="content"),
    4: required string name (api.form="name"),
    5: optional string contact (api.form="contact")
}

struct AddReplyReq {
    1: required i64 feedback_id (api.form="feedback_id"),
    2: required string content (api.form="content")
}

struct GetReplyReq {
    1: required string feedback_uuid (api.query="feedback_uuid")
}

struct UnreadFeedbackData {
    1: required i64 id,
    2: required string title
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
