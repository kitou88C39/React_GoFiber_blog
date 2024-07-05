package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/blog/database"
	"github.com/neerajbg/blog/router"
	"gorm.io/gorm/logger"
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

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}