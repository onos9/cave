package target_group

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
	target_group := new(models.Target_group)
	target_groupId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(target_group)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := db.Instance.Database.Collection("target_group")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: target_groupId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Target_group{}
	rawRecord.Decode(record)

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: target_group.Name},
				{Key: "salary", Value: target_group.Salary},
				{Key: "age", Value: target_group.Age},
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
