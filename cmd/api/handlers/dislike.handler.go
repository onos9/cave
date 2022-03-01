package handlers

import (
	"log"
	"time"

	"github.com/cave/cmd/api/mods"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
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
func (ctrl *DislikeController) create(ctx *gin.Context) {
	// get values
	// build into struct

	var uploadBody DislikeCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToDislike()
	if err != nil {
		log.Printf("error in dislike get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	log.Printf("dislike => %+v", vid)
}

// DislikeLoginRequest spec for login request
type DislikeLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// DislikeCreateRequest spec for signup request
type DislikeCreateRequest struct {
	IsDisliked bool       `json:"describe"`
	Video      mods.Video `json:"video"`
	User       mods.User  `json:"user"`
}

// ToDislike converts DislikeCreateRequest to Dislike object
func (dislikeCreateRequest *DislikeCreateRequest) ToDislike() (*mods.Dislike, error) {
	if dislikeCreateRequest == nil {
		return nil, errors.New("Null Dislike Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := dislikeCreateRequest.DislikeID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	dislike := &mods.Dislike{
		IsDisliked: dislikeCreateRequest.IsDisliked,
		Video:      dislikeCreateRequest.Video,
		User:       dislikeCreateRequest.User,
	}
	return dislike, nil
}

// DislikeInfoUpdateRequest - spec for updating dislike info
type DislikeInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
