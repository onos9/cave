package controller

import (
	"net/http"
	"net/url"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/mailer"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/wallet"
	"github.com/gofiber/fiber/v2"
)

var webhook *Webhook

// Webhook is an anonymous struct for user controller
type Webhook struct{}

func (c *Webhook) payment(ctx *fiber.Ctx) error {

	user := models.User{}
	m := fiber.Map{}

	if err := ctx.BodyParser(&m); err != nil {
		return ctx.Status(http.StatusOK).JSON(err)
	}

	if _, ok := m["html"]; !ok {
		return ctx.Status(http.StatusOK).JSON(ok)
	}

	id, err := wallet.ProcessPayment(m, &user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	data := new(url.Values)
	data.Set("userId", id)
	u, _ := url.ParseRequestURI(m["redirect_uri"].(string))
	u.Path = vt

	mail := fiber.Map{
		"fromAddress": "support@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Payment Confirmation",
		"content": fiber.Map{
			"filename":     "payment.html",
			"redirect_uri": u.String() + "?" + data.Encode(),
		},
	}

	mailer := new(mailer.Mail)
	resp, err := mailer.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resp,
	})
}
