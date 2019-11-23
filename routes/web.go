package routes

import (
	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/kataras/iris/v12"
)

// Web description is missing.
func Web(app *iris.Application) {
	app.Get("/", controllers.Welcome)
}
