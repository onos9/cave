package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/auth"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateVideo    = errors.New("Unable to create Video")
	// errUnableToFetchVideo     = errors.New("Unable to fetch video")
	// errUnableToFetchVideoList = errors.New("Unable to fetch video list")
	// errUnableToUpdateVideo    = errors.New("Unable to update video")
	// errUnableToDeleteVideo    = errors.New("Unable to delete video")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	video *VideoController
)

// VideoController is an anonymous struct for video controller
type VideoController struct{}

// SignUp registers video
func (ctrl *VideoController) upload(ctx *gin.Context) {
	// get values
	// build into struct

	tokenAuth, err := auth.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userid, err := auth.FetchAuth(tokenAuth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var video mods.Video
	ctx.BindJSON(&video)

	video.ID = strconv.FormatUint(userid, 10)

	err = video.Create()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error creating video",
			"error":   err,
		})
		return
	}

	err = video.FetchByID()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error can not fetch video",
			"error":   err,
		})
		return
	}

	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message":  err,
		"video_id": video.VideoID,
		"user_id": userid,
	})
}

// VideoLoginRequest spec for login request
type VideoLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// VideoCreateRequest spec for signup request
type VideoCreateRequest struct {
	Date        *time.Time
	Title       string `json:"title"`
	Thumnail    string `json:"thumnail"`
	VideoID     string `json:"videoId"`
	Ip          string `json:"ip"`
	Description string `json:"description"`
}

// ToVideo converts VideoCreateRequest to Video object
func (videoCreateRequest *VideoCreateRequest) ToVideo() (*mods.Video, error) {
	if videoCreateRequest == nil {
		return nil, errors.New("Null Video Create Request")
	}

	passwordSalt := uuid.NewRandom().String()
	saltedPassword := videoCreateRequest.VideoID + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Error generating password hash")
	}

	video := &mods.Video{
		Title:       videoCreateRequest.Title,
		Thumnail:    videoCreateRequest.Thumnail,
		Description: videoCreateRequest.Description,
		VideoID:     string(passwordHash),
	}
	return video, nil
}

// VideoInfoUpdateRequest - spec for updating video info
type VideoInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
