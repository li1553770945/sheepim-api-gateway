package feedback

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/feedback"
)

func (c FeedbackControllerImpl) GetFeedbackCategories(ctx context.Context) *feedback.GetFeedbackCategoryResp {
	hlog.CtxInfof(ctx, "调用反馈服务获取反馈类别列表")

	rpcResp, err := c.FeedbackRpcClient.GetFeedbackCategories(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.GetFeedbackCategoryResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)
	if rpcResp.Data != nil {
		hlog.CtxInfof(ctx, "共获取到%d条数据", len(rpcResp.Data))
	}
	httpResp := assembler.GetFeedbackCategoryRespRpcToHttp(rpcResp)
	return httpResp
}

func (c FeedbackControllerImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) *feedback.GetFeedbackResp {
	hlog.CtxInfof(ctx, "调用反馈服务获取反馈详情")

	rpcReq := assembler.GetFeedbackReqHttpToRpc(req)
	rpcResp, err := c.FeedbackRpcClient.GetFeedback(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.GetFeedbackResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	return assembler.GetFeedbackRespRpcToHttp(rpcResp)
}

func (c FeedbackControllerImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedbackReq) *feedback.AddFeedbackResp {
	hlog.CtxInfof(ctx, "调用反馈服务添加反馈")

	rpcReq := assembler.AddFeedbackReqHttpToRpc(req)
	rpcResp, err := c.FeedbackRpcClient.AddFeedback(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.AddFeedbackResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	return assembler.AddFeedbackRespRpcToHttp(rpcResp)
}

func (c FeedbackControllerImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) *feedback.AddReplyResp {
	hlog.CtxInfof(ctx, "调用反馈服务添加回复")

	rpcReq := assembler.AddReplyReqHttpToRpc(req)
	rpcResp, err := c.FeedbackRpcClient.AddReply(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.AddReplyResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	return assembler.AddReplyRespRpcToHttp(rpcResp)
}

func (c FeedbackControllerImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) *feedback.GetReplyResp {
	hlog.CtxInfof(ctx, "调用反馈服务获取回复")

	rpcReq := assembler.GetReplyReqHttpToRpc(req)
	rpcResp, err := c.FeedbackRpcClient.GetReply(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.GetReplyResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	return assembler.GetReplyRespRpcToHttp(rpcResp)
}

func (c FeedbackControllerImpl) GetUnreadFeedback(ctx context.Context) *feedback.GetUnreadFeedbackResp {
	hlog.CtxInfof(ctx, "调用反馈服务获取未读反馈")

	rpcResp, err := c.FeedbackRpcClient.GetUnreadFeedback(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用反馈服务失败: %s", err.Error())
		return &feedback.GetUnreadFeedbackResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用反馈服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用反馈服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	return assembler.GetUnreadFeedbackRespRpcToHttp(rpcResp)
}
