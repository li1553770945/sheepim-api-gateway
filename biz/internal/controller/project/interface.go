package project

import (
	"context"
	"github.com/li1553770945/personal-project-service/kitex_gen/project/projectservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/project"
)

type ProjectControllerImpl struct {
	ProjectRpcClient projectservice.Client
}

type IProjectController interface {
	GetProjects(ctx context.Context, req *project.ProjectsReq) *project.ProjectsResp
	GetProjectNum(ctx context.Context) *project.ProjectNumResp
}

func NewProjectController(projectRpcClient projectservice.Client) IProjectController {
	return &ProjectControllerImpl{
		ProjectRpcClient: projectRpcClient,
	}
}
