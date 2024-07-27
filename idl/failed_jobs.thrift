// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go failed_jobs
namespace py failed_jobs
namespace java failed_jobs
namespace rs failed_jobs


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Failed_jobs {

1:i32 id
2:string uuid
3:string connection
4:string queue
5:string payload
6:string exception
7:i64 failed_at


}

//Create_service
struct CreateFailed_jobsRequest{

1:string uuid  (api.body="uuid", api.form="uuid")
2:string connection  (api.body="connection", api.form="connection")
3:string queue  (api.body="queue", api.form="queue")
4:string payload  (api.body="payload", api.form="payload")
5:string exception  (api.body="exception", api.form="exception")
6:i64 failed_at  (api.body="failed_at", api.form="failed_at")


}

struct CreateFailed_jobsResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryFailed_jobsRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryFailed_jobsResponse{
   1: Code code
   2: string msg
   3: list<Failed_jobs> failed_jobss
   4: i32 totoal
}

//Delete_service
struct DeleteFailed_jobsRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteFailed_jobsResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateFailed_jobsRequest{

1:i32 id  (api.body="id", api.form="id")
2:string uuid  (api.body="uuid", api.form="uuid")
3:string connection  (api.body="connection", api.form="connection")
4:string queue  (api.body="queue", api.form="queue")
5:string payload  (api.body="payload", api.form="payload")
6:string exception  (api.body="exception", api.form="exception")
7:i64 failed_at  (api.body="failed_at", api.form="failed_at")


}

struct UpdateFailed_jobsResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service Failed_jobsService {
   CreateFailed_jobsResponse CreateFailed_jobs(1:CreateFailed_jobsRequest req)(api.post="/api/failed_jobs/create/")
   QueryFailed_jobsResponse  QueryFailed_jobs(1: QueryFailed_jobsRequest req)(api.get="/api/failed_jobs/query/")
   DeleteFailed_jobsResponse DeleteFailed_jobs(1:DeleteFailed_jobsRequest req)(api.delete="/api/failed_jobs/delete/:id")
   UpdateFailed_jobsResponse UpdateFailed_jobs(1:UpdateFailed_jobsRequest req)(api.put="/api/failed_jobs/update/:id")
}