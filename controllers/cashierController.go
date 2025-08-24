package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rossencho/go-api/config"
	"github.com/rossencho/go-api/models"
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
	id := c.Params("id")
	var cashier models.Cashier

	config.DB.Select("*").Where("id = ?", id).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier details",
			"data":    cashier,
		})

}

func CashierList(c *fiber.Ctx) error {
	var cashiers []models.Cashier
	var count int64

	limit := c.QueryInt("limit", 100)
	offset := c.QueryInt("offset", 0)

	config.DB.Select("*").Limit(limit).Offset(offset).Find(&cashiers).Count(&count)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier list",
			"data":    cashiers,
		})
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
