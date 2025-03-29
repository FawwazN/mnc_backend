package routers

import (
	"github.com/FawwazN/mnc-backend/api/controllers"
	"github.com/FawwazN/mnc-backend/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/login", controllers.LoginController)
	api.Post("/payment", middlewares.Authenticate, controllers.PaymentController)
	api.Post("/logout", middlewares.Authenticate, controllers.LogoutController)

	api.Post("/customer/register", controllers.CreateCustomerController)
	api.Get("/customer/:id", controllers.GetCustomerByIdController)
	api.Get("/customers", controllers.GetAllCustomersController)
	api.Get("/customer/username/:username", controllers.GetCustomerByUsernameController)
	api.Put("/customer/update", controllers.UpdateCustomerController)
}
