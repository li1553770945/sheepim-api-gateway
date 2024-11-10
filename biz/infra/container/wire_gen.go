// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/rpc"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/auth"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/feedback"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/controller/user"
)

// Injectors from wire.go:

func GetContainer(config2 *config.Config) *Container {
	client := rpc.NewUserClient(config2)
	iUserController := user.NewUserController(client)
	authserviceClient := rpc.NewAuthClient(config2)
	iAuthController := auth.NewAuthController(authserviceClient)
	projectserviceClient := rpc.NewProjectClient(config2)
	iProjectController := project.NewProjectController(projectserviceClient)
	feedbackserviceClient := rpc.NewFeedbackClient(config2)
	iFeedbackController := feedback.NewFeedbackController(feedbackserviceClient)
	container := NewContainer(config2, iUserController, iAuthController, iProjectController, iFeedbackController, authserviceClient, client)
	return container
}
