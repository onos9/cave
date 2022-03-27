package controller

import (
	"log"
	"net/http"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

var (
	// errAuthenticationFailure = errors.New("Authentication failed")
	// errorNotFound            = errors.New("Entity not found")
	// errForbidden             = errors.New("Attempted action is not allowed")
	// errUnableToCreateCandidate    = errors.New("Unable to create Candidate")
	// errUnableToFetchCandidate     = errors.New("Unable to fetch candidate")
	// errUnableToFetchCandidateList = errors.New("Unable to fetch candidate list")
	// errUnableToUpdateCandidate    = errors.New("Unable to update candidate")
	// errUnableToDeleteCandidate    = errors.New("Unable to delete candidate")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	candidate *CandidateController
)

// CandidateController is an anonymous struct for candidate controller
type CandidateController struct {
	mods.Candidate
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	GoogleJWT string `json:"google_jwt"`
}

func (c *CandidateController) setRoute(r *gin.RouterGroup) {
	// candRouter := r.Group("/candidate")

	// candRouter.POST("/google", c.google)
	// candRouter.POST("/login", c.login)
	// candRouter.POST("/logout", c.logout)

}

func (c *CandidateController) register(ctx *fiber.Ctx) error {
	var candidate mods.Candidate

	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	// Hash Password
	hashedPass, _ := utils.HashPassword(c.Password)
	candidate.PasswordHash = []byte(hashedPass)

	// Save User To DB
	// if err := candidate.Create(); err != nil {
	// 	response := HTTPResponse(http.StatusInternalServerError, "User Not Registered", err.Error())
	// 	return ctx.JSON(response)
	// }

	response := HTTPResponse(http.StatusCreated, "User Registered", candidate)
	return ctx.Status(http.StatusCreated).JSON(response)

}

func (c *CandidateController) google(ctx *fiber.Ctx) error {

	// Validate Input
	if err := utils.ParseBodyAndValidate(ctx, &c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	// Validate the JWT is valid
	claims, err := auth.ValidateGoogleJWT(c.GoogleJWT)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if claims.Email != c.Email {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	var candidate mods.Candidate
	candidate.Email = c.Email
	err = candidate.Create()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	ts, _ := auth.CreateToken(candidate.ID.String())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	saveErr := auth.CreateAuth(candidate.ID.String(), ts)
	if saveErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	return ctx.Status(http.StatusOK).JSON(tokens)
}

// Login user
func (c *CandidateController) login(ctx *fiber.Ctx) error {
	// Validate Input
	if err := utils.ParseBodyAndValidate(ctx, &c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	var cand mods.Candidate
	cand.Email = c.Email

	// Check if Password is Correct (Hash and Compare DB Hash)
	passwordIsCorrect := utils.CheckPasswordHash(string(cand.PasswordHash), c.Password)
	if !passwordIsCorrect {
		errorList = nil
		errorList = append(
			errorList,
			&Response{
				Code:    http.StatusUnauthorized,
				Message: "Email or Password is Incorrect",
			},
		)
		return ctx.Status(http.StatusUnauthorized).JSON(HTTPErrorResponse(errorList))
	}

	ts, err := auth.CreateToken(cand.ID.String())
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
	}

	err = auth.CreateAuth(cand.ID.String(), ts)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	return ctx.Status(http.StatusCreated).JSON(tokens)
}

// SignUp registers user
func (c *CandidateController) logout(ctx *fiber.Ctx) error {
	// au, err := auth.ExtractTokenMetadata(ctx.Request())
	// if err != nil {
	// 	return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
	// }

	// deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	// if delErr != nil || deleted == 0 { //if any goes wrong
	// 	return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
	// }

	return ctx.Status(http.StatusOK).JSON("success")
}

func (c *CandidateController) create(ctx *fiber.Ctx) error {

	var cand mods.Candidate

	// Validate Input
	if err := utils.ParseBodyAndValidate(ctx, &c); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	pass, err := password.Generate(8, 0, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	passwordSalt := uuid.New().String()
	saltedPassword := pass + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	cand.PasswordSalt = passwordSalt
	cand.PasswordHash = passwordHash
	cand.Email = "user@gmail.com"

	err = cand.Create()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	// if err := z.Mail.Send(); err != nil {
	// 	return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	// }

	return ctx.Status(http.StatusOK).JSON(cand)
}

func (c *CandidateController) getAll(ctx *fiber.Ctx) error {

	route := ctx.Route()

	return ctx.Status(http.StatusOK).JSON(route)
}
