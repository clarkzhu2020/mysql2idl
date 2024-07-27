// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go sessions
namespace py sessions
namespace java sessions
namespace rs sessions


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Sessions {

1:string id
2:i32 user_id
3:string ip_address
4:string user_agent
5:string payload
6:i32 last_activity


}

//Create_service
struct CreateSessionsRequest{

1:i32 user_id  (api.body="user_id", api.form="user_id")
2:string ip_address  (api.body="ip_address", api.form="ip_address")
3:string user_agent  (api.body="user_agent", api.form="user_agent")
4:string payload  (api.body="payload", api.form="payload")
5:i32 last_activity  (api.body="last_activity", api.form="last_activity")


}

struct CreateSessionsResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QuerySessionsRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QuerySessionsResponse{
   1: Code code
   2: string msg
   3: list<Sessions> sessionss
   4: i32 totoal
}

//Delete_service
struct DeleteSessionsRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteSessionsResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateSessionsRequest{

1:string id  (api.body="id", api.form="id")
2:i32 user_id  (api.body="user_id", api.form="user_id")
3:string ip_address  (api.body="ip_address", api.form="ip_address")
4:string user_agent  (api.body="user_agent", api.form="user_agent")
5:string payload  (api.body="payload", api.form="payload")
6:i32 last_activity  (api.body="last_activity", api.form="last_activity")


}

struct UpdateSessionsResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service SessionsService {
   CreateSessionsResponse CreateSessions(1:CreateSessionsRequest req)(api.post="/api/sessions/create/")
   QuerySessionsResponse  QuerySessions(1: QuerySessionsRequest req)(api.get="/api/sessions/query/")
   DeleteSessionsResponse DeleteSessions(1:DeleteSessionsRequest req)(api.delete="/api/sessions/delete/:id")
   UpdateSessionsResponse UpdateSessions(1:UpdateSessionsRequest req)(api.put="/api/sessions/update/:id")
}