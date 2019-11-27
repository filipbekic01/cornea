package app

import (
	"log"

	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/filipbekic01/cornea/app/middleware"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Cornea .
type Cornea struct {
	Environment map[string]string
	Iris        *iris.Application
}

func getEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}

// Run .
func Run() {
	cornea := new(Cornea)

	// Load environment variables and initialize Iris framework
	cornea.Environment = getEnv()
	cornea.Iris = iris.New()

	// View engine
	var isDebug bool = false
	if cornea.Environment["DEBUG"] == "TRUE" {
		isDebug = true
	}

	cornea.Iris.RegisterView(iris.HTML("./public/views", ".html").Reload(isDebug))

	// MVC configuration
	mvc.Configure(cornea.Iris.Party("/"), func(app *mvc.Application) {
		app.Router.Use(middleware.GeneralMiddleware)

		app.Register()

		app.Handle(new(controllers.HomeController))
	})

	// Serve static files
	cornea.Iris.HandleDir("/", "./public")

	// Run iris
	cornea.Iris.Run(iris.Addr(":8080"))
}
