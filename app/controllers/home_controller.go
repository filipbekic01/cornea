package controllers

import (
	"github.com/filipbekic01/cornea/app/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// HomeController .
type HomeController struct {
	Ctx     iris.Context
	Service services.UserService
}

// Get .
func (c *HomeController) Get() mvc.Result {

	_ = c.Service.GetAll()

	return mvc.View{
		Name: "home.html",
		Data: nil,
	}
}
