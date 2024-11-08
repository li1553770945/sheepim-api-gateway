package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/personal-project-service/kitex_gen/project/projectservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

func NewProjectClient(config *config.Config) projectservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	projectClient, err := projectservice.NewClient(
		config.RpcConfig.ProjectServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("项目 RPC 客户端启动失败" + err.Error())
	}
	return projectClient
}
