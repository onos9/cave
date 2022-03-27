package controller

import (
	"net/http"

	z "github.com/cave/pkg/zoho"
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
	Code string `json:"code"`
}

func (c *MailerController) code(ctx *fiber.Ctx) error {

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := z.Mail.RequestTokens(c.Code); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(z.Mail.Respons)
}

func (c *MailerController) credential(ctx *fiber.Ctx) error {
	if err := z.Mail.GetCredential(); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusCreated).JSON(z.Mail.Respons)
}

func (c *MailerController) token(ctx *fiber.Ctx) error {
	if err := z.Mail.GetNewToken(); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(z.Mail.Respons)
}

func (c *MailerController) sendMail(ctx *fiber.Ctx) error {

	if err := ctx.BodyParser(&z.Mail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := z.Mail.Send(); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusCreated).JSON(z.Mail.Respons)
}
