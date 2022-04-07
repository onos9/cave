package controller

import (
	"net/http"

	"github.com/cave/pkg/zoho"
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateCandidate    = errors.New("Unable to create Candidate")
	// errUnableToFetchCandidate     = errors.New("Unable to fetch candidate")
	// errUnableToFetchCandidateList = errors.New("Unable to fetch candidate list")
	// errUnableToUpdateCandidate    = errors.New("Unable to update candidate")
	// errUnableToDeleteCandidate    = errors.New("Unable to delete candidate")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
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

	if err := ctx.BodyParser(&zoho.Mailer{}); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	mail := zoho.Mailer{
		ToAddress: c.To,
		Subject:   c.Subject,
		Content:   c.Body,
		AskReceip: c.AskReceip,
	}
	if err := zoho.Send(mail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error)
	}
	return ctx.Status(http.StatusCreated).JSON(Resp{
		"message": "success",
	})
}
