package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

var userAuth *AuthController

// AuthController is an anonymous struct for user controller
type AuthController struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthController) signup(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	// Hash Password
	hashedPass, _ := utils.HashPassword(c.Password)
	user.PasswordHash = []byte(hashedPass)
	user.Email = c.Email
	user.EnrollProgress = 0
	user.Role = "candidate"

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(Resp{
			"message": "User Not Registered",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(Resp{
		"message":      "success",
		"redirect_uri": fmt.Sprintf("http://localhost:3000/#/sign-in/?email=%s", c.Email),
	})

}

// Login user
func (c *AuthController) signin(ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	user.Email = c.Email

	err := user.FetchByEmail()
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(Resp{
			"message": "User does not exist",
			"error":   err.Error(),
		})
	}

	// Check if Password is Correct (Hash and Compare DB Hash)
	passwordIsCorrect := utils.CheckPasswordHash(user.PasswordHash, c.Password)
	if !passwordIsCorrect {
		return ctx.Status(http.StatusForbidden).JSON(Resp{
			"message": "Incorrect Password",
		})
	}

	t, err := auth.IssueAccessToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	rt, err := auth.IssueRefreshToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    rt,
		Expires:  time.Now().Add((24 * time.Hour) * 14),
		HTTPOnly: true,
		Secure:   false,
		Domain:   "localhost",
		SameSite: "Lax",
		Path:     "/",
	}

	ctx.Cookie(&cookie)

	roles := []string{"admin", "candidate", "guest"}

	return ctx.Status(http.StatusCreated).JSON(Resp{
		"accessToken": t,
		"user":        user,
		"roles":       roles,
	})
}

func (c *AuthController) signout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (c *AuthController) newToken(ctx *fiber.Ctx) error {

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
