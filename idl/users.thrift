// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go users
namespace py users
namespace java users
namespace rs users


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Users {

1:i32 id
2:string name
3:string email
4:i64 email_verified_at
5:string password
6:string remember_token
7:i64 created_at
8:i64 updated_at


}

//Create_service
struct CreateUsersRequest{

1:string name  (api.body="name", api.form="name")
2:string email  (api.body="email", api.form="email")
3:i64 email_verified_at  (api.body="email_verified_at", api.form="email_verified_at")
4:string password  (api.body="password", api.form="password")
5:string remember_token  (api.body="remember_token", api.form="remember_token")


}

struct CreateUsersResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryUsersRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryUsersResponse{
   1: Code code
   2: string msg
   3: list<Users> userss
   4: i32 totoal
}

//Delete_service
struct DeleteUsersRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteUsersResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateUsersRequest{

1:i32 id  (api.body="id", api.form="id")
2:string name  (api.body="name", api.form="name")
3:string email  (api.body="email", api.form="email")
4:i64 email_verified_at  (api.body="email_verified_at", api.form="email_verified_at")
5:string password  (api.body="password", api.form="password")
6:string remember_token  (api.body="remember_token", api.form="remember_token")


}

struct UpdateUsersResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service UsersService {
   CreateUsersResponse CreateUsers(1:CreateUsersRequest req)(api.post="/api/users/create/")
   QueryUsersResponse  QueryUsers(1: QueryUsersRequest req)(api.get="/api/users/query/")
   DeleteUsersResponse DeleteUsers(1:DeleteUsersRequest req)(api.delete="/api/users/delete/:id")
   UpdateUsersResponse UpdateUsers(1:UpdateUsersRequest req)(api.put="/api/users/update/:id")
}