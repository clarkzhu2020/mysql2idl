package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	users "goravel/hertz_gen/users"
)

type DeleteUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUsersService(Context context.Context, RequestContext *app.RequestContext) *DeleteUsersService {
	return &DeleteUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUsersService) Run(req *users.DeleteUsersRequest) (resp *users.DeleteUsersResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
