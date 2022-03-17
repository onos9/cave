package handlers

import (
	"github.com/cave/pkg/database"
	"github.com/pkg/errors"

	"github.com/cave/cmd/api/mods"
	"github.com/gin-gonic/gin"
)

var (
	//authenticator   *auth.Authenticator
	ErrResetExpired = errors.New("Reset expired")
)

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, db *database.Database) {
	mods.SetRepoDB(db)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", user.signup)
		auth.POST("/", user.googleAuth)
	}

	apiV1 := r.Group("/api/v1")
	candidate.setRoute(apiV1)

	userRouter := apiV1.Group("/user")
	{
		userRouter.GET("/", user.signup)
	}

	videoRouter := apiV1.Group("/video")
	{
		videoRouter.POST("/", video.upload)
	}

	categoryRouter := apiV1.Group("/category")
	{
		categoryRouter.GET("/", category.create)
	}

	channelRouter := apiV1.Group("/channel")
	{
		channelRouter.POST("/", channel.create)
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
