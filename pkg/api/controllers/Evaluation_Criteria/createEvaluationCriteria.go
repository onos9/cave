package evaluationcriteria

import (
	"github.com/asaskevich/govalidator"
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CreateNew(ctx *fiber.Ctx) error {

	collection := db.Instance.Database.Collection("evaluationcriteria")

	// create a new record
	evaluationcriteria := new(models.EvaluationCriteria)
	evaluationcriteria.CreatedAt = utils.MakeTimestamp()
	evaluationcriteria.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(evaluationcriteria); errors != nil {
		_, err := govalidator.ValidateStruct(evaluationcriteria)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), employe); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdevaluationcriteria := &models.EvaluationCriteria{}
			createdRecord.Decode(createdevaluationcriteria)

			return helpers.CrudResponse(ctx, "Create", createdevaluationcriteria)
		}
	}
}
