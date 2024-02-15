package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	users "grom/hertz_gen/users"
)

type CreateUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateUsersService(Context context.Context, RequestContext *app.RequestContext) *CreateUsersService {
	return &CreateUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateUsersService) Run(req *users.CreateUsersRequest) (resp *users.CreateUsersResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
