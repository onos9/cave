package controller

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateDislike    = errors.New("Unable to create Dislike")
	// errUnableToFetchDislike     = errors.New("Unable to fetch dislike")
	// errUnableToFetchDislikeList = errors.New("Unable to fetch dislike list")
	// errUnableToUpdateDislike    = errors.New("Unable to update dislike")
	// errUnableToDeleteDislike    = errors.New("Unable to delete dislike")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	dislike *DislikeController
)

// DislikeController is an anonymous struct for dislike controller
type DislikeController struct{}

// SignUp registers dislike
func (ctrl *DislikeController) create(ctx *fiber.Ctx) {
}
