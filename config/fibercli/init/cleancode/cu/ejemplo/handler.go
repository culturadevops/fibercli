package ejemplo

import (
	"%FIBERCLI%/infra"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	S *Service
}

func (t *Handler) NewHandler(repo *infra.RepositorySql) error {
	service := new(Service)
	service.NewService(repo)
	t.S = service
	return nil
}

func (t *Handler) ListAllFor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
