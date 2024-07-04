package controller

import (
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error{

	context := fiber.Map{}
	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error{
	context := fiber.Map{}
	c.Status(200)
	return c.JSON(context)
	
}

func BlogUpdate(c *fiber.Ctx) error{
	context := fiber.Map{}
	c.Status(200)
	return c.JSON(context)
	
}

func BlogDelete(c *fiber.Ctx) error{
	context := fiber.Map{}
	c.Status(200)
	return c.JSON(context)
	
}