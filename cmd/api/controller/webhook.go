package controller

import (
	"net/http"
	"net/url"

	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/mail"
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

	err := wallet.ProcessPayment(m, &user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	
	if user.UserID == "" {
		naration := m["TransactionNarration"].(string)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid payment code: \"" + naration + "\"",
		})
	}

	vt, err := auth.IssueVerificationToken(user)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(err.Error())
	}

	query := url.Values{}
	query.Set("userId", user.UserID)
	u, _ := url.ParseRequestURI(m["redirect_uri"].(string))
	urlStr := u.String() + "/#/sign-in/"

	data := fiber.Map{
		"fromAddress": "support@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Payment Confirmation",
		"content": map[string]interface{}{
			"filename":     "payment.html",
			"redirect_uri": urlStr + vt + "?" + query.Encode(),
		},
	}

	mail := new(mail.Mail)
	resp, err := mail.SendMail(data)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resp,
	})
}
