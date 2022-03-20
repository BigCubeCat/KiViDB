package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ParseArg(context *fiber.Ctx, arg string) (string, error) {
	id := context.Params(arg)
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": arg + " cannot be empty",
			"data":    []string{},
		})
		return "", errors.New("")
	}
	return id, nil
}
