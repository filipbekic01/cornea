package app

import (
	"log"

	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/filipbekic01/cornea/app/middleware"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
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
func Run() *Cornea {
	cornea := new(Cornea)

	// Environment
	cornea.Environment = getEnv()

	// Iris
	cornea.Iris = iris.New()

	// Migrations
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=cornea password=postgres1")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer db.Close()

	// View
	cornea.Iris.RegisterView(iris.HTML("./public/views", ".html"))

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
