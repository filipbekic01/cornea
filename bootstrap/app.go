package bootstrap

import (
	"github.com/filipbekic01/cornea/routes"
	"github.com/kataras/iris"
)

func App() {
	// Start iris app
	app := iris.New()

	// Define routes
	routes.Web(app)

	// Run app
	app.Run(iris.Addr(":8080"))
}
