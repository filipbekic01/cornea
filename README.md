# Cornea [work in progress]

<img src="https://ya-webdesign.com/images250_/vector-eyeball-cornea-3.png"
     alt="Cornea Go web framework"
     style="float: left; margin-right: 10px; width:90px" />

Cornea is Go web framework is based on most convinient and popular libraries such as [iris](https://github.com/kataras/iris), [gorm](https://github.com/jinzhu/gorm), [godotenv](https://github.com/joho/godotenv), [migrate](https://github.com/golang-migrate/migrate) and other. Goal is not to reinvent the wheel, but to help user take off with project as quickly as possible with a lot of basic and more advanced features out-of-box.

<div style="clear:both"></div>

## Quick Start
Download latest [release](https://github.com/filipbekic01/cornea/releases) or clone repository for nightly build:

```
$ git clone git@github.com:filipbekic01/cornea.git
```

Copy environemnt file, and set `DEBUG` to `TRUE` to avoid views and assets caching. 

```
$ copy .env.example .env
$ vim .env
```

Download latest migrate library [release](https://github.com/golang-migrate/migrate/releases) as binary and put it in Cornea root folder. Now you can fire migrations with following line.

```
$ ./migrate -path database/migrations -database postgres://username:password@localhost:5432/dbname up
```

Run the application.

```
$ go run main.go
Now listening on: http://0.0.0.0:8080
Application started. Press CTRL+C to shut down.
```


## Structure

Application structure is ment to be simple and clean. Root folder contains configuration files mostly. Folder *app* contains all data structures and methods. Once applicated is compiled, everything from this folder goes to one single binary file. Folder *assets* contains raw files which are compiled to public folder. Folder *database* contains database related files such as migrations. Folder *public* is the only one that goes on server together with main binary file.

```
├── app
│   ├── controllers
│   │   └── HomeController.go
│   ├── middleware
│   │   └── GeneralMiddleware.go
│   ├── models
│   │   └── User.go
│   ├── services
│   │   └── UserService.go
│   └── kernel.go
├── assets
│   ├── js
│   │   └── app.js
│   └── sass
│       └── app.scss
├── database
│   └── migrations
│       ├── 1574794561_create_users.down.sql
│       └── 1574794561_create_users.up.sql
├── public
│   ├── views
│   │   ├── layouts
│   │   │   └── default.html
│   │   └── home.html
│   └── favicon.ico
├── go.mod
├── LICENSE
├── main.go
├── package.json
├── README.md
└── webpack.mix.js
``` 

Once you download files, create `.env` to fit your needs. Make sure to set `DEBUG` to `TRUE` in order to disable views and assets caching on each page request.

## ORM

Default library is [gorm](https://github.com/jinzhu/gorm). Every defined model in application should be located in *app/models* folder. Since there is no model mapper for now, you have to create models on your own. Quick preview:

```
// Create
db.Create(&User{Username: "John"})

// Read
var user User
db.First(&user, 1)
db.First(&user, "username = ?", "John")

// Update
db.Model(&product).Update("Username", "Philip")

// Delete
db.Delete(&product)
```

For more detailed instructions, it would be perfect to visit their [official documentation](https://gorm.io/docs/). Make sure you skip *AutoMigrate* feature. Otherwise it may conflict with our *migrate* library.

## Migrations

Even though Cornea uses [gorm](https://github.com/jinzhu/gorm) library which has out-of-box migrations that feature looked like incomplete and bit confusing - it's hard to keep track every single migration. Therefore, using [migrate](https://github.com/golang-migrate/migrate) library Cornea migrations are way more flexible and supports more database drivers. Great thing here is that mgirations are written in clean SQL code, in files with sql extension - the way it should be.

Download their latest [release](https://github.com/golang-migrate/migrate/releases) to Cornea root folder. This is how simple usage is:

```
./migrate -path database/migrations -database postgres://username:password@localhost:5432/dbname up
```

For more detailed instructions, open up their [official cli documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

## Assets

No matter if you're making multiple or single page application, you assets should always compile to publich folder and you're supposed to use them from there. Git will ignore compiled files since you'll have different environments. Cornea uses webpack, [laravel-mix](https://laravel-mix.com/) library since they offer the most elegant syntax for minifying, uglying and moving files. In rare cases, you can write your own [webpack](https://webpack.js.org/) script.

Package [laravel-mix](https://laravel-mix.com/) gives you three commands to use. Command *watch* will recompile assets whenever you modify file. Configuration file *webpack.mix.js* is pretty simple and straightforward.

```
npm run production
npm run development
npm run watch
```





