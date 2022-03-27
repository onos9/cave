package controller

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateSubscription    = errors.New("Unable to create Subscription")
	// errUnableToFetchSubscription     = errors.New("Unable to fetch subscription")
	// errUnableToFetchSubscriptionList = errors.New("Unable to fetch subscription list")
	// errUnableToUpdateSubscription    = errors.New("Unable to update subscription")
	// errUnableToDeleteSubscription    = errors.New("Unable to delete subscription")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	subscription *SubscriptionController
)

// SubscriptionController is an anonymous struct for subscription controller
type SubscriptionController struct{}

// SignUp registers subscription
func (ctrl *SubscriptionController) create(ctx *fiber.Ctx) {
}
