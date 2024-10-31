package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"sheepim-api-gateway/biz/infra/config"
	"sheepim-user-service/kitex_gen/user/userservice"
)

func NewUserClient(config *config.Config) userservice.Client {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdConfig.Endpoint})
	userClient, err := userservice.NewClient(
		"sheepim-user-service",
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	return userClient
}
