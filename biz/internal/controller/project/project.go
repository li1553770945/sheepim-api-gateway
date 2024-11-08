package project

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/project"
)

func (c ProjectControllerImpl) GetProjects(ctx context.Context, req *project.ProjectsReq) *project.ProjectsResp {
	hlog.CtxInfof(ctx, "调用项目服务获取项目列表")
	rpcReq := assembler.GetProjectsReqHttpToRpc(req)
	// 调用 RPC 客户端获取项目列表
	rpcResp, err := c.ProjectRpcClient.GetProjects(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用项目服务失败: %s", err.Error())
		return &project.ProjectsResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用项目服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用项目服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	// 使用 assembler 函数转换响应
	return assembler.GetProjectsRespRpcToHttp(rpcResp)
}

func (c ProjectControllerImpl) GetProjectNum(ctx context.Context) *project.ProjectNumResp {
	hlog.CtxInfof(ctx, "获取项目总数")

	hlog.CtxInfof(ctx, "调用项目服务获取总数")
	rpcResp, err := c.ProjectRpcClient.GetProjectNum(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用项目服务失败:"+err.Error())
		return &project.ProjectNumResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用用户服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用项目服务成功，服务返回状态码:%d", rpcResp.BaseResp.Code)
	return assembler.GetProjectNumRespRpcToHttp(rpcResp)
}
