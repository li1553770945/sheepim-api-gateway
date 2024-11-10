package container

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/feedback"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/user"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
	"sync"
)

type Container struct {
	Config *config.Config

	AuthRpcClient authservice.Client
	UserRpcClient userservice.Client

	UserController     user.IUserController
	AuthController     auth.IAuthController
	ProjectController  project.IProjectController
	FeedbackController feedback.IFeedbackController
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
	feedbackController feedback.IFeedbackController,

	authRpcClient authservice.Client,
	userRpcClient userservice.Client,
) *Container {
	return &Container{
		Config: config,

		AuthRpcClient: authRpcClient,
		UserRpcClient: userRpcClient,

		UserController:     userController,
		AuthController:     authController,
		ProjectController:  projectController,
		FeedbackController: feedbackController,
	}

}
