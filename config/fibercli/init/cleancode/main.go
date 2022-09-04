package main

import (
	"%FIBERCLI%/cu/ejemplo"
	"%FIBERCLI%/infra"
	"%FIBERCLI%/infra/driver"

	"github.com/gofiber/fiber/v2"
)

var drive *infra.RepositorySql

func chooseRepositorySql(bd string) *infra.RepositorySql {
	repo := new(infra.RepositorySql)
	if bd == "mysql" {
		var c driver.MysqlDriver
		c.Configure("./configs", "mysql")
		c.NewDriver()
		repo.SetDriver(&c)
	}
	if bd == "postgres" {
		var c driver.DriverPostgres
		c.Configure("./configs", "postgres")
		c.NewDriver()
		repo.SetDriver(&c)
	}

	return repo
}
func init() {
	drive = chooseRepositorySql("mysql")
}
func router(app *fiber.App) {
	ejemploHandler := new(ejemplo.Handler)
	ejemploHandler.NewHandler(drive)
	ejemplo.Router(app, ejemploHandler)
}

func main() {
	app := fiber.New()
	router(app)
	app.Listen(":3000")
}
