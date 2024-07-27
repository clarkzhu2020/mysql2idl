// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go password_reset_tokens
namespace py password_reset_tokens
namespace java password_reset_tokens
namespace rs password_reset_tokens


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Password_reset_tokens {

1:string email
2:string token
3:i64 created_at


}

//Create_service
struct CreatePassword_reset_tokensRequest{

1:string email  (api.body="email", api.form="email")
2:string token  (api.body="token", api.form="token")


}

struct CreatePassword_reset_tokensResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryPassword_reset_tokensRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryPassword_reset_tokensResponse{
   1: Code code
   2: string msg
   3: list<Password_reset_tokens> password_reset_tokenss
   4: i32 totoal
}

//Delete_service
struct DeletePassword_reset_tokensRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeletePassword_reset_tokensResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdatePassword_reset_tokensRequest{

1:string email  (api.body="email", api.form="email")
2:string token  (api.body="token", api.form="token")


}

struct UpdatePassword_reset_tokensResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service Password_reset_tokensService {
   CreatePassword_reset_tokensResponse CreatePassword_reset_tokens(1:CreatePassword_reset_tokensRequest req)(api.post="/api/password_reset_tokens/create/")
   QueryPassword_reset_tokensResponse  QueryPassword_reset_tokens(1: QueryPassword_reset_tokensRequest req)(api.get="/api/password_reset_tokens/query/")
   DeletePassword_reset_tokensResponse DeletePassword_reset_tokens(1:DeletePassword_reset_tokensRequest req)(api.delete="/api/password_reset_tokens/delete/:id")
   UpdatePassword_reset_tokensResponse UpdatePassword_reset_tokens(1:UpdatePassword_reset_tokensRequest req)(api.put="/api/password_reset_tokens/update/:id")
}