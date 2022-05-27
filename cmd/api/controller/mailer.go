package controller

import (
	"net/http"

	"github.com/cave/pkg/database"
	"github.com/cave/pkg/mailer"
	"github.com/gofiber/fiber/v2"
)

var (
	mail *Mailer
)

// CandidateController is an anonymous struct for candidate controller
type Mailer struct {
	Code      string `json:"code"`
	From      string `json:"fromAddress"`
	To        string `json:"toAddress"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	AskReceip string `json:"askReceip"`
}

func (c *Mailer) send(ctx *fiber.Ctx) error {

	ml := fiber.Map{}

	if err := ctx.BodyParser(&ml); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	m := new(mailer.Mail)
	resp, err := m.SendMail(ml)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error)
	}

	cred, err := m.GetCredential()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"respons":     resp,
		"credentials": cred,
	})
}

func (c *Mailer) zohoCode(ctx *fiber.Ctx) error {

	m := new(mailer.Mail)
	cfg := m.GetMailConfig()
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"credentials": cfg,
	})
}

func (c *Mailer) token(ctx *fiber.Ctx) error {

	rdb := database.RedisClient(0)
	defer rdb.Close()

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	m := new(mailer.Mail)
	resp, err := m.RequestTokens(c.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"data": resp,
		})
	}

	rt, err := rdb.Get(ctx.UserContext(), "zohoRefreshToken").Result()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	token, err := rdb.Get(ctx.UserContext(), "accessToken").Result()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	cred, err := m.GetCredential()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"zohoRefreshToken": rt,
		"zohoAccessToken":  token,
		"mail":             cred,
	})
}
