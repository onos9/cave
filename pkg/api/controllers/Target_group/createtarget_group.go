package target_group

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

	collection := db.Instance.Database.Collection("target_group")

	// create a new record
	target_group := new(models.Target_group)
	target_group.CreatedAt = utils.MakeTimestamp()
	target_group.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(target_group); errors != nil {
		_, err := govalidator.ValidateStruct(target_group)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), target_group); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdtarget_group := &models.Target_group{}
			createdRecord.Decode(createdtarget_group)

			return helpers.CrudResponse(ctx, "Create", createdtarget_group)
		}
	}
}
