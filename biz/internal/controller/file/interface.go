package file

import (
	"context"

	"github.com/li1553770945/personal-file-service/kitex_gen/file/fileservice"
	"github.com/li1553770945/sheepim-api-gateway/biz/model/file"
)

type FileControllerImpl struct {
	fileRpcClient fileservice.Client
}

type IFileController interface {
	UploadFile(ctx context.Context, req *file.UploadFileReq) (resp *file.UploadFileResp)
	DownloadFile(ctx context.Context, req *file.DownloadFileReq) (resp *file.DownloadFileResp)
	DeleteFile(ctx context.Context, req *file.DeleteFileReq) (resp *file.DeleteFileResp)
	FileInfo(ctx context.Context, req *file.FileInfoReq) (resp *file.FileInfoResp)
	DirectDownload(ctx context.Context, req *file.DownloadFileReq) (resp *file.DownloadFileResp)
}

func NewFileController(fileRpcClient fileservice.Client) IFileController {
	return &FileControllerImpl{
		fileRpcClient: fileRpcClient,
	}
}
