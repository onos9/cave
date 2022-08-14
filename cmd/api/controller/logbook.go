package controller

import (
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
	err := logBook.Create()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "LogBook Not Found",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
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

func (c *LogBook) getOneByUserId(ctx *fiber.Ctx) error {

	var logBook models.LogBook
	logBook.UserID = ctx.Params("userId")

	logBookList := models.LogBookList{}
	err := logBook.FetchByUserId(&logBookList)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	if len(logBookList) == 0 {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"success": false,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"list":    logBookList,
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
	if err := ctx.BodyParser(&logBook); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	Id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	temp := logBook
	logBook = models.LogBook{}
	_ = logBook.FetchByID(ctx.Params("id"))

	if t := temp.Evangelism; t != nil {
		logBook.Evangelism = append([]*models.Evangelism{t[0]}, logBook.Evangelism...)
	}

	if t := temp.Prayer; t != nil {
		logBook.Prayer = append([]*models.Prayer{t[0]}, logBook.Prayer...)
	}

	if t := temp.Exercise; t != nil {
		logBook.Exercise = append([]*models.Exercise{t[0]}, logBook.Exercise...)
	}

	logBook.Id = Id
	logBook.Status = temp.Status
	err = logBook.UpdateOne()
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

func (c *LogBook) updateMany(ctx *fiber.Ctx) error {

	var logBook models.LogBook
	if err := ctx.BodyParser(&logBook); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	var status string
	if logBook.Status == "Open" {
		status = "Closed"
	} else {
		status = "Open"
	}

	err := logBook.UpdateManyWhere("status", status)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	var logBookList models.LogBookList
	err = logBook.FetchAll(&logBookList)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"list": logBookList,
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
