package handlers

import (
	"%FIBERCLI%/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Ejemplo struct {
}

type ejemplorqt struct {
	//	Account  string `json:"account" form:"account" query:"account"`
}

func (t *Ejemplo) Holamundo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mdl models.Ejemplo
		data := mdl.Holamundo()
		return c.Status(http.StatusOK).JSON(data)
	}
}

/*
func (t *Ejemplo) Add() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(ejemplorqt)
		err := c.BodyParser(u)
		data := ""
		if err != nil {
			// falta el agregar model
			return c.Status(http.StatusInternalServerError).JSON(err)
		}
		return c.Status(http.StatusOK).JSON(data)
	}
}

func (t *Ejemplo) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//data := models.Ejemplo.List()
		return c.Status(http.StatusOK).JSON(data)
	}
}
func (t *Ejemplo) Show() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
		//data, _ := models.Ejemplo.Info(uint(id))
		return c.Status(http.StatusOK).JSON(data)
	}
}
func (t *Ejemplo) Del() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
		//data, _ := models.Ejemplo.Del(uint(id))
		return c.Status(http.StatusOK).JSON(data)
	}
}
func (t *Ejemplo) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
		u := new(ejemplorqt)
		if err := c.BodyParser(r); err != nil {
			return err
		}
		//models.Ejemplo.Update(uint(id),u)
		return c.Status(http.StatusOK).JSON("Registro actualizado")

	}
}*/
