package handlers

import (
	"fmt"
	"log"
	"net/http"

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
	// errUnableToCreateUser    = errors.New("Unable to create User")
	// errUnableToFetchUser     = errors.New("Unable to fetch user")
	// errUnableToFetchUserList = errors.New("Unable to fetch user list")
	// errUnableToUpdateUser    = errors.New("Unable to update user")
	// errUnableToDeleteUser    = errors.New("Unable to delete user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	user *UserController
)

// UserController is an anonymous struct for user controller
type UserController struct{}

// SignUp registers user
func (ctrl *UserController) register(ctx *gin.Context) {

	var usr mods.User
	ctx.BindJSON(&userReq)

	passwordSalt := uuid.NewRandom().String()
	saltedPassword := userReq.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Error generating password hash")
	}

	usr{
		passwordHash: passwordHash,
		passwordSalt: passwordSalt,
	}

	value := usr.Create()
	token, _ := authenticator.GenerateToken(auth.Claims{})

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   token,
	})
}

// Login user
func (ctrl *UserController) login(ctx *gin.Context) {

	userReq := mods.User
	ctx.BindJSON(&userReq)

	passwordSalt := uuid.NewRandom().String()
	saltedPassword := userReq.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Error generating password hash")
	}

	usr := mods.User
	err = usr.FetchByID()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", req.Username),
		})
		return
	}

	if usr.PasswordHash != passwordHash {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	token, err := authenticator.GenerateToken(auth.Claims{})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": err,
		"token":   token,
	})
}

// UserLoginRequest spec for login request
type UserLogoutRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// SignUp registers user
func (ctrl *UserController) logout(ctx *gin.Context) {
	// get values
	// build into struct

	var signupBody UserCreateRequest
	ctx.BindJSON(&signupBody)
	usr, err := signupBody.ToUser()
	if err != nil {
		log.Printf("error in user get => %+v", err.Error())
	}
	value := usr.Create()
	token, _ := authenticator.GenerateToken(auth.Claims{})

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   token,
	})
}
