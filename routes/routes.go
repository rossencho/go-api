package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rossencho/go-api/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers", controllers.CreateCashier)
	app.Get("/cashiers", controllers.CashierList)
	app.Get("/cashiers/:id", controllers.GetCashierById)
	app.Patch("/cashiers/:id", controllers.UpdateCashier)
	app.Delete("/cashiers/:id", controllers.DeleteCashier)
}
