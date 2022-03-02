package handlers

import (
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
	// errUnableToCreateCategory    = errors.New("Unable to create Category")
	// errUnableToFetchCategory     = errors.New("Unable to fetch category")
	// errUnableToFetchCategoryList = errors.New("Unable to fetch category list")
	// errUnableToUpdateCategory    = errors.New("Unable to update category")
	// errUnableToDeleteCategory    = errors.New("Unable to delete category")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	//ErrResetExpired = errors.New("Reset expired")

	category *CategoryController
)

// CategoryController is an anonymous struct for category controller
type CategoryController struct{}

// SignUp registers category
func (ctrl *CategoryController) create(ctx *gin.Context) {
	// get values
	// build into struct

	var uploadBody CategoryCreateRequest
	ctx.BindJSON(&uploadBody)
	vid, err := uploadBody.ToCategory()
	if err != nil {
		log.Printf("error in category get => %+v", err.Error())
	}
	//value := vid.Create()
	ctx.JSON(200, gin.H{
		"message": nil,
		"respons": "Ok!",
	})
	s := utils.PrettyPrint(vid)
	log.Printf("category => %+v", s)
}

// CategoryLoginRequest spec for login request
type CategoryLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// CategoryCreateRequest spec for signup request
type CategoryCreateRequest struct {
	Title string `json:"title"`
}

// ToCategory converts CategoryCreateRequest to Category object
func (categoryCreateRequest *CategoryCreateRequest) ToCategory() (*mods.Category, error) {
	if categoryCreateRequest == nil {
		return nil, errors.New("Null Category Create Request")
	}

	// passwordSalt := uuid.NewRandom().String()
	// saltedPassword := categoryCreateRequest.CategoryID + passwordSalt
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Error generating password hash")
	// }

	category := &mods.Category{
		Title: categoryCreateRequest.Title,
	}
	return category, nil
}

// CategoryInfoUpdateRequest - spec for updating category info
type CategoryInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}
