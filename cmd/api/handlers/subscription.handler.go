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
func (ctrl *SubscriptionController) create(ctx *gin.Context) {
	// get values
	// build into struct

	var uploadBody SubscriptionCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToSubscription()
	if err != nil {
		log.Printf("error in subscription get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	log.Printf("subscription => %+v", vid)
}

// SubscriptionLoginRequest spec for login request
type SubscriptionLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// SubscriptionCreateRequest spec for signup request
type SubscriptionCreateRequest struct {
	IsSubscribed bool         `json:"IsSubscriptiond"`
	Channel      mods.Channel `json:"channel"`
	User         mods.User    `json:"user"`
}

// ToSubscription converts SubscriptionCreateRequest to Subscription object
func (subscriptionCreateRequest *SubscriptionCreateRequest) ToSubscription() (*mods.Subscription, error) {
	if subscriptionCreateRequest == nil {
		return nil, errors.New("Null Subscription Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := subscriptionCreateRequest.SubscriptionID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	subscription := &mods.Subscription{
		IsSubscribed: subscriptionCreateRequest.IsSubscriptiond,
		Channel:      subscriptionCreateRequest.Channel,
		User:         subscriptionCreateRequest.User,
	}
	return subscription, nil
}

// SubscriptionInfoUpdateRequest - spec for updating subscription info
type SubscriptionInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
