package assembler

import (
	rpc "github.com/li1553770945/personal-file-service/kitex_gen/file"
	http "github.com/li1553770945/sheepim-api-gateway/biz/model/file"
)

func UploadFileReqHttpToRpc(httpReq *http.UploadFileReq) (rpcReq *rpc.UploadFileReq) {
	return &rpc.UploadFileReq{
		Name:        httpReq.Name,
		MaxDownload: httpReq.MaxDownload,
		Key:         httpReq.Key,
	}
}

func UploadFileRespRpcToHttp(rpcResp *rpc.UploadFileResp) (httpResp *http.UploadFileResp) {
	return &http.UploadFileResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data: &http.UploadFileRespData{
			Key:       rpcResp.Key,
			SignedUrl: rpcResp.SignedUrl,
		},
	}
}

func DownloadFileReqHttpToRpc(httpReq *http.DownloadFileReq) (rpcReq *rpc.DownloadFileReq) {
	return &rpc.DownloadFileReq{
		Key: httpReq.Key,
	}
}

func DownloadFileRespRpcToHttp(rpcResp *rpc.DownloadFileResp) (httpResp *http.DownloadFileResp) {
	return &http.DownloadFileResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
		Data: &http.DownloadFileRespData{
			Name:      rpcResp.Name,
			SignedUrl: rpcResp.SignedUrl,
		},
	}
}

func DeleteFileReqHttpToRpc(httpReq *http.DeleteFileReq) (rpcReq *rpc.DeleteFileReq) {
	return &rpc.DeleteFileReq{
		Key: httpReq.Key,
	}
}

func DeleteFileRespRpcToHttp(rpcResp *rpc.DeleteFileResp) (httpResp *http.DeleteFileResp) {
	return &http.DeleteFileResp{
		Code:    rpcResp.BaseResp.Code,
		Message: rpcResp.BaseResp.Message,
	}
}
