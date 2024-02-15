package users

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"grom/biz/service"
	"grom/biz/utils"
	users "grom/hertz_gen/users"
)

// UpdateUsers .
// @router /v1/users/update/:id [POST]
func UpdateUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UpdateUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUpdateUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteUsers .
// @router /v1/users/delete/:id [POST]
func DeleteUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.DeleteUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewDeleteUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// QueryUsers .
// @router /v1/users/query/ [POST]
func QueryUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.QueryUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewQueryUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateUsers .
// @router /v1/users/create/ [POST]
func CreateUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.CreateUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCreateUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
