package controller

import (
	"net/http"

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

	err := wallet.ProcessPayment(m, &user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	mail := fiber.Map{
		"fromAddress": "admin@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Adullam",
		"content": fiber.Map{
			"filename":     "payment.html",
			"redirect_uri": m["redirect_uri"].(string) + "/" + vt,
		},
	}

	mailer := new(mailer.Mail)
	_, err = mailer.SendMail(mail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
