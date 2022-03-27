package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateChannel    = errors.New("Unable to create Channel")
	// errUnableToFetchChannel     = errors.New("Unable to fetch channel")
	// errUnableToFetchChannelList = errors.New("Unable to fetch channel list")
	// errUnableToUpdateChannel    = errors.New("Unable to update channel")
	// errUnableToDeleteChannel    = errors.New("Unable to delete channel")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	channel *ChannelController
)

// ChannelController is an anonymous struct for channel controller
type ChannelController struct{}

// SignUp registers channel
func (ctrl *ChannelController) create(ctx *gin.Context) {
}
