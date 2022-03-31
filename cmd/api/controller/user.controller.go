package controller

import (
	"net/http"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

var user *UserController

// UserController is an anonymous struct for user controller
type UserController struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (c *UserController) signup(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	// Hash Password
	hashedPass, _ := utils.HashPassword(c.Password)
	user.PasswordHash = []byte(hashedPass)
	user.Email = c.Email
	user.Role = c.Role

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "User Not Registered",
			"error":   err,
		})
	}

	// Issue Token
	accessToken, err := auth.IssueAccessToken(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "User Not Registered",
			"error":   err,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(Resp{
		"token": accessToken,
		"user":  user,
	})

}

// Login user
func (c *UserController) signin(ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	var user models.User
	user.Email = c.Email

	err := user.FetchByEmail()
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(Resp{
			"message": "User does not exist",
			"error":   err,
		})
	}

	// Check if Password is Correct (Hash and Compare DB Hash)
	passwordIsCorrect := utils.CheckPasswordHash(string(user.PasswordHash), c.Password)
	if !passwordIsCorrect {
		return ctx.Status(http.StatusUnauthorized).JSON(Resp{
			"message": "Incorrect Password",
		})
	}

	t, err := auth.IssueAccessToken(user)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(Resp{
		"token": t,
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

	var class models.User

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	var slice *models.UserList
	err := class.FetchAll(slice)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(slice)
}

func (c *UserController) updateOne(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err := user.UpdateOne()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(user)
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
