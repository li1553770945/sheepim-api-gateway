//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/log"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/rpc"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/trace"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/feedback"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/file"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/room"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/user"
)

func GetContainer(env string) *Container {
	panic(wire.Build(
		//infra
		config.GetConfig,
		trace.InitTrace,
		log.InitLog,

		rpc.NewUserClient,
		rpc.NewAuthClient,
		rpc.NewProjectClient,
		rpc.NewFeedbackClient,
		rpc.NewRoomClient,
		rpc.NewFileClient,

		//controller
		user.NewUserController,
		auth.NewAuthController,
		project.NewProjectController,
		feedback.NewFeedbackController,
		room.NewRoomController,
		file.NewFileController,

		NewContainer,
	))
}
