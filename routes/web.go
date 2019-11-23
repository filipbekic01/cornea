package routes

import (
	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/kataras/iris"
)

func Web(app *iris.Application) {
	app.Get("/", controllers.Welcome)
}
