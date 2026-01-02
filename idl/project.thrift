namespace go project

struct ProjectsReq{
    1: optional i32 start (api.query="start");
    2: optional i32 end (api.query="end")
    3: optional string order (api.query="order")
    4: optional i32 status (api.query="status")
}
struct ProjectData{
     1: string name,
     2: string desc,
     3: string link,
     4: i32 status,
     5: i32 volume_of_work,
     6: i32 difficulty,
     7: i64 start_date,      
     8: optional i64 end_date
   }

struct ProjectsResp {
    1: required i32 code
    2: required string message
    3: optional list<ProjectData> data
}

struct ProjectNumResp{
    1: required i32 code
    2: required string message
    3: optional i64 data
}
service ProjectController {
    ProjectsResp GetProjects(ProjectsReq req) (api.get="/api/projects")
    ProjectNumResp GetProjectNum() (api.get="/api/projects/num")
}
