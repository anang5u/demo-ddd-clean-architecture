package routes

import (
	"demo-ddd-clean-architecture/app/controller"
	"demo-ddd-clean-architecture/app/services/common"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Handle
func Handle(app *fiber.App) {
	cmn := common.Apply()
	app.Use(cors.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	api := app.Group(cmn.Config.Get("APP_ENPOINT"))

	api.Get("/info", controller.Info)
	api.Get("/inquiry/:contract_number", controller.Inquiry)
	api.Post("/payment", controller.Payment)
	api.Get("/payment/:installment_id", controller.PaymentDetail)
	api.Post("/token", controller.Token)
}
