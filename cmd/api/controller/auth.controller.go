package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"github.com/cave/pkg/zoho"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	hashedPass, _ := utils.EncryptPassword(c.Password)
	user.PasswordHash = []byte(hashedPass)
	user.Email = c.Email
	user.EnrollProgress = 0
	user.Role = "prospective"

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	mail := fiber.Map{
		"fromAddress": "admin@adullam.ng",
		"toAddress":   c.Email,
		"subject":     "Activate Your Adullam Account",
		"content":     fmt.Sprintf("http://localhost:3000/#/sign-in/%s", vt),
	}

	m := new(zoho.Mailer)
	_, err = m.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"emailed": false,
			"success": true,
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"emailed": true,
		"success": true,
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
	passwordIsCorrect := utils.VerifyPassword(user.PasswordHash, c.Password)
	if !passwordIsCorrect {
		return ctx.Status(http.StatusForbidden).JSON(Resp{
			"message": "Incorrect Password",
		})
	}

	at, err := auth.IssueAccessToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	rt, err := auth.IssueRefreshToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	// m := new(zoho.Mailer)
	// cred, err := m.GetCredential()
	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	// }

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

	roles := []string{"admin", "prospective", "guest"}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		// "mail":        cred,
		"accessToken": at,
		"user":        user,
		"roles":       roles,
		"login":       true,
	})
}

func (c *AuthController) signout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"accessToken": nil,
		"login":       false,
	})
}

func (c *AuthController) token(ctx *fiber.Ctx) error {
	var user models.User

	token := ctx.Cookies("token")
	claims, err := auth.ParseToken(token)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"login": false,
		})
	}

	user.Role = claims["role"].(string)
	err = user.FetchByID(claims["userID"].(string))
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	t, err := auth.IssueAccessToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	roles := []string{"admin", "prospective", "guest"}

	// m := new(zoho.Mailer)
	// cred, err := m.GetCredential()
	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	// }

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		// "mail":        cred,
		"accessToken": t,
		"user":        user,
		"roles":       roles,
		"login":       true,
	})
}

func (c *AuthController) verify(ctx *fiber.Ctx) error {
	var user models.User

	token := ctx.Params("token")
	claims, err := auth.ParseToken(token)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"isVerified": false,
		})
	}

	user.Role = claims["role"].(string)
	user.Id, err = primitive.ObjectIDFromHex(claims["userID"].(string))
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	user.IsVerified = true
	user.UpdateOne()
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	return ctx.JSON(fiber.Map{
		"accessToken": nil,
		"login":       false,
		"isVerified":  true,
	})
}
