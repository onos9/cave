package handlers

import (
	"fmt"
	"net/http"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
type UserController struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// SignUp registers user
func (ctrl *UserController) signup(ctx *gin.Context) {

	var usr mods.User
	ctx.BindJSON(&ctrl)

	passwordSalt := uuid.New().String()
	saltedPassword := ctrl.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error generating password hash",
			"error":   err,
		})
		return
	}

	usr.PasswordSalt = passwordSalt
	usr.PasswordHash = passwordHash
	usr.Email = ctrl.Email

	err = usr.Create()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error creating user",
			"error":   err,
		})
		return
	}

	ts, _ := auth.CreateToken(usr.ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = auth.CreateAuth(usr.ID, ts)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	ctx.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    ctrl,
		"token":   tokens,
	})
}

// Login user
func (ctrl *UserController) login(ctx *gin.Context) {
	ctx.BindJSON(&ctrl)

	var usr mods.User
	usr.Email = ctrl.Email
	err := usr.FetchByEmail()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", ctrl.Email),
		})
		return
	}

	byteHash := []byte(usr.PasswordHash)
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(ctrl.Password+usr.PasswordSalt))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err,
			"message": "incorrect password",
		})
		return
	}

	ts, _ := auth.CreateToken(usr.ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := auth.CreateAuth(usr.ID, ts)
	if saveErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	ctx.JSON(200, gin.H{
		"message": err,
		"tokens":  tokens,
	})
}

// SignUp registers user
func (ctrl *UserController) logout(ctx *gin.Context) {
	// get values
	// build into struct
	ctx.JSON(200, gin.H{
		"message": nil,
		"token":   "logout",
	})
}
