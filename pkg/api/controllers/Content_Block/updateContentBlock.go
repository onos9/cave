package contentblock

import (
	"time"

	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Controller) UpdateSingle(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	contentblock := new(models.ContentBlock)
	contentblockId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(contentblock)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := db.Instance.Database.Collection("contentblock")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: contentblockId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.ContentBlock{}
	rawRecord.Decode(record)

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: contentblock.Name},
				{Key: "salary", Value: contentblock.Salary},
				{Key: "age", Value: contentblock.Age},
				{Key: "updatedAt", Value: time.Now()},
			},
		},
	}
	result, updateError := collection.UpdateOne(ctx.Context(), query, update)
	if updateError != nil {
		return helpers.ServerResponse(ctx, updateError.Error(), updateError.Error())
	}

	return helpers.CrudResponse(ctx, "Update", result)
}