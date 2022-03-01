package controllers

import (
	"github.com/cave/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Employe EmployeController
}

var ControllersRoutes = &Controller{}

func Index(c *fiber.Ctx) error {
	return helpers.MsgResponse(c, "Welcome to Boillerplate Fiber With Mongo", nil)
}
