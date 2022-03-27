package controller

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateCategory    = errors.New("Unable to create Category")
	// errUnableToFetchCategory     = errors.New("Unable to fetch category")
	// errUnableToFetchCategoryList = errors.New("Unable to fetch category list")
	// errUnableToUpdateCategory    = errors.New("Unable to update category")
	// errUnableToDeleteCategory    = errors.New("Unable to delete category")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	category *CategoryController
)

// CategoryController is an anonymous struct for category controller
type CategoryController struct{}

// SignUp registers category
func (ctrl *CategoryController) create(ctx *fiber.Ctx) {
}
