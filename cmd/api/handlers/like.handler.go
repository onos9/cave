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
	// get values
	// build into struct

	var uploadBody LikeCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToLike()
	if err != nil {
		log.Printf("error in like get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	log.Printf("like => %+v", vid)
}

// LikeLoginRequest spec for login request
type LikeLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// LikeCreateRequest spec for signup request
type LikeCreateRequest struct {
	IsLiked bool       `json:"IsLiked"`
	Video   mods.Video `json:"video"`
	User    mods.User  `json:"user"`
}

// ToLike converts LikeCreateRequest to Like object
func (likeCreateRequest *LikeCreateRequest) ToLike() (*mods.Like, error) {
	if likeCreateRequest == nil {
		return nil, errors.New("Null Like Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := likeCreateRequest.LikeID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	like := &mods.Like{
		IsLiked: likeCreateRequest.IsLiked,
		Video:   likeCreateRequest.Video,
		User:    likeCreateRequest.User,
	}
	return like, nil
}

// LikeInfoUpdateRequest - spec for updating like info
type LikeInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
