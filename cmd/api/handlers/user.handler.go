package handlers

import (
	"fmt"
	"net/http"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/auth"

	"github.com/pkg/errors"

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
type UserController struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

// SignUp registers user
func (ctrl *UserController) register(ctx *gin.Context) {

	var usr mods.User
	ctx.BindJSON(&usr)

	passwordSalt := uuid.New().String()
	saltedPassword := usr.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		errors.Wrap(err, "Error generating password hash")
	}

	usr.PasswordSalt = passwordSalt
	usr.PasswordHash = passwordHash

	value := usr.Create()
	token, _ := auth.CreateToken(usr.ID)

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   token,
	})
}

// Login user
func (ctrl *UserController) login(ctx *gin.Context) {

	var userReq mods.User
	ctx.BindJSON(&userReq)

	passwordSalt := uuid.New().String()
	saltedPassword := userReq.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		errors.Wrap(err, "Error generating password hash")
	}

	var usr mods.User
	err = usr.FetchByID()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", userReq.Username),
		})
		return
	}

	if string(usr.PasswordHash) != string(passwordHash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	//claims := auth.Claims{}

	t, err := auth.CreateToken(usr.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// saveErr := authenticator.CreateAuth(usr.ID, ts, mods.RedisClient)
	// if saveErr != nil {
	// 	ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	// }
	// tokens := map[string]string{
	// 	"access_token":  ts.AccessToken,
	// 	"refresh_token": ts.RefreshToken,
	// }

	ctx.JSON(200, gin.H{
		"message": err,
		"tokens":  t,
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
