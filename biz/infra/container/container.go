package container

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/service/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/service/user"
	"sync"
)

type Container struct {
	Config      *config.Config
	UserService user.IUserService
	AuthService auth.IAuthService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(config *config.Config) {
	once.Do(func() {
		APP = GetContainer(config)
	})
}
func NewContainer(config *config.Config,
	userService user.IUserService,
	authService auth.IAuthService,

) *Container {
	return &Container{
		Config:      config,
		UserService: userService,
		AuthService: authService,
	}

}
