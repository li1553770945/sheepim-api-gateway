//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/rpc"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/service/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/service/user"
)

func GetContainer(config *config.Config) *Container {
	panic(wire.Build(

		rpc.NewUserClient,
		rpc.NewAuthClient,

		//service
		user.NewUserService,
		auth.NewUserService,

		NewContainer,
	))
}
