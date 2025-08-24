package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rossencho/go-api/config"
	models "github.com/rossencho/go-api/models"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Failed to parse body",
			})
	}

	if data["name"] == "" || data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Name, passcode are required",
			})
	}

	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	config.DB.Create(&cashier)

	return c.Status(201).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier created",
			"data":    cashier,
		})
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}

func CashierList(c *fiber.Ctx) error {
	return c.SendString("List of cashiers")
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
