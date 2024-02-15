package controllers

import (
	_ "demo/app/model"
	"demo/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {

	var user models.User
	facades.Orm().Query().First(&user)
	//user.Name = "clark1"
	return ctx.Response().Success().Json(http.Json{
		"Hello": user.Name,
	})
}
