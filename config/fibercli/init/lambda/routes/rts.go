package routes

import (
	"%FIBERCLI%/handlers"

	"github.com/gofiber/fiber/v2"
)

func Ejemplo(e *fiber.App, Handler *handlers.Ejemplo) {
	e.Get("/index", func(c *fiber.Ctx) error {
		return c.SendFile("index")
	})
	r := e.Group("/ejemplo")
	r.Get("/holamundo", Handler.Holamundo())

	/*r.Get("/show/:id", Handler.Show())
	r.Post("/list", Handler.Add())
	r.Put("/update/:id", Handler.Update())
	r.Delete("/del/:id", Handler.Del())
	*/
}
