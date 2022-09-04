package handlers

import (
	"mvslocal/libs"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ResponseBasicStruct struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Bind(c interface{}, r interface{}) error {
	return c.(*fiber.Ctx).BodyParser(r)
}
func Response(c interface{}, code int, data interface{}) error {
	return c.(*fiber.Ctx).Status(code).JSON(data)
}

func GetRequest(c interface{}, v interface{}) error {
	if err := Bind(c, v); err != nil {
		return err
	}
	return nil
}
func ResponseInternalServerError(c interface{}, errorText string, a ...interface{}) error {
	libs.Error(3, errorText, a...)
	return Response(c, http.StatusInternalServerError, errorText)

}
func ResponseStatusUnauthorized(c interface{}, errorText string, a ...interface{}) error {
	libs.Error(3, errorText, a...)
	return Response(c, http.StatusUnauthorized, errorText)

}
func GetParamUInt(c interface{}, idName string) uint {
	id, _ := strconv.ParseUint(c.(*fiber.Ctx).Params(idName), 10, 32)
	return uint(id)
}

func ReponseBasic(message string, data interface{}) *ResponseBasicStruct {
	r := new(ResponseBasicStruct)
	r.Message = message
	r.Data = data
	return r
}
func ReponseSuccess(c interface{}, data interface{}) error {
	return Response(c, http.StatusOK, data)
}

func ReponseSuccesswithMessage(c interface{}, message string, data interface{}) error {
	return Response(c, http.StatusOK, ReponseBasic(message, data))
}
func ResponseTradicional(c interface{}, err error, data interface{}) error {
	if err != nil {
		return ResponseInternalServerError(c, err.Error(), err)
	} else {
		return ReponseSuccess(c, data)
	}
}
