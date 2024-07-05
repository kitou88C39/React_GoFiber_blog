package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/blog/controller"
)

func SetupRotes (app *fiber.Ctx) error {
	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/", controller.BlogUpdate)
	app.Delete("/", controller.BlogDelete())
}