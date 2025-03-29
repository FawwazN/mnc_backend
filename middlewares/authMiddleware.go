package middlewares

import (
	"github.com/FawwazN/mnc-backend/api/repositories"
	"github.com/FawwazN/mnc-backend/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Authenticate(c *fiber.Ctx) error {
	tokenHeader := c.Get("Authorization")
	if tokenHeader == "" {
		return unauthorizedResponse(c, "Missing or invalid token")
	}

	token := utils.ExtractToken(tokenHeader)

	session, err := repositories.FetchSession(token)
	if err != nil || session.Token == "" {
		return unauthorizedResponse(c, "Invalid token")
	}

	validationResult, err := utils.VerifyToken(token)
	if err != nil {
		repositories.DeleteSession(token)
		return unauthorizedResponse(c, "Invalid token")
	}

	c.Locals("username", validationResult.Claims["username"])
	return c.Next()
}

func unauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": message,
	})
}
