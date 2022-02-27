package answeroption

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
  answerOption := new(models.AnswerOption)
	answerOptionId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(answerOption)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := db.Instance.Database.Collection("answerOption")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: answerOptionId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.AnswerOption{}
	rawRecord.Decode(record)

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: answerOption.Name},
				{Key: "salary", Value: answerOption.Salary},
				{Key: "age", Value: answerOption.Age},
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
