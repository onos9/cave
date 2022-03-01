package issuedcertificate

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

	collection := db.Instance.Database.Collection("issuedcertificate")

	// create a new record
	isssuedcertificate := new(models.IssuedCertificate)
	issuedcertificate.CreatedAt = utils.MakeTimestamp()
	issuedcertificate.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(issuedcertificate); errors != nil {
		_, err := govalidator.ValidateStruct(issuedcertificate)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), issuedcertificate); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdissuedcertificate := &models.IssuedCertificate{}
			createdRecord.Decode(createdissuedcertificate)

			return helpers.CrudResponse(ctx, "Create", createdissuedcertificate)
		}
	}
}
