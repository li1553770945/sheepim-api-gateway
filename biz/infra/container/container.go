package container

import (
	"github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/li1553770945/personal-file-service/kitex_gen/file/fileservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/trace"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/feedback"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/file"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/room"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/user"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
	"sync"
)

type Container struct {
	Config      *config.Config
	TraceStruct *trace.TraceSturct
	Logger      *logrus.Logger

	AuthRpcClient authservice.Client
	UserRpcClient userservice.Client
	FileRpcClient fileservice.Client

	UserController     user.IUserController
	AuthController     auth.IAuthController
	ProjectController  project.IProjectController
	FeedbackController feedback.IFeedbackController
	RoomController     room.IRoomController
	FileController     file.IFileController
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
	traceStruct *trace.TraceSturct,
	logger *logrus.Logger,

	userController user.IUserController,
	authController auth.IAuthController,
	projectController project.IProjectController,
	feedbackController feedback.IFeedbackController,
	roomController room.IRoomController,
	fileController file.IFileController,

	authRpcClient authservice.Client,
	userRpcClient userservice.Client,
	fileRpcClient fileservice.Client,
) *Container {
	return &Container{
		Config:      config,
		TraceStruct: traceStruct,
		Logger:      logger,

		AuthRpcClient: authRpcClient,
		UserRpcClient: userRpcClient,
		FileRpcClient: fileRpcClient,

		UserController:     userController,
		AuthController:     authController,
		ProjectController:  projectController,
		FeedbackController: feedbackController,
		RoomController:     roomController,
		FileController:     fileController,
	}

}
