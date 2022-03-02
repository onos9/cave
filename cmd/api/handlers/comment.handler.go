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
	// get values
	// build into struct

	var uploadBody CommentCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToComment()
	if err != nil {
		log.Printf("error in comment get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	s := utils.PrettyPrint(vid)
	fmt.Printf("%+v\n", s)
}

// CommentLoginRequest spec for login request
type CommentLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// CommentCreateRequest spec for signup request
type CommentCreateRequest struct {
	Described string     `json:"describe"`
	Video     mods.Video `json:"video"`
	User      mods.User  `json:"user"`
}

// ToComment converts CommentCreateRequest to Comment object
func (commentCreateRequest *CommentCreateRequest) ToComment() (*mods.Comment, error) {
	if commentCreateRequest == nil {
		return nil, errors.New("Null Comment Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := commentCreateRequest.CommentID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	comment := &mods.Comment{
		Described: commentCreateRequest.Described,
		Video:     commentCreateRequest.Video,
		User:      commentCreateRequest.User,
	}
	return comment, nil
}

// CommentInfoUpdateRequest - spec for updating comment info
type CommentInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
