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
	app.Listen(":8000")


		
	app.Listen(":8000")
}