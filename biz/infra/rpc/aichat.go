package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/personal-aichat-service/kitex_gen/aichat/aichatservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

func NewAIChatClient(config *config.Config) aichatservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	if err != nil {
		panic("创建 AI 聊天服务的 etcd 解析器失败: " + err.Error())
	}
	
	aichatClient, err := aichatservice.NewClient(
		config.RpcConfig.AIChatServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("创建 AI 聊天服务 RPC 客户端失败: " + err.Error())
	}
	return aichatClient
}
