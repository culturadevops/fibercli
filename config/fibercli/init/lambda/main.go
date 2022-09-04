package main

import (
	"context"

	"%FIBERCLI%/routes"
	"%FIBERCLI%/adapter"
	"%FIBERCLI%/handlers"
	"%FIBERCLI%/mid"
	"%FIBERCLI%/libs"
	
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var fiberLambda *adapter.FiberLambda

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
func bdconfig(){
	var c libs.MysqlDriver
	c.Configure("./configs", "mysql")
	c.NewDriver()
	libs.DB = c.GetClient()
}
func init() {



	app := fiber.New()
	app.Static("/", "./static/public")
	mainRoutes(app)
	mainMiddleware(app)
	fiberLambda = adapter.New(app)
}

// Handler proxies our app requests to aws lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
