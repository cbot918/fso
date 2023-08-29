package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	App *fiber.App
}

func NewAPI(app *fiber.App) *API {

	api := new(API)
	app.Use(logger.New())
	app.Use(recover.New())

	api.App = app

	api.setupRouter()

	return api
}

func (api *API) setupRouter() {

	api.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
