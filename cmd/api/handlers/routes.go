package handlers

import (
	"github.com/cave/pkg/auth"

	"github.com/cave/cmd/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticator *auth.Authenticator
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, auth *auth.Authenticator, db *gorm.DB) {
	models.SetRepoDB(db)
	authenticator = auth
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/ping", pingHandler)
	}
}
