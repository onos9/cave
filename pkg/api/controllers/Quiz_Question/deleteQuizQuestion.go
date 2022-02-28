package quizquestion

import (
	. "github.com/cave/pkg/api/controllers"
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Controller) DeleteSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")

	quizquestionId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	// get collection
	collection := db.Instance.Database.Collection("quizquestion")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: quizquestionId}}
	result, deleteError := collection.DeleteOne(ctx.Context(), &query)

	if deleteError != nil {
		return helpers.ServerResponse(ctx, deleteError.Error(), deleteError.Error())
	}

	// check if item was deleted
	if result.DeletedCount < 1 {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Deleted", result)
	}
}

func DeleteAll(ctx *fiber.Ctx) error {
	// get collection
	collection := db.Instance.Database.Collection("quizquestion")

	// check if the record is there
	deleteResult := collection.Drop(ctx.Context())

	return helpers.CrudResponse(ctx, "Deleted", deleteResult)
}
