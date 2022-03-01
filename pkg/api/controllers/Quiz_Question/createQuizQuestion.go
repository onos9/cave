package quizquestion

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

	collection := db.Instance.Database.Collection("quizquestion")

	// create a new record
	quizquestion := new(models.QuizQuestion)
	quizquestion.CreatedAt = utils.MakeTimestamp()
	quizquestion.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(quizquestion); errors != nil {
		_, err := govalidator.ValidateStruct(quizquestion)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), quizquestion); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdquizquestion := &models.QuizQuestion{}
			createdRecord.Decode(createdquizquestion)

			return helpers.CrudResponse(ctx, "Create", createdquizquestion)
		}
	}
}
