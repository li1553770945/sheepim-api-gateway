//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"sheepim-api-gateway/biz/infra/config"
	"sheepim-api-gateway/biz/infra/rpc"
	"sheepim-api-gateway/biz/internal/service/auth"
	"sheepim-api-gateway/biz/internal/service/user"
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
