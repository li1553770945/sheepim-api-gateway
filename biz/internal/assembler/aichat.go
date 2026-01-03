package assembler

import (
	rpc "github.com/li1553770945/personal-aichat-service/kitex_gen/aichat"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/aichat"
)

func AIChatReqHttpToRpc(req *aichat.AIChatReq) *rpc.SendMessageReq {
	return &rpc.SendMessageReq{
		Message:        req.Message,
		ConversationId: req.ConversationID,
	}
}
