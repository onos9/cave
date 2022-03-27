package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateLike    = errors.New("Unable to create Like")
	// errUnableToFetchLike     = errors.New("Unable to fetch like")
	// errUnableToFetchLikeList = errors.New("Unable to fetch like list")
	// errUnableToUpdateLike    = errors.New("Unable to update like")
	// errUnableToDeleteLike    = errors.New("Unable to delete like")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	like *LikeController
)

// LikeController is an anonymous struct for like controller
type LikeController struct{}

// SignUp registers like
func (ctrl *LikeController) create(ctx *gin.Context) {
}
