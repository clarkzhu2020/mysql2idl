// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go jobs
namespace py jobs
namespace java jobs
namespace rs jobs


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Jobs {

1:i32 id
2:string queue
3:string payload
4:i8 attempts
5:i32 reserved_at
6:i32 available_at
7:i32 created_at


}

//Create_service
struct CreateJobsRequest{

1:string queue  (api.body="queue", api.form="queue")
2:string payload  (api.body="payload", api.form="payload")
3:i8 attempts  (api.body="attempts", api.form="attempts")
4:i32 reserved_at  (api.body="reserved_at", api.form="reserved_at")
5:i32 available_at  (api.body="available_at", api.form="available_at")


}

struct CreateJobsResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryJobsRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryJobsResponse{
   1: Code code
   2: string msg
   3: list<Jobs> jobss
   4: i32 totoal
}

//Delete_service
struct DeleteJobsRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteJobsResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateJobsRequest{

1:i32 id  (api.body="id", api.form="id")
2:string queue  (api.body="queue", api.form="queue")
3:string payload  (api.body="payload", api.form="payload")
4:i8 attempts  (api.body="attempts", api.form="attempts")
5:i32 reserved_at  (api.body="reserved_at", api.form="reserved_at")
6:i32 available_at  (api.body="available_at", api.form="available_at")


}

struct UpdateJobsResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service JobsService {
   CreateJobsResponse CreateJobs(1:CreateJobsRequest req)(api.post="/api/jobs/create/")
   QueryJobsResponse  QueryJobs(1: QueryJobsRequest req)(api.get="/api/jobs/query/")
   DeleteJobsResponse DeleteJobs(1:DeleteJobsRequest req)(api.delete="/api/jobs/delete/:id")
   UpdateJobsResponse UpdateJobs(1:UpdateJobsRequest req)(api.put="/api/jobs/update/:id")
}