package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cave/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

var logBook *LogBook

// LogBook is an anonymous struct for logBook controller
type LogBook struct{}

func (c *LogBook) create(ctx *fiber.Ctx) error {
	var logBook models.LogBook

	if err := ctx.BodyParser(&logBook); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	//Save LogBook To DB
	if err := logBook.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "LogBook Not Registered",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
	})

}

func (c *LogBook) getOne(ctx *fiber.Ctx) error {

	var logBook models.LogBook
	id := ctx.Params("id")

	err := logBook.FetchByID(id)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"id": logBook,
	})
}

func (c *LogBook) getAll(ctx *fiber.Ctx) error {

	var logBook models.LogBook
	var logBookList models.LogBookList
	err := logBook.FetchAll(&logBookList)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"list": logBookList,
	})
}

func (c *LogBook) updateOne(ctx *fiber.Ctx) error {

	var logBook models.LogBook
	err := json.Unmarshal(ctx.Body(), &logBook)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	logBook.Id, err = primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = logBook.UpdateOne()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = logBook.FetchByID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"logBook": logBook,
	})
}

func (c *LogBook) deleteOne(ctx *fiber.Ctx) error {

	var logBook models.LogBook

	if err := ctx.BodyParser(&logBook); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err := logBook.Delete()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"deleted": logBook,
	})
}
