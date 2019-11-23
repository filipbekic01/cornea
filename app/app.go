package app

import (
	"github.com/filipbekic01/cornea/routes"
	"github.com/kataras/iris/v12"
)

// Run description is missing.
func Run() {
	app := iris.New()

	// Register routes
	routes.Web(app)

	// Register views
	app.RegisterView(iris.HTML("./resources/views", ".html"))

	app.Run(iris.Addr(":8080"))
}
