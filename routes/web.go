package routes

import (
	"github.com/kataras/iris"
)

func Web(app *iris.Application) {
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
}
