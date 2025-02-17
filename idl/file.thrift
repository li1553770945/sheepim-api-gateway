namespace go file

struct UploadFileReq{
    1: required string name;
    2: required i32 maxDownload;
    3: optional string key;
}

struct UploadFileRespData{
    1: required string ak;
    2: required string sk;
    3: required string key;
    4: required string ossPath;
}

struct UploadFileResp{
    1: required i32 code;
    2: required string message;
    3: optional UploadFileRespData data;
}

struct DownloadFileReq{
    1: required string key;
}
struct DownloadFileRespData{
    1: required string ak;
    2: required string sk;
    3: required string ossPath;
    4: required string name;
}

struct DownloadFileResp{
    1: required i32 code;
    2: required string message;
    3: optional DownloadFileRespData data;
}

struct DeleteFileReq{
    1: required string key;
}
struct DeleteFileResp{
    1: required i32 code;
    2: required string message;
}
service FileController {
    UploadFileResp UploadFile(UploadFileReq req)(api.post="/api/file");
    DownloadFileResp DownloadFile(DownloadFileReq req)(api.post="/api/file");
    DeleteFileResp DeleteFile(DeleteFileReq req)(api.delete="/api/file");
}
