package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

var user *User

// User is an anonymous struct for user controller
type User struct{}

func (c *User) create(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	password := utils.GeneratePassword()
	user.IsVerified = true

	// Hash Password
	hashedPass, _ := utils.EncryptPassword(password)
	user.PasswordHash = []byte(hashedPass)

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "User Not Registered",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"password": password,
	})

}

func (c *User) getOne(ctx *fiber.Ctx) error {

	var user models.User
	id := ctx.Params("id")

	err := user.FetchByID(id)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"id": user,
	})
}

func (c *User) getAll(ctx *fiber.Ctx) error {

	var user models.User
	var userList models.UserList
	err := user.FetchAll(&userList)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"list": userList,
	})
}

func (c *User) updateOne(ctx *fiber.Ctx) error {

	var user models.User
	body := ctx.Body()
	err := json.Unmarshal(body, &user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	user.Id, err = primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = user.UpdateOne()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = user.FetchByID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func (c *User) deleteOne(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err := user.Delete()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"deleted": user,
	})
}
