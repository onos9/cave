package role

import (
	. "github.com/cave/configs"
	"github.com/cave/pkg/helpers"
	. "github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	roleId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	collection := Instance.Database.Collection("role")

	query := bson.D{{Key: "_id", Value: roleId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &Role{}
	rawRecord.Decode(record)

	if rawRecord.Err() != nil {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Get", record)
	}
}
