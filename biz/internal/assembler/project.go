package assembler

import (
	projectRpc "github.com/li1553770945/personal-project-service/kitex_gen/project"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/project"
)

func GetProjectsReqHttpToRpc(httpReq *project.ProjectsReq) *projectRpc.ProjectsReq {
	return &projectRpc.ProjectsReq{
		Start:  httpReq.Start,
		End:    httpReq.End,
		Order:  httpReq.Order,
		Status: httpReq.Status,
	}
}

func GetProjectNumRespRpcToHttp(rpcResp *projectRpc.ProjectNumResp) *project.ProjectNumResp {
	// 定义默认的 code 和 message，如果 baseResp 为 nil，可以设置默认值
	var code int32 = 200 // 假设 200 表示成功
	var message string = "Success"
	var data *int64 = nil

	// 检查 rpcResp 的 BaseResp 是否存在，获取其中的 code 和 message
	if rpcResp.BaseResp != nil {
		code = rpcResp.BaseResp.Code       // 假设 BaseResp 中有 Code 字段
		message = rpcResp.BaseResp.Message // 假设 BaseResp 中有 Message 字段
	}

	// 获取 rpcResp 中的 Num 字段
	data = rpcResp.Num

	// 返回 HTTP 层的 ProjectNumResp
	return &project.ProjectNumResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func GetProjectsRespRpcToHttp(rpcResp *projectRpc.ProjectsResp) *project.ProjectsResp {
	// 初始化 HTTP 层的项目列表
	var projectDataList []*project.ProjectData

	// 遍历 RPC 响应中的项目列表，将每个项目转换为 HTTP 层的 ProjectData
	for _, rpcData := range rpcResp.ProjectData {
		projectData := &project.ProjectData{
			Name:         rpcData.Name,
			Desc:         rpcData.Desc,
			Link:         rpcData.Link,
			Status:       rpcData.Status,
			VolumeOfWork: rpcData.VolumeOfWork,
			Difficulty:   rpcData.Difficulty,
			StartDate:    rpcData.StartDate,
		}

		// 将 EndDate 从 int64 转换为 *int64，如果为空则保持 nil
		if rpcData.EndDate != nil {
			endDate := *rpcData.EndDate
			projectData.EndDate = &endDate
		}

		// 将转换后的项目添加到列表中
		projectDataList = append(projectDataList, projectData)
	}

	// 构建并返回 HTTP 层的 ProjectsResp
	return &project.ProjectsResp{
		Code:    rpcResp.BaseResp.Code,    // 假设 BaseResp 中包含 Code 字段
		Message: rpcResp.BaseResp.Message, // 假设 BaseResp 中包含 Message 字段
		Data:    projectDataList,
	}
}
