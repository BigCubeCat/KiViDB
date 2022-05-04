package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ParseArg(context *fiber.Ctx, arg string) (string, error) {
	param := context.Params(arg)
	if param == "" {
		return "", errors.New("value of param cannot be empty")
	}
	return param, nil
}
