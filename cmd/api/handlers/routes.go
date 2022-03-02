package handlers

import (
	"github.com/cave/pkg/auth"
	"github.com/pkg/errors"

	"github.com/cave/cmd/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticator   *auth.Authenticator
	ErrResetExpired = errors.New("Reset expired")
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
	apiV1.GET("/ping", pingHandler)

	userRouter := apiV1.Group("/user", user.SignUp)
	{
		userRouter.GET("/", user.SignUp)
	}

	videoRouter := apiV1.Group("/video")
	{
		videoRouter.GET("/", video.Upload)
	}

	categoryRouter := apiV1.Group("/category")
	{
		categoryRouter.GET("/", category.create)
	}

	channelRouter := apiV1.Group("/channel")
	{
		channelRouter.GET("/", channel.create)
	}

	commentRouter := apiV1.Group("/comment")
	{
		commentRouter.GET("/", comment.create)

	}
	dislikeRouter := apiV1.Group("/dislike")
	{
		dislikeRouter.GET("/", dislike.create)
	}
	likeRouter := apiV1.Group("/like")
	{
		likeRouter.GET("/", like.create)
	}
	subscriptionRouter := apiV1.Group("/subscription")
	{
		subscriptionRouter.GET("/", subscription.create)
	}
}
