package controllers

import (
	"github.com/FawwazN/mnc-backend/api/models"
	"github.com/FawwazN/mnc-backend/api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateCustomerController(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Balance  int    `json:"balance"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	customer := models.NewCustomer(input.Username, input.Password, input.Balance)

	err := services.CreateCustomerService(customer)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
func GetAllCustomersController(c *fiber.Ctx) error {
	customers, err := services.GetAllCustomersService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to retrieve customers"})
	}

	return c.JSON(customers)
}
func GetCustomerByIdController(c *fiber.Ctx) error {
	id := c.Params("id")
	customer, err := services.GetCustomerByIdService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "customer not found"})
	}

	return c.JSON(customer)
}
func GetCustomerByUsernameController(c *fiber.Ctx) error {
	username := c.Params("username")
	customer, err := services.GetCustomerByUsernameService(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "customer not found"})
	}
	return c.JSON(customer)
}
func UpdateCustomerController(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	err := services.UpdateCustomerService(customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}
