namespace go file

struct UploadFileReq{
    1: required string name;
    2: required i32 maxDownload;
    3: optional string key;
}

struct UploadFileRespData{
    1: required string key;
    2: required string signedUrl;
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
    1: required string signedUrl;
    2: required string name;
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

struct FileInfoReq{
    1: required string key;
}

struct FileInfoRespData{
    1: required string name;
    2: required string uploader_name;
    3: required i32 uploader_uid;
    4: required string created_at;
}
struct FileInfoResp{
    1: required i32 code;
    2: required string message;
    3: required FileInfoRespData data;
}


service FileController {
    UploadFileResp UploadFile(UploadFileReq req)(api.post="/api/files");
    DownloadFileResp DownloadFile(DownloadFileReq req)(api.get="/api/files/download");
    DeleteFileResp DeleteFile(DeleteFileReq req)(api.delete="/api/files");
    FileInfoResp FileInfo(FileInfoReq req)(api.get="/api/files");
}
