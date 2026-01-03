package aichat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/sse"
	aichatConstant "github.com/li1553770945/personal-aichat-service/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/aichat"
)

func (client *AIChatControllerImpl) SendMessage(ctx context.Context, req *aichat.AIChatReq, c *app.RequestContext) (resp *aichat.AIChatResp) {
	if req == nil {
		hlog.Error("HTTP Request req is nil")
		return &aichat.AIChatResp{Message: "非法请求：req为空"}
	}

	// 执行转换
	rpcReq := assembler.AIChatReqHttpToRpc(req)

	if client.AIChatRpcClient == nil {
		print("AIChatRpcClient is nil")
		hlog.Error("AIChatRpcClient is nil")
		return &aichat.AIChatResp{Message: "系统错误：AIChatRpcClient未初始化"}
	}
	stream, err := client.AIChatRpcClient.SendMessage(ctx, rpcReq)
	if err != nil {
		return &aichat.AIChatResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用聊天服务失败:" + err.Error(),
		}
	}
	w := sse.NewWriter(c)

	for {
		resp, err := stream.Recv(stream.Context())
		if errors.Is(err, io.EOF) {
			err = w.Close()
			if err != nil {
				hlog.Warn(err)
			}
			// 正常结束不通过标准HTTP返回
			return nil
		}
		if err != nil {
			hlog.Warnf("从stream读取aichat消息出错:%s", err.Error())
			sseResp := &aichat.AIChatSSEData{
				EventType: aichatConstant.EventTypeError,
				Data:      fmt.Sprintf("从stream读取aichat消息出错:%s", err.Error()),
			}
			jsonBytes, _ := json.Marshal(sseResp)
			err = w.WriteEvent("id-x", "message", jsonBytes)
			if err != nil {
				hlog.Errorf("sse遇到错误，但将错误发送到前端失败:%s", err.Error())
			}
			return nil
		}
		sseResp := &aichat.AIChatSSEData{
			EventType: resp.EventType,
			Data:      resp.Data,
		}
		jsonBytes, _ := json.Marshal(sseResp)
		err = w.WriteEvent("id-x", "message", jsonBytes)
		if err != nil {
			hlog.Errorf("sse发送消息失败:%s", err.Error())
			return nil
		}
	}
}
