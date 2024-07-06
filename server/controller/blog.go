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
		context["msg"]="Something went wrong."
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
	
	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
	}

    context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg": "",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not found."
		
		return c.JSON(context)
	}
	
	resulet := database.DBConn.First(&record)

	if resulet.Error != nil {
		context["msg"] = "Something went wrong"
		return c.JSON(context)
	}

	context["statusText"] = "OK."
	context["msg"] = "Record deleted successfully."
	c.Status(200)
	return c.JSON(context)
}