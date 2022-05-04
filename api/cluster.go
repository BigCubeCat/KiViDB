package api

import (
	"github.com/gofiber/fiber/v2"
	"kiviDB/core"
	"log"
	"net/http"
)

func GetClusterHandler(context *fiber.Ctx) error {
	var id string
	var values []core.Object
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[GET Cluster] Unable to get cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	if err = core.DBCore.ClusterExists(id); err != nil {
		log.Printf("[GET Cluster] Unable to get cluster: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
	}
	values, err = core.DBCore.ClusterValues(id)
	if err != nil {
		log.Printf("[GET Cluster] Ubable to get cluster's values: %v\n", err)
	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"data":    values})
}

func PostClusterHandler(context *fiber.Ctx) error {
	var id string
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[POST Cluster] Unable to get cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster	",
		})
	}
	err = core.DBCore.CreateCluster(id)
	if err != nil {
		log.Printf("[POST Cluster] Ubable to create cluster: %v\n", err)
		return context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "unable to create cluster",
		})
	}
	log.Printf("Cluster with id: `%v` is created", id)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
}

func DeleteClusterHandler(context *fiber.Ctx) error {
	var id string
	var err error
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[DELETE Cluster] Unable to parse cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	if err = core.DBCore.ClusterExists(id); err != nil {
		log.Printf("[DELETE Cluster] Unable to delete cluster: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
		})

	}
	err = core.DBCore.DropCluster(id)
	if err != nil {
		log.Printf("[DELETE Cluster] Unable to delete cluster: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "unable to delete cluster",
		})
	}
	log.Printf("Cluster with id: `%v` is deleted", id)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
}
