package assembler

import (
	rpcFeedback "github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/feedback"
)

func GetFeedbackCategoryRespRpcToHttp(rpcResp *rpcFeedback.GetFeedbackCategoryResp) *feedback.GetFeedbackCategoryResp {
	return &feedback.GetFeedbackCategoryResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data:    rpcCategoryDataRpcToHttp(rpcResp.Data),
	}
}

func rpcCategoryDataRpcToHttp(data []*rpcFeedback.GetFeedbackCategoryRespData) []*feedback.GetFeedbackCategoryRespData {
	var result = make([]*feedback.GetFeedbackCategoryRespData, 0)
	for _, d := range data {
		result = append(result, &feedback.GetFeedbackCategoryRespData{
			ID:   d.Id,
			Name: d.Name,
		})
	}
	return result
}

func GetFeedbackReqHttpToRpc(req *feedback.GetFeedbackReq) *rpcFeedback.GetFeedbackReq {
	return &rpcFeedback.GetFeedbackReq{
		Uuid: req.UUID, // 修正字段名
	}
}

func GetFeedbackRespRpcToHttp(rpcResp *rpcFeedback.GetFeedbackResp) *feedback.GetFeedbackResp {
	return &feedback.GetFeedbackResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data: &feedback.GetFeedbackRespData{
			ID:        rpcResp.Id,
			Title:     rpcResp.Title,
			Content:   rpcResp.Content,
			Name:      rpcResp.Name,
			Contact:   rpcResp.Contact,
			Category:  rpcResp.Category,
			CreatedAt: rpcResp.CreatedAt,
		},
	}
}

func AddFeedbackReqHttpToRpc(req *feedback.AddFeedbackReq) *rpcFeedback.AddFeedBackReq {
	return &rpcFeedback.AddFeedBackReq{ // 使用正确的类型名
		CategoryId: req.CategoryId, // 修正字段名
		Title:      req.Title,
		Content:    req.Content,
		Name:       req.Name,
		Contact:    req.Contact,
	}
}

func AddFeedbackRespRpcToHttp(rpcResp *rpcFeedback.AddFeedbackResp) *feedback.AddFeedbackResp {
	return &feedback.AddFeedbackResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data: &feedback.AddFeedbackRespData{
			UUID: *rpcResp.Uuid,
		},
	}
}

func AddReplyReqHttpToRpc(req *feedback.AddReplyReq) *rpcFeedback.AddReplyReq {
	return &rpcFeedback.AddReplyReq{
		FeedbackId: req.FeedbackId, // 修正字段名
		Content:    req.Content,
	}
}

func AddReplyRespRpcToHttp(rpcResp *rpcFeedback.AddReplyResp) *feedback.AddReplyResp {
	return &feedback.AddReplyResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
	}
}

func GetReplyReqHttpToRpc(req *feedback.GetReplyReq) *rpcFeedback.GetReplyReq {
	return &rpcFeedback.GetReplyReq{
		FeedbackUuid: req.FeedbackUuid,
	}
}

func GetReplyRespRpcToHttp(rpcResp *rpcFeedback.GetReplyResp) *feedback.GetReplyResp {
	if rpcResp.BaseResp.Code != 0 {
		return &feedback.GetReplyResp{
			Code:    rpcResp.BaseResp.Code,
			Message: rpcResp.BaseResp.Message,
		}
	}
	var content string
	// 检查 Content 是否为 nil
	if rpcResp.Content != nil {
		content = *rpcResp.Content
	}

	return &feedback.GetReplyResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data: &feedback.GetReplyRespData{
			Content:   content,
			CreatedAt: *rpcResp.CreatedAt,
		},
	}
}

func GetUnreadFeedbackRespRpcToHttp(rpcResp *rpcFeedback.GetUnreadFeedbackResp) *feedback.GetUnreadFeedbackResp {
	return &feedback.GetUnreadFeedbackResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data:    rpcUnreadFeedbackDataRpcToHttp(rpcResp.Data),
	}
}

func rpcUnreadFeedbackDataRpcToHttp(data []*rpcFeedback.UnreadFeedbackData) []*feedback.UnreadFeedbackData {
	if data == nil {
		return nil
	}
	var result []*feedback.UnreadFeedbackData
	for _, d := range data {
		result = append(result, &feedback.UnreadFeedbackData{
			ID:        d.Id,
			Title:     d.Title,
			Name:      d.Name,
			CreatedAt: d.CreatedAt,
			UUID:      d.Uuid,
		})
	}
	return result
}
