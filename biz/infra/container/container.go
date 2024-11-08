package container

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/user"
	"sync"
)

type Container struct {
	Config            *config.Config
	UserController    user.IUserController
	AuthController    auth.IAuthController
	ProjectController project.IProjectController
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
	userController user.IUserController,
	authController auth.IAuthController,
	projectController project.IProjectController,

) *Container {
	return &Container{
		Config:            config,
		UserController:    userController,
		AuthController:    authController,
		ProjectController: projectController,
	}

}
