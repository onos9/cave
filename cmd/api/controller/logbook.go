package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cave/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
)

var practicum *LogBook

// LogBook is an anonymous struct for logBook controller
type LogBook struct{}

func (c *LogBook) create(ctx *fiber.Ctx) error {
	var logBook models.LogBook

	if err := ctx.BodyParser(&logBook); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	//Save LogBook To DB
	err := logBook.Create()
	if err == nil {
		_ = logBook.FetchByEmail()
	}

	if err != nil {
		e := err.(mongo.WriteException)
		if c := e.WriteErrors[0].Code; c == 11000 {
			_ = logBook.FetchByEmail()
		}
		err = nil
	}
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "LogBook Not Found",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"login":   true,
		"logBook": logBook,
	})

}

func (c *LogBook) getByEmail(ctx *fiber.Ctx) error {
	var logBook models.LogBook

	logBook.Email = ctx.Params("email")
	err := logBook.FetchByEmail()
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"logBook": logBook,
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
		"logBook": logBook,
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
	temp := logBook

	logBook.Id, err = primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	

	if len(temp.Evangelism) > 0 {
		t := temp.Evangelism[0]
		logBook.Evangelism = nil
		_ = logBook.FetchByID(ctx.Params("id"))
		all := append([]interface{}{t}, logBook.Evangelism...)
		logBook.Evangelism = all
	}

	if len(temp.Exercise) > 0 {
		t := logBook.Exercise[0]
		logBook.Exercise = nil
		_ = logBook.FetchByID(ctx.Params("id"))
		all := append([]interface{}{t}, logBook.Exercise...)
		logBook.Exercise = all
	}

	if len(temp.Prayer) > 0 {
		t := logBook.Prayer[0]
		logBook.Prayer = nil
		_ = logBook.FetchByID(ctx.Params("id"))
		all := append([]interface{}{t}, logBook.Prayer...)
		logBook.Prayer = all
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
