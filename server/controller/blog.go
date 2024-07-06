package controller

import (
	"log"

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

	var records []model.Blog

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
	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil{
		log.Println("Error in parsing request.")
		context["statusText"]=""
		context["msg"]="Something went wrong."
	}

	result := database.DBConn.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"]=""
	}
	context["msg"] = "Record is saved successully."
	context["data"] = record

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