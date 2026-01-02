package aichat

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/personal-aichat-service/kitex_gen/aichat/aichatservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/aichat"
)

type AIChatControllerImpl struct {
	AIChatRpcClient aichatservice.Client
}

type IAIChatController interface {
	SendMessage(ctx context.Context, req *aichat.AIChatReq, c *app.RequestContext) (resp *aichat.AIChatResp)
}

func NewAIChatController(aichatRpcClient aichatservice.Client) IAIChatController {
	return &AIChatControllerImpl{
		AIChatRpcClient: aichatRpcClient,
	}
}
