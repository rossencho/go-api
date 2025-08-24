package routes

import (
	"github.com/gofiber/fiber/v2"
	Controllers "github.com/rossencho/go-api/controllers"
)

func Setup(app *fiber.App) {
	// app.Post("cashiers/:cashierId/login", Controllers.Login)
	// app.Get("cashiers/:cashierId/logout", Controllers.Logout)
	// app.Post("cashiers/:cashierId/passcode", Controllers.Passcode)

	app.Post("/cashiers", Controllers.CreateCashier)
	app.Get("/cashiers", Controllers.CashierList)
	app.Get("/cashiers/:id", Controllers.GetCashierById)
	app.Patch("/cashiers/:id", Controllers.UpdateCashier)
	app.Delete("/cashiers/:id", Controllers.DeleteCashier)
}
