package feedback

import (
	"context"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback/feedbackservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/feedback"
)

type FeedbackControllerImpl struct {
	FeedbackRpcClient feedbackservice.Client
}

type IFeedbackController interface {
	GetFeedbackCategories(ctx context.Context) *feedback.GetFeedbackCategoryResp
	GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) *feedback.GetFeedbackResp
	AddFeedback(ctx context.Context, req *feedback.AddFeedbackReq) *feedback.AddFeedbackResp
	AddReply(ctx context.Context, req *feedback.AddReplyReq) *feedback.AddReplyResp
	GetReply(ctx context.Context, req *feedback.GetReplyReq) *feedback.GetReplyResp
	GetUnreadFeedback(ctx context.Context) *feedback.GetUnreadFeedbackResp
}

func NewFeedbackController(feedbackRpcClient feedbackservice.Client) IFeedbackController {
	return &FeedbackControllerImpl{
		FeedbackRpcClient: feedbackRpcClient,
	}
}
