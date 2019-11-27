
#  Cornea [work in progress]

Cornea is Go web framework based on most convinient and popular libraries such as [iris](https://github.com/kataras/iris), [gorm](https://github.com/jinzhu/gorm), [godotenv](https://github.com/joho/godotenv), [migrate](https://github.com/golang-migrate/migrate) and others. Goal is not to reinvent the wheel, but to help user take off with project as quickly as possible with a lot of basic and more advanced features out-of-box. Idea is to keep project community-driven. Therefore, please send all suggestions and feel free to make pull requests if you wish.

## Requirements

The only requirement you need is latest [Go](https://golang.org/) binary and proper [workspace](https://golang.org/doc/code.html) setup. If you're Go newbie, make sure you understand how workspaces and modules work.

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

Install npm modules using `npm` or `yarn` tool, then compile assets.
```
$ npm install
$ npm run development
```

Run the application.

```
$ go run main.go
Now listening on: http://0.0.0.0:8080
Application started. Press CTRL+C to shut down.
```

## Structure

Application structure is intended to be simple, clean and intuitive.

Root folder contains configuration files mostly. Folder `app` contains all data structures and methods. Everything from this folder compiles into single binary file. Folder `assets` contains raw files which are compiled and put in public folder. Folder `database` contains database related files such as migrations. Folder `public` is the only one that goes on server together with main binary file.

File `app/kernel.go` is heart of application. It initializes web framework itself as well as other related things such as routing, dependency injection, template engine, static file serving, etc...

File `.env.example` should be cloned to `.env` file. It contains all configuration variables.

There is no explicit `routes` file since they are defined right in controllers - method names define them.

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

Once compiled, files you need for server are next.
```
├── public
│   ├── js
│   ├── css
│   ├── views
│   ...
├── cornea
└── .env
``` 

## ORM

The library we decided to go with is [gorm](https://github.com/jinzhu/gorm) as they offer associations (has one, has many, belongs to, many to many, polymorphism), hooks, preloading, transactions, composite key, sql builder, logger, etc. However, we wanted to keep migrations and models away from each other. Even though there may be inconsistency in theory, this way you can keep track of every single migration. Library [gorm](https://github.com/jinzhu/gorm) does not provide as fluent as [migrate](https://github.com/golang-migrate/migrate) migrations.

You decide weather your struct will extend `gorm.Model` or not. All it has are predefined attributes such as id, created, upated and deleted timestamps. We recommend using it since it gives you soft delete out of box.

Quick preview of library usage:

```
db.Create(&User{Username: "John"})

var user User
db.First(&user, 1)
db.First(&user, "username = ?", "John")

db.Model(&product).Update("Username", "Philip")

db.Delete(&product)
```

For more detailed instructions, it would be perfect to visit their [official documentation](https://gorm.io/docs/). Make sure you ignore `AutoMigrate` feature. Otherwise it may conflict with our *migrate* library.

## Migrations

Even though Cornea uses [gorm](https://github.com/jinzhu/gorm) library which has out of box migrations, that feature looked like incomplete and bit confusing - it's hard to keep track every single migration. Therefore, using [migrate](https://github.com/golang-migrate/migrate) library Cornea migrations are way more flexible and supports more database drivers. Great thing here is that migrations are written in clean sql code, in files with sql extension - the way as it should be.

For more detailed instructions, please read their [official cli documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

## Assets

No matter if you're making multiple or single page application, you assets should always compile to publich folder and you're supposed to use them from there. Git will ignore compiled files since you'll have different environments. Cornea uses webpack, [laravel-mix](https://laravel-mix.com/) library since they offer the most elegant syntax for minifying, uglying and moving files. In rare cases, you can write your own [webpack](https://webpack.js.org/) script.

Package [laravel-mix](https://laravel-mix.com/) gives you three commands to use. Command *watch* will recompile assets whenever you modify file. Configuration file *webpack.mix.js* is pretty simple and straightforward.

```
npm run production
npm run development
npm run watch
```





