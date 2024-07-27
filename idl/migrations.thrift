// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go migrations
namespace py migrations
namespace java migrations
namespace rs migrations


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Migrations {

1:i32 id
2:string migration
3:i32 batch


}

//Create_service
struct CreateMigrationsRequest{

1:string migration  (api.body="migration", api.form="migration")
2:i32 batch  (api.body="batch", api.form="batch")


}

struct CreateMigrationsResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryMigrationsRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryMigrationsResponse{
   1: Code code
   2: string msg
   3: list<Migrations> migrationss
   4: i32 totoal
}

//Delete_service
struct DeleteMigrationsRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteMigrationsResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateMigrationsRequest{

1:i32 id  (api.body="id", api.form="id")
2:string migration  (api.body="migration", api.form="migration")
3:i32 batch  (api.body="batch", api.form="batch")


}

struct UpdateMigrationsResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service MigrationsService {
   CreateMigrationsResponse CreateMigrations(1:CreateMigrationsRequest req)(api.post="/api/migrations/create/")
   QueryMigrationsResponse  QueryMigrations(1: QueryMigrationsRequest req)(api.get="/api/migrations/query/")
   DeleteMigrationsResponse DeleteMigrations(1:DeleteMigrationsRequest req)(api.delete="/api/migrations/delete/:id")
   UpdateMigrationsResponse UpdateMigrations(1:UpdateMigrationsRequest req)(api.put="/api/migrations/update/:id")
}