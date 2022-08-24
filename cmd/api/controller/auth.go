package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cave/config"
	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/mail"
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
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

func (c *Auth) signup(ctx *fiber.Ctx) error {
	var user models.User
	rdb := config.RedisClient(0)
	defer rdb.Close()

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(999999))

	// Hash Password
	hashedPass, _ := utils.EncryptPassword(c.Password)
	user.PasswordHash = []byte(hashedPass)
	user.Email = c.Email
	user.Role = c.Role
	user.EnrollProgress = 0
	user.Wallet = 0
	user.UserID = code
	user.FullName = c.FullName

	//Save User To DB
	if err := user.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	////////////////////////////////////////////////////////////////
	//
	// REDIS IS GIVEN THE FOLLOWING ERROR ON THE SERVER "READONLY You can't write against a read only replica."
	// RESOLVE THIS ERROR BEFOR ENABLING THE FOLLOWING REDIS CODE
	//
	// POSIBLE SOLUTIONS:
	//
	// Because the master-slave replication cluster was configured before, the configuration was changed disorderly
	// There are two solutions:
	// 1. Open the configuration file corresponding to the redis service,
	// 	  and change the value of the attribute slave read-only to no, so that it can be written.
	// 2. Open the client mode through the redis cli command, and enter the slave of no one command
	//
	// https://programmerah.com/error-readonly-you-cant-write-against-a-read-only-replica-43956/
	//
	/////////////////////////////////////////////////////////////////

	// err := rdb.Set(ctx.UserContext(), code, user.Id.Hex(), 0).Err()
	// if err != nil {
	// 	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
	// 		"success": false,
	// 		"error":   err.Error(),
	// 	})
	// }

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

	if t := ctx.Params("type"); t == "logbook" {
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
		}

		ctx.Cookie(&cookie)

		roles := []string{"admin", "prospective", "guest", "student"}

		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"accessToken": at,
			"user":        user,
			"roles":       roles,
			"login":       true,
		})
	}

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	query := url.Values{}
	query.Set("reg_tk", vt)
	u, _ := url.ParseRequestURI(os.Getenv("APP_HOST"))
	urlStr := u.String() + "/#/platform/sign-in/register"

	data := fiber.Map{
		"fromAddress": "support@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Payment Confirmation",
		"content": map[string]interface{}{
			"filename":     "payment.html",
			"redirect_uri": urlStr + "?" + query.Encode(),
		},
	}

	// id, err := rdb.Get(ctx.UserContext(), code).Result()
	// if err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"emailed": false,
	// 		"success": true,
	// 		"error":   err.Error(),
	// 	})
	// }

	m := mail.Mail{}
	_, err = m.SendMail(data)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"emailed": false,
			"success": true,
			"error":   err,
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
