package controllers

import (
	"time"

	"github.com/asaskevich/govalidator"
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/helpers"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeController struct {
	Name  string
	Group string
}

func (c *EmployeController) CreateNew(ctx *fiber.Ctx) error {
	collection := db.Instance.Database.Collection("employe")

	// create a new record
	employe := new(models.Employe)
	employe.CreatedAt = utils.MakeTimestamp()
	employe.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(employe); errors != nil {
		_, err := govalidator.ValidateStruct(employe)

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
			createdemploye := &models.Employe{}
			createdRecord.Decode(createdemploye)

			return helpers.CrudResponse(ctx, "Create", createdemploye)
		}
	}
}

func (c *EmployeController) DeleteSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")

	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	// get collection
	collection := db.Instance.Database.Collection("employe")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: employeId}}
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

func (c *EmployeController) DeleteAll(ctx *fiber.Ctx) error {
	// get collection
	collection := db.Instance.Database.Collection("employe")

	// check if the record is there
	deleteResult := collection.Drop(ctx.Context())

	return helpers.CrudResponse(ctx, "Deleted", deleteResult)
}

func (c *EmployeController) GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	collection := db.Instance.Database.Collection("employe")

	query := bson.D{{Key: "_id", Value: employeId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Employe{}
	rawRecord.Decode(record)

	if rawRecord.Err() != nil {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Get", record)
	}
}

func (c *EmployeController) GetAll(ctx *fiber.Ctx) error {

	var employe []models.Employe = make([]models.Employe, 0)

	query := bson.D{{}}
	cursor, queryError := db.Instance.Database.Collection("employe").Find(ctx.Context(), query)

	if queryError != nil {
		return helpers.CrudResponse(ctx, "Get", queryError)
	}

	// iterate the cursor and decode each item into a Todo
	if err := cursor.All(ctx.Context(), &employe); err != nil {
		return helpers.MsgResponse(ctx, "get data unsuccesfully", err)
	}

	return helpers.CrudResponse(ctx, "Get", employe)
}

func (c *EmployeController) UpdateSingle(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	employe := new(models.Employe)
	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(employe)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := db.Instance.Database.Collection("employe")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: employeId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Employe{}
	rawRecord.Decode(record)

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employe.Name},
				{Key: "salary", Value: employe.Salary},
				{Key: "age", Value: employe.Age},
				{Key: "updatedAt", Value: time.Now()},
			},
		},
	}
	result, updateError := collection.UpdateOne(ctx.Context(), query, update)
	if updateError != nil {
		return helpers.ServerResponse(ctx, updateError.Error(), updateError.Error())
	}

	return helpers.CrudResponse(ctx, "Update", result)
}
