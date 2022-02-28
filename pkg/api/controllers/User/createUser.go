package user

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

	collection := db.Instance.Database.Collection("user")

	// create a new record
	user := new(models.User)
	user.CreatedAt = utils.MakeTimestamp()
	user.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(user); errors != nil {
		_, err := govalidator.ValidateStruct(user)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), user); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createduser := &models.User{}
			createdRecord.Decode(createduser)

			return helpers.CrudResponse(ctx, "Create", createduser)
		}
	}
}
