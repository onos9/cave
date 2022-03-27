package controller

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure  = errors.New("Authentication failed")
	// errorNotFound             = errors.New("Entity not found")
	// errForbidden              = errors.New("Attempted action is not allowed")
	// errUnableToCreateVideo    = errors.New("Unable to create Video")
	// errUnableToFetchVideo     = errors.New("Unable to fetch video")
	// errUnableToFetchVideoList = errors.New("Unable to fetch video list")
	// errUnableToUpdateVideo    = errors.New("Unable to update video")
	// errUnableToDeleteVideo    = errors.New("Unable to delete video")

	//ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	video *VideoController
)

// VideoController is an anonymous struct for video controller
type VideoController struct{}

// SignUp registers video
func (ctrl *VideoController) upload(ctx *fiber.Ctx) {

}
