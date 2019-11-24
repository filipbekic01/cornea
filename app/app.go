package app

import (
	"log"

	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/filipbekic01/cornea/app/middleware"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Cornea .
type Cornea struct {
	Environment map[string]string
	Iris        *iris.Application
}

// Run .
func Run() *Cornea {
	cornea := new(Cornea)

	// Iris
	cornea.Iris = iris.New()

	// Environment
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cornea.Environment = env

	// View
	cornea.Iris.RegisterView(iris.HTML("./resources/views", ".html"))

	// MVC
	mvc.Configure(cornea.Iris.Party("/"), func(app *mvc.Application) {
		app.Router.Use(middleware.GeneralMiddleware)

		app.Register()

		app.Handle(new(controllers.HomeController))
	})

	// Files
	cornea.Iris.HandleDir("/", "./public")

	cornea.Iris.Run(iris.Addr(":8080"))

	return cornea
}
