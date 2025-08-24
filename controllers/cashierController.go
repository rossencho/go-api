package controllers

import "github.com/gofiber/fiber/v2"

func CreateCashier(c *fiber.Ctx) error {
	return nil
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
