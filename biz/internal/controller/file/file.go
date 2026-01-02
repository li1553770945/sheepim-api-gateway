package file

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"github.com/li1553770945/sheepim-api-gateway/biz/internal/assembler"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/file"
)

func (c *FileControllerImpl) UploadFile(ctx context.Context, req *file.UploadFileReq) (resp *file.UploadFileResp) {
	hlog.CtxInfof(ctx, "调用上传文件服务")

	rpcReq := assembler.UploadFileReqHttpToRpc(req)
	rpcResp, err := c.fileRpcClient.UploadFile(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用上传文件服务失败: %s", err.Error())
		return &file.UploadFileResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用上传文件服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用上传文件服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	httpResp := assembler.UploadFileRespRpcToHttp(rpcResp)
	return httpResp
}
func (c *FileControllerImpl) DownloadFile(ctx context.Context, req *file.DownloadFileReq) (resp *file.DownloadFileResp) {
	hlog.CtxInfof(ctx, "调用下载文件服务")

	rpcReq := assembler.DownloadFileReqHttpToRpc(req)
	rpcResp, err := c.fileRpcClient.DownloadFile(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用下载文件服务失败: %s", err.Error())
		return &file.DownloadFileResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用下载文件服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用下载文件服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	httpResp := assembler.DownloadFileRespRpcToHttp(rpcResp)
	return httpResp
}
func (c *FileControllerImpl) DeleteFile(ctx context.Context, req *file.DeleteFileReq) (resp *file.DeleteFileResp) {
	hlog.CtxInfof(ctx, "调用删除文件服务")

	rpcReq := assembler.DeleteFileReqHttpToRpc(req)
	rpcResp, err := c.fileRpcClient.DeleteFile(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用删除文件服务失败: %s", err.Error())
		return &file.DeleteFileResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用删除文件服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用删除文件服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	httpResp := assembler.DeleteFileRespRpcToHttp(rpcResp)
	return httpResp
}

func (c *FileControllerImpl) FileInfo(ctx context.Context, req *file.FileInfoReq) (resp *file.FileInfoResp) {
	hlog.CtxInfof(ctx, "调用文件信息服务")

	rpcReq := assembler.FileInfoReqHttpToRpc(req)
	rpcResp, err := c.fileRpcClient.FileInfo(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用删除文件服务失败: %s", err.Error())
		return &file.FileInfoResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用删除文件服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用删除文件服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	httpResp := assembler.FileInfoRespRpcToHttp(rpcResp)
	return httpResp
}

func (c *FileControllerImpl) DirectDownload(ctx context.Context, req *file.DownloadFileReq) (resp *file.DownloadFileResp) {
	hlog.CtxInfof(ctx, "调用直接下载文件服务")

	rpcReq := assembler.DownloadFileReqHttpToRpc(req)
	rpcResp, err := c.fileRpcClient.DownloadFile(ctx, rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "调用直接下载文件服务失败: %s", err.Error())
		return &file.DownloadFileResp{
			Code:    constant.SystemError,
			Message: "系统错误：调用直接下载文件服务失败",
		}
	}
	hlog.CtxInfof(ctx, "调用直接下载文件服务成功，服务返回状态码: %d", rpcResp.BaseResp.Code)

	httpResp := assembler.DownloadFileRespRpcToHttp(rpcResp)
	return httpResp
}
