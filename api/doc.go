package api

import (
	"github.com/gofiber/fiber/v2"
	"kiviDB/core"
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
		return nil
	}
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		return nil
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
		return err
	}
	document, err = core.DBCore.Get(cluster, id)
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"value":   string(document),
	})
	return nil
}

func CreateDocumentHandler(context *fiber.Ctx) error {
	var (
		cluster     string
		err         error
		generatedId string
	)
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		return nil
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
		return err
	}
	model := docModel{}

	err = context.BodyParser(&model)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	generatedId, err = core.DBCore.Add(cluster, []byte(model.Value))
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": ""})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "new doc created",
		"key":     generatedId,
	})
	return nil
}

func PostDocumentHandler(context *fiber.Ctx) error {
	var (
		id      string
		cluster string
		err     error
	)
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		return nil
	}
	if err = core.DBCore.ClusterExists(cluster); err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "cluster not found",
			"data":    []string{},
		})
		return err
	}
	id = context.Params("id")
	model := docModel{}

	err = context.BodyParser(&model)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = core.DBCore.Set(cluster, id, []byte(model.Value))
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Error",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
	return nil
}

func DeleteDocumentHandler(context *fiber.Ctx) error {
	var (
		id      string
		cluster string
		err     error
	)
	id, err = ParseArg(context, "id")
	if err != nil {
		return nil
	}
	cluster, err = ParseArg(context, "cluster")
	if err != nil {
		return nil
	}
	err = core.DBCore.Delete(cluster, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
	return nil
}
