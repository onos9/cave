package target_version

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

	collection := db.Instance.Database.Collection("target_version")

	// create a new record
	target_version := new(models.Target_version)
	target_version.CreatedAt = utils.MakeTimestamp()
	target_version.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(target_version); errors != nil {
		_, err := govalidator.ValidateStruct(target_version)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), target_version); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdtarget_version := &models.Target_version{}
			createdRecord.Decode(createdtarget_version)

			return helpers.CrudResponse(ctx, "Create", createdtarget_version)
		}
	}
}
