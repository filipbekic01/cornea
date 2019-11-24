package controllers

import (
	"github.com/kataras/iris/v12/mvc"
)

// HomeController .
type HomeController struct{}

// Get .
func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home.html",
		Data: nil,
	}
}
