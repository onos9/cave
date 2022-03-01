package target

import (
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (c Controller) GetAll(ctx *fiber.Ctx) error {

	var target []models.Target = make([]models.Target, 0)

	query := bson.D{{}}
	cursor, queryError := db.Instance.Database.Collection("target").Find(ctx.Context(), query)

	if queryError != nil {
		return helpers.CrudResponse(ctx, "Get", queryError)
	}

	// iterate the cursor and decode each item into a Todo
	if err := cursor.All(ctx.Context(), &employe); err != nil {
		return helpers.MsgResponse(ctx, "get data unsuccesfully", err)
	}

	return helpers.CrudResponse(ctx, "Get", target)
}
