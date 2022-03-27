package controller

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateUser    = errors.New("Unable to create User")
	// errUnableToFetchUser     = errors.New("Unable to fetch user")
	// errUnableToFetchUserList = errors.New("Unable to fetch user list")
	// errUnableToUpdateUser    = errors.New("Unable to update user")
	// errUnableToDeleteUser    = errors.New("Unable to delete user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	user *UserController
)

// UserController is an anonymous struct for user controller
type UserController struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	GoogleJWT string `json:"google_jwt"`
}

func (ctrl *UserController) GoogleAuth(ctx *fiber.Ctx) {

}

// SignUp registers user
func (ctrl *UserController) Signup(ctx *fiber.Ctx) {

}
