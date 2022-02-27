package role

import (
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)


func GetAll(ctx *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, queryError := db.Instance.Database.Collection("role").Find(ctx.Context(), query)
	if queryError != nil {
		return helpers.CrudResponse(ctx, "Get", queryError)
	}

	var role []models.Role = make([]models.Role, 0)

	// iterate the cursor and decode each item into a Todo
	if err := cursor.All(ctx.Context(), &role); err != nil {
		return helpers.MsgResponse(ctx, "get data unsuccesfully", err)
	}

	return helpers.CrudResponse(ctx, "Get", role)
}
