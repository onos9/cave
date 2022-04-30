package controller

import (
	"net/http"

	"github.com/cave/pkg/zoho"
	"github.com/gofiber/fiber/v2"
)

var (
	mailer *MailerController
)

// CandidateController is an anonymous struct for candidate controller
type MailerController struct {
	Code      string `json:"code"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	AskReceip string `json:"askReceip"`
}

func (c *MailerController) send(ctx *fiber.Ctx) error {

	var mail fiber.Map

	if err := ctx.BodyParser(&mail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	m := new(zoho.Mailer)
	resp, err := m.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error)
	}
	
	return ctx.Status(http.StatusCreated).JSON(Resp{
		"respons": resp,
	})
}

func (c *MailerController) zohoCode(ctx *fiber.Ctx) error {

	m := new(zoho.Mailer)
	cfg := m.GetMailConfig()
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"credentials": cfg,
	})
}

func (c *MailerController) token(ctx *fiber.Ctx) error {

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	m := new(zoho.Mailer)
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
