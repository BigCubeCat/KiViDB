package api

import (
	"github.com/gofiber/fiber/v2"
	"kiviDB/core"
	"net/http"
)

func GetClusterHandler(context *fiber.Ctx) error {
	var id string
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		return nil
	}
	if err := core.DBCore.ClusterExists(id); err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
		return err
	}
	clusters, _ := core.DBCore.GetAll(id)
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"data":    clusters})
	return nil
}

func PostClusterHandler(context *fiber.Ctx) error {
	var id string
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		return nil
	}
	_ = core.DBCore.CreateCluster(id) // Ignore error bcs already handle cluster exists
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
	return nil
}

func DeleteClusterHandler(context *fiber.Ctx) error {
	var id string
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		return nil
	}
	if err := core.DBCore.ClusterExists(id); err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
		})
		return err
	}
	_ = core.DBCore.CreateCluster(id) // Ignore error bcs already handle cluster exists
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
	return nil
}
