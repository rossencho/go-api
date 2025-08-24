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
	app.Get("/cashiers/:cashierId", Controllers.GetCashierDetails)
	app.Patch("/cashiers/:cashierId", Controllers.UpdateCashier)
	app.Delete("/cashiers/:cashierId", Controllers.DeleteCashier)
}
