package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/blog/database"
	"github.com/neerajbg/blog/model"
)

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg": "Blog List",
	}
	db := database.DBConn

	var records model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg": "Add a Blog",
	}
	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg": "Update Blog",
	}
	c.Status(200)
	return c.JSON(context)
	
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg": "Delete Blog for the given ID",
	}
	c.Status(200)
	return c.JSON(context)	
}