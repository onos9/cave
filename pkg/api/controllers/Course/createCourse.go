package course

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

	collection := db.Instance.Database.Collection("course")

	// create a new record
	course := new(models.Course)
	course.CreatedAt = utils.MakeTimestamp()
	course.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(course); errors != nil {
		_, err := govalidator.ValidateStruct(course)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), course); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdcourse := &models.Course{}
			createdRecord.Decode(createdcourse)

			return helpers.CrudResponse(ctx, "Create", createdcourse)
		}
	}
}
