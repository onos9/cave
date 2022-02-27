package role

import (
	"github.com/asaskevich/govalidator"
	. "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	. "github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNew(ctx *fiber.Ctx) error {

	collection := Instance.Database.Collection("role")

	// create a new record
	role := new(Role)
	role.CreatedAt = utils.MakeTimestamp()
	role.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(role); errors != nil {
		_, err := govalidator.ValidateStruct(role)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), role); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdRole := &Role{}
			createdRecord.Decode(createdRole)

			return helpers.CrudResponse(ctx, "Create", createdRole)
		}
	}
}
