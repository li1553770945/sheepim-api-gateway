package container

import (
	"sheepim-api-gateway/biz/infra/config"
	"sheepim-api-gateway/biz/internal/service/user"
	"sheepim-user-service/kitex_gen/user/userservice"
	"sync"
)

type Container struct {
	Config        *config.Config
	UserRpcClient *userservice.Client
	UserService   user.IUserService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(env string) {
	once.Do(func() {
		APP = GetContainer(env)
	})
}
func NewContainer(config *config.Config,
	userService user.IUserService,
	userRpcClient *userservice.Client,
) *Container {
	return &Container{
		Config:        config,
		UserService:   userService,
		UserRpcClient: userRpcClient,
	}

}
