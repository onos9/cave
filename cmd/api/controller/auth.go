package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/database"
	"github.com/cave/pkg/mailer"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

var userAuth *Auth

// Auth is an anonymous struct for user controller
type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Auth) signup(ctx *fiber.Ctx) error {
	var user models.User
	rdb := database.RedisClient(0)
	defer rdb.Close()

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(999999))
	userID := user.Id.Hex()

	// Hash Password
	hashedPass, _ := utils.EncryptPassword(c.Password)
	user.PasswordHash = []byte(hashedPass)
	user.Email = c.Email
	user.EnrollProgress = 0
	user.Role = "prospective"
	user.Wallet = 0
	user.UserID = userID

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err := rdb.Set(ctx.UserContext(), code, userID, 0).Err()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// mail := fiber.Map{
	// 	"fromAddress": os.Getenv("EMAIL_FROM"),
	// 	"toAddress":   c.Email,
	// 	"subject":     "Adullam|Signup",
	// 	"content": map[string]interface{}{
	// 		"filename":    "signup.html",
	// 		"paymentCode": "10-" + code,
	// 		"payment_url": os.Getenv("FINANCIAL"),
	// 	},
	// }

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	data := url.Values{}
	data.Set("userId", user.UserID)
	u, _ := url.ParseRequestURI(`https://portal.adullam.ng`)
	urlStr := u.String() + "/#/sign-in/"

	mail := fiber.Map{
		"fromAddress": "support@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Payment Confirmation",
		"content": map[string]interface{}{
			"filename":     "payment.html",
			"redirect_uri": urlStr + vt + "?" + data.Encode(),
		},
	}

	id, err := rdb.Get(ctx.UserContext(), code).Result()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"emailed": false,
			"success": true,
			"error":   err.Error(),
		})
	}

	m := new(mailer.Mail)
	_, err = m.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"ID":      id,
			"userId":  code,
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
func (c *Auth) signin(ctx *fiber.Ctx) error {
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

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    rt,
		Expires:  time.Now().Add((24 * time.Hour) * 14),
		HTTPOnly: true,
		Secure:   false,
		Domain:   os.Getenv("APP_HOST"),
	}

	ctx.Cookie(&cookie)

	roles := []string{"admin", "prospective", "guest"}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"accessToken": at,
		"user":        user,
		"roles":       roles,
		"login":       true,
	})
}

func (c *Auth) signout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Domain:   os.Getenv("APP_HOST"),
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"accessToken": nil,
		"login":       false,
		"message":     "signing out, Buy!",
	})
}

func (c *Auth) token(ctx *fiber.Ctx) error {
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

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"accessToken": t,
		"user":        user,
		"roles":       roles,
		"login":       true,
	})
}

func (c *Auth) verify(ctx *fiber.Ctx) error {
	var user models.User

	userId := ctx.Query("userId")
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

	at, err := auth.IssueAccessToken(user)
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

	roles := []string{"admin", "prospective", "guest"}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"accessToken": at,
		"user":        user,
		"userId":      userId,
		"roles":       roles,
		"login":       true,
		"isVerified":  user.IsVerified,
	})
}
