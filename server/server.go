package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/blog/database"
)

func main(){
	database.ConnectDB()
	sqlDb, err := database.DBConn.DB()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {


		return c.JSON(fiber.Map{"message": "Welcome to may first web Application"})
	})
	app.Listen(":8000")
}