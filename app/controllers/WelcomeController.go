package controllers

import "github.com/kataras/iris/v12"

// Welcome description missing.
func Welcome(ctx iris.Context) {
	ctx.ViewData("message", "This is dynamic message")
	ctx.View("welcome.html")
}
