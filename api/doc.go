package api

import (
	"github.com/gofiber/fiber/v2"
	"kiviDB/core"
	"log"
	"net/http"
)

type docModel struct {
	Value string
}

type GetAndDeleteJSON struct {
	Cluster string
	Id      string
}

func GetDocumentHandler(context *fiber.Ctx) error {
	var (
		id       string
		cluster  string
		err      error
		document []byte
	)
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[GET Document] Unable to parse document id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the document",
		})
	}
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		log.Printf("[GET Document] Unable to parse cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		log.Printf("[GET Document] Unable to get document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
	}
	document, err = core.DBCore.Get(cluster, id)
	if err != nil {
		log.Printf("[GET Document] Unable to get document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"value":   string(document),
	})
}

func CreateDocumentHandler(context *fiber.Ctx) error {
	var (
		cluster     string
		err         error
		generatedId string
	)
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		log.Printf("[POST Document] Unable to parse cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		log.Printf("[POST Document] Unable to create document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
		})
	}
	model := docModel{}

	err = context.BodyParser(&model)
	if err != nil {
		log.Printf("[POST Document] Unable to create document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "unable to parse body",
		})
	}
	generatedId, err = core.DBCore.Add(cluster, []byte(model.Value))
	if err != nil {
		log.Printf("[POST Document] Unable to create document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "unable to generate new document id",
		})
	}
	log.Printf("Document with id: `%v` is created", generatedId)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"key":     generatedId,
	})
}

func PostDocumentHandler(context *fiber.Ctx) error {
	var (
		id      string
		cluster string
		err     error
	)
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		log.Printf("[POST Document] Unable to parse cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		log.Printf("[POST Document] Unable to create document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
		})
	}
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[POST Document] Unable to parse document id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the document",
		})
	}
	model := docModel{}

	err = context.BodyParser(&model)
	if err != nil {
		log.Printf("[POST Document] Unable to parse body: %v\n", err)
		return context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failure"})
	}
	err = core.DBCore.Set(cluster, id, []byte(model.Value))
	if err != nil {
		log.Printf("[POST Document] Unable to create document: %v\n", err)
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "unable to set value for document",
		})
	}
	log.Printf("Document with id: `%v` is created", id)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"key":     id,
	})
}

func DeleteDocumentHandler(context *fiber.Ctx) error {
	var (
		id      string
		cluster string
		err     error
	)
	id, err = ParseArg(context, "id")
	if err != nil {
		log.Printf("[DELETE Document] Unable to parse document id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the document",
		})
	}
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		log.Printf("[DELETE Document] Unable to parse cluster id: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to get id of the cluster",
		})
	}
	err = core.DBCore.Delete(cluster, id)
	if err != nil {
		log.Printf("[DELETE Document] Unable to delete document: %v\n", err)
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "unable to delete document",
		})
	}
	log.Printf("Document with id: `%v` is deleted", id)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
}
