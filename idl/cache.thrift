// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go cache
namespace py cache
namespace java cache
namespace rs cache


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Cache {

1:string key
2:string value
3:i32 expiration


}

//Create_service
struct CreateCacheRequest{

1:string key  (api.body="key", api.form="key")
2:string value  (api.body="value", api.form="value")
3:i32 expiration  (api.body="expiration", api.form="expiration")


}

struct CreateCacheResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryCacheRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryCacheResponse{
   1: Code code
   2: string msg
   3: list<Cache> caches
   4: i32 totoal
}

//Delete_service
struct DeleteCacheRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteCacheResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateCacheRequest{

1:string key  (api.body="key", api.form="key")
2:string value  (api.body="value", api.form="value")
3:i32 expiration  (api.body="expiration", api.form="expiration")


}

struct UpdateCacheResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service CacheService {
   CreateCacheResponse CreateCache(1:CreateCacheRequest req)(api.post="/api/cache/create/")
   QueryCacheResponse  QueryCache(1: QueryCacheRequest req)(api.get="/api/cache/query/")
   DeleteCacheResponse DeleteCache(1:DeleteCacheRequest req)(api.delete="/api/cache/delete/:id")
   UpdateCacheResponse UpdateCache(1:UpdateCacheRequest req)(api.put="/api/cache/update/:id")
}