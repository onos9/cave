package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/utils"
	"github.com/gin-gonic/gin"
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
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	GoogleJWT string `json:"google_jwt"`
}

func (c *CandidateController) setRoute(r *gin.RouterGroup) {
	candRouter := r.Group("/candidate")
	candRouter.POST("/", c.create)
	candRouter.GET("/", c.getAll)
	candRouter.POST("/bio", c.bio)
	candRouter.POST("/background", c.background)
	candRouter.POST("/health", c.health)
	candRouter.POST("/qualification", c.qualification)
	candRouter.POST("/terms", c.terms)
	candRouter.POST("/referee", c.referee)

	candRouter.POST("/google", c.google)
	candRouter.POST("/login", c.login)
	candRouter.POST("/logout", c.logout)
}

func (c *CandidateController) google(ctx *gin.Context) {
	ctx.BindJSON(&c)

	// Validate the JWT is valid
	claims, err := auth.ValidateGoogleJWT(c.GoogleJWT)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Invalid google auth",
		})
		return
	}

	if claims.Email != c.Email {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Emails don't match",
		})
		return
	}

	var cand mods.Candidate
	cand.Email = c.Email
	err = cand.UpdateOrCreateByEmail()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", c.Email),
		})
		return
	}

	ts, _ := auth.CreateToken(cand.ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err,
			"message": "Couldn't make authentication token",
		})
		return
	}

	saveErr := auth.CreateAuth(cand.ID, ts)
	if saveErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "Couldn't make authentication token",
		})
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"tokens":  tokens,
	})
}

// Login user
func (ctrl *CandidateController) login(ctx *gin.Context) {
	ctx.BindJSON(&ctrl)

	var cand mods.Candidate
	cand.Email = ctrl.Email
	err := cand.FetchByEmail()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", ctrl.Email),
		})
		return
	}

	byteHash := []byte(cand.PasswordHash)
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(ctrl.Password+cand.PasswordSalt))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err,
			"message": "incorrect password",
		})
		return
	}

	ts, _ := auth.CreateToken(cand.ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := auth.CreateAuth(cand.ID, ts)
	if saveErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"tokens":  tokens,
	})
}

// SignUp registers user
func (c *CandidateController) logout(ctx *gin.Context) {
	au, err := auth.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "unauthorized",
		})
		return
	}

	deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "unauthorized",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func (ctrl *CandidateController) create(ctx *gin.Context) {

	var cand mods.Candidate

	ctx.BindJSON(&ctrl)
	pass, err := password.Generate(8, 0, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	mail := utils.Mail{}
	mail.Send()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("unable to send mail to %s", cand.Email),
			"error":   err,
		})
		return
	}

	passwordSalt := uuid.New().String()
	saltedPassword := pass + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

	cand.PasswordSalt = passwordSalt
	cand.PasswordHash = passwordHash
	cand.Email = ctrl.Email
	if !cand.IsVerified {
		cand.IsVerified = true
	}

	err = cand.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": cand,
		"link":      fmt.Sprintf("http://locahost:33000/%s", ctrl.Password),
	})
}

func (ctrl *CandidateController) getAll(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	var cands mods.CandidateList
	err = cand.FetchAll(&cands)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": cands,
	})
}

// create bio data
func (ctrl *CandidateController) bio(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	bio := cand.Bio
	err = bio.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": bio,
	})
}

// create bio data
func (ctrl *CandidateController) qualification(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	qualification := cand.Qualification
	err = qualification.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": qualification,
	})
}

// create bio data
func (ctrl *CandidateController) background(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	background := cand.Background
	err = background.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": background,
	})
}

// create bio data
func (ctrl *CandidateController) health(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	health := cand.Health
	err = health.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": health,
	})
}

// create bio data
func (ctrl *CandidateController) referee(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("bio data for %s not found", data),
			"error":   err,
		})
		return
	}

	referee := cand.Referee
	err = referee.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("can not creat data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": referee,
	})
}

// create bio data
func (ctrl *CandidateController) terms(ctx *gin.Context) {

	data, _ := ioutil.ReadAll(ctx.Request.Body)

	var cand mods.Candidate
	err := json.Unmarshal(data, &cand)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "unable to unmarshal data",
			"error":   err,
		})
		return
	}

	terms := cand.Terms
	err = terms.Create()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("terms data for %s not found", cand.Email),
			"error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "success",
		"candidate": terms,
	})
}
