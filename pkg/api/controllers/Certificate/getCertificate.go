package certificate

import (
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Controller) GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	certificateId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	collection := db.Instance.Database.Collection("certificate")

	query := bson.D{{Key: "_id", Value: certificateId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Certificate{}
	rawRecord.Decode(record)

	if rawRecord.Err() != nil {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Get", record)
	}
}
