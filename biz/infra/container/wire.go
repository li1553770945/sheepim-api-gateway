//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"sheepim-api-gateway/biz/infra/config"
	"sheepim-api-gateway/biz/internal/service/user"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.InitConfig,

		//service
		user.NewUserService,

		NewContainer,
	))
}
