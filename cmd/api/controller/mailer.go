package controller

import (
	"net/http"

	"github.com/cave/pkg/mailer"
	"github.com/gofiber/fiber/v2"
)

var (
	mail *Mailer
)

// CandidateController is an anonymous struct for candidate controller
type Mailer struct {
	Code      string `json:"code"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	AskReceip string `json:"askReceip"`
}

func (c *Mailer) send(ctx *fiber.Ctx) error {

	var mail fiber.Map

	if err := ctx.BodyParser(&mail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	m := new(mailer.Mail)
	resp, err := m.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"respons": resp,
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

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	m := new(mailer.Mail)
	token, err := m.RequestTokens(c.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	cred, err := m.GetCredential()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"token": token,
		"mail":  cred,
	})
}
