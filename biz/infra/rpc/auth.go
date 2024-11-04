package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sheepim-api-gateway/biz/infra/config"
	"sheepim-auth-service/kitex_gen/auth/authservice"
)

func NewAuthClient(config *config.Config) authservice.Client {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdConfig.Endpoint})
	userClient, err := authservice.NewClient(
		config.RpcConfig.AuthServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		panic("认证 RPC 客户端启动失败" + err.Error())
	}
	return userClient
}
