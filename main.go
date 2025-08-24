package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/rossencho/go-api/config"
	routes "github.com/rossencho/go-api/routes"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(app)

	routes.Setup(app)
	app.Listen(":8080")
}
