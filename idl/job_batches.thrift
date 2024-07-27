// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go job_batches
namespace py job_batches
namespace java job_batches
namespace rs job_batches


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Job_batches {

1:string id
2:string name
3:i32 total_jobs
4:i32 pending_jobs
5:i32 failed_jobs
6:string failed_job_ids
7:string options
8:i32 cancelled_at
9:i32 created_at
10:i32 finished_at


}

//Create_service
struct CreateJob_batchesRequest{

1:string name  (api.body="name", api.form="name")
2:i32 total_jobs  (api.body="total_jobs", api.form="total_jobs")
3:i32 pending_jobs  (api.body="pending_jobs", api.form="pending_jobs")
4:i32 failed_jobs  (api.body="failed_jobs", api.form="failed_jobs")
5:string failed_job_ids  (api.body="failed_job_ids", api.form="failed_job_ids")
6:string options  (api.body="options", api.form="options")
7:i32 cancelled_at  (api.body="cancelled_at", api.form="cancelled_at")
8:i32 finished_at  (api.body="finished_at", api.form="finished_at")


}

struct CreateJob_batchesResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryJob_batchesRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryJob_batchesResponse{
   1: Code code
   2: string msg
   3: list<Job_batches> job_batchess
   4: i32 totoal
}

//Delete_service
struct DeleteJob_batchesRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteJob_batchesResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateJob_batchesRequest{

1:string id  (api.body="id", api.form="id")
2:string name  (api.body="name", api.form="name")
3:i32 total_jobs  (api.body="total_jobs", api.form="total_jobs")
4:i32 pending_jobs  (api.body="pending_jobs", api.form="pending_jobs")
5:i32 failed_jobs  (api.body="failed_jobs", api.form="failed_jobs")
6:string failed_job_ids  (api.body="failed_job_ids", api.form="failed_job_ids")
7:string options  (api.body="options", api.form="options")
8:i32 cancelled_at  (api.body="cancelled_at", api.form="cancelled_at")
9:i32 finished_at  (api.body="finished_at", api.form="finished_at")


}

struct UpdateJob_batchesResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service Job_batchesService {
   CreateJob_batchesResponse CreateJob_batches(1:CreateJob_batchesRequest req)(api.post="/api/job_batches/create/")
   QueryJob_batchesResponse  QueryJob_batches(1: QueryJob_batchesRequest req)(api.get="/api/job_batches/query/")
   DeleteJob_batchesResponse DeleteJob_batches(1:DeleteJob_batchesRequest req)(api.delete="/api/job_batches/delete/:id")
   UpdateJob_batchesResponse UpdateJob_batches(1:UpdateJob_batchesRequest req)(api.put="/api/job_batches/update/:id")
}