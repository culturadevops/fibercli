package ejemplo

import (
	"github.com/gofiber/fiber/v2"
)

func Router(e *fiber.App, Handler *Handler) {
	e.Post("/list", Handler.ListAllFor())
}
