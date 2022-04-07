package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

var user *UserController

// UserController is an anonymous struct for user controller
type UserController struct{}

func (c *UserController) create(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	password := utils.GeneratePassword()
	user.IsVerified = true

	// Hash Password
	hashedPass, _ := utils.HashPassword(password)
	user.PasswordHash = []byte(hashedPass)

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "User Not Registered",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(Resp{
		"password": password,
	})

}

func (c *UserController) getOne(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err := user.FetchByID()
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

func (c *UserController) getAll(ctx *fiber.Ctx) error {

	var user models.User
	var userList models.UserList
	err := user.FetchAll(&userList)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(userList)
}

func (c *UserController) updateOne(ctx *fiber.Ctx) error {

	var user models.User
	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user.Id, err = primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return err
	}

	err = user.UpdateOne()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON("success")
}

func (c *UserController) deleteOne(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err := user.Delete()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(user)
}
