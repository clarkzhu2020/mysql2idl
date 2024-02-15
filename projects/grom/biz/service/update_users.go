package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	users "grom/hertz_gen/users"
)

type UpdateUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUsersService(Context context.Context, RequestContext *app.RequestContext) *UpdateUsersService {
	return &UpdateUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUsersService) Run(req *users.UpdateUsersRequest) (resp *users.UpdateUsersResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
