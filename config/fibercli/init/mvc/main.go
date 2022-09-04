package main

import (
	"%FIBERCLI%/extra/appConfigs"
	"%FIBERCLI%/handlers"
	"%FIBERCLI%/mid"
	"%FIBERCLI%/routes"
	"log"

	//"%FIBERCLI%/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var AppConfig appConfigs.Web

func init() {
	AppConfig.Configure("./configs", "app")
}

func mainRoutes(app *fiber.App) {
	routes.Ejemplo(app, new(handlers.Ejemplo))
}
func mainMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(mid.NotFound)
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
}
func main() {
	app := fiber.New()
	mainRoutes(app)
	mainMiddleware(app)
	app.Static("/", "./static/public")
	log.Fatal(app.Listen(AppConfig.Port))
}
