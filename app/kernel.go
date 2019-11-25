package app

import (
	"log"

	"github.com/filipbekic01/cornea/app/controllers"
	"github.com/filipbekic01/cornea/app/middleware"
	"github.com/filipbekic01/cornea/database"
	"github.com/filipbekic01/cornea/database/migrations"
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

	// Migrations
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=cornea password=postgres1")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer db.Close()

	kernel := &database.Kernel{DB: db}
	createUserTable := migrations.CreateUsersTable{Kernel: kernel}
	createUserTable.Up()

	return nil

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
