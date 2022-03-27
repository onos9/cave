package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateComment    = errors.New("Unable to create Comment")
	// errUnableToFetchComment     = errors.New("Unable to fetch comment")
	// errUnableToFetchCommentList = errors.New("Unable to fetch comment list")
	// errUnableToUpdateComment    = errors.New("Unable to update comment")
	// errUnableToDeleteComment    = errors.New("Unable to delete comment")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	comment *CommentController
)

// CommentController is an anonymous struct for comment controller
type CommentController struct{}

// SignUp registers comment
func (ctrl *CommentController) create(ctx *gin.Context) {
}
