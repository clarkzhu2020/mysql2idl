// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go cache_locks
namespace py cache_locks
namespace java cache_locks
namespace rs cache_locks


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Cache_locks {

1:string key
2:string owner
3:i32 expiration


}

//Create_service
struct CreateCache_locksRequest{

1:string key  (api.body="key", api.form="key")
2:string owner  (api.body="owner", api.form="owner")
3:i32 expiration  (api.body="expiration", api.form="expiration")


}

struct CreateCache_locksResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryCache_locksRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryCache_locksResponse{
   1: Code code
   2: string msg
   3: list<Cache_locks> cache_lockss
   4: i32 totoal
}

//Delete_service
struct DeleteCache_locksRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteCache_locksResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateCache_locksRequest{

1:string key  (api.body="key", api.form="key")
2:string owner  (api.body="owner", api.form="owner")
3:i32 expiration  (api.body="expiration", api.form="expiration")


}

struct UpdateCache_locksResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service Cache_locksService {
   CreateCache_locksResponse CreateCache_locks(1:CreateCache_locksRequest req)(api.post="/api/cache_locks/create/")
   QueryCache_locksResponse  QueryCache_locks(1: QueryCache_locksRequest req)(api.get="/api/cache_locks/query/")
   DeleteCache_locksResponse DeleteCache_locks(1:DeleteCache_locksRequest req)(api.delete="/api/cache_locks/delete/:id")
   UpdateCache_locksResponse UpdateCache_locks(1:UpdateCache_locksRequest req)(api.put="/api/cache_locks/update/:id")
}