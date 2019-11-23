package controllers

import "github.com/kataras/iris"

func Welcome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "Welcome to Cornea web framework.",
	})
}
