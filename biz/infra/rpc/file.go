package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/personal-file-service/kitex_gen/file/fileservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

func NewFileClient(config *config.Config) fileservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	fileClient, err := fileservice.NewClient(
		config.RpcConfig.FileServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("文件 RPC 客户端启动失败" + err.Error())
	}
	return fileClient
}
