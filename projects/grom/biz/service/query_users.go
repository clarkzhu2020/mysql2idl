package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	users "grom/hertz_gen/users"
)

type QueryUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryUsersService(Context context.Context, RequestContext *app.RequestContext) *QueryUsersService {
	return &QueryUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryUsersService) Run(req *users.QueryUsersRequest) (resp *users.QueryUsersResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
