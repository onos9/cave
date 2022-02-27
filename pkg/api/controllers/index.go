package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cave/pkg/helpers"
)

func Index(c *fiber.Ctx) error {
	return helpers.MsgResponse(c, "Welcome to Boillerplate Fiber With Mongo", nil)
}