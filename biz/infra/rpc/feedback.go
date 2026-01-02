package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback/feedbackservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

func NewFeedbackClient(config *config.Config) feedbackservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	c, err := feedbackservice.NewClient(
		config.RpcConfig.FeedbackServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("留言 RPC 客户端启动失败" + err.Error())
	}
	return c
}
