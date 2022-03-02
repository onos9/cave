package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/utils"

	"github.com/pkg/errors"

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
	// get values
	// build into struct

	var uploadBody ChannelCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToChannel()
	if err != nil {
		log.Printf("error in channel get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	s := utils.PrettyPrint(vid)
	fmt.Printf("%+v\n", s)
}

// ChannelLoginRequest spec for login request
type ChannelLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// ChannelCreateRequest spec for signup request
type ChannelCreateRequest struct {
	Name     string    `json:"name"`
	Thumnail string    `json:"thumnail"`
	Banner   string    `json:"banner"`
	About    string    `json:"about"`
	User     mods.User `json:"user"`
}

// ToChannel converts ChannelCreateRequest to Channel object
func (channelCreateRequest *ChannelCreateRequest) ToChannel() (*mods.Channel, error) {
	if channelCreateRequest == nil {
		return nil, errors.New("Null Channel Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := channelCreateRequest.ChannelID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	channel := &mods.Channel{
		Name:     channelCreateRequest.Name,
		Thumnail: channelCreateRequest.Thumnail,
		Banner:   channelCreateRequest.Banner,
		About:    channelCreateRequest.About,
		User:     channelCreateRequest.User,
	}
	return channel, nil
}

// ChannelInfoUpdateRequest - spec for updating channel info
type ChannelInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
