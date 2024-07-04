package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/blog/database"
)

func init(){
	database.ConnectDB()
}
func main(){
	
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}
	defer sqlDb.Close()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {


		return c.JSON(fiber.Map{"message": "Welcome to may first web Application"})
	})
	app.Listen(":8000")
}