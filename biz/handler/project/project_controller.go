// Code generated by hertz generator.

package project

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/container"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	project "github.com/li1553770945/sheepim-api-gateway/biz/model/project"
)

// GetProjects .
// @router /api/projects [GET]
func GetProjects(ctx context.Context, c *app.RequestContext) {
	var req project.ProjectsReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": constant.InvalidInput,
			"msg":  "参数校验失败:" + err.Error(),
		})
		return
	}
	App := container.GetGlobalContainer()
	resp := App.ProjectController.GetProjects(ctx, &req)
	c.JSON(consts.StatusOK, resp)
}

// GetProjectNum .
// @router /api/projects/num [GET]
func GetProjectNum(ctx context.Context, c *app.RequestContext) {
	App := container.GetGlobalContainer()
	resp := App.ProjectController.GetProjectNum(ctx)
	c.JSON(consts.StatusOK, resp)
}
