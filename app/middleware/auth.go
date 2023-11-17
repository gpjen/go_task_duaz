package middleware

import (
	"strings"
	"user-product-management/app/auth"
	"user-product-management/app/users"
	"user-product-management/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(userService users.UserService, authService auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		if token == "" || !strings.Contains(token, "Bearer") {
			return helper.Response(c, 401, "unauthenticated", nil)
		}

		token = strings.Replace(token, "Bearer ", "", -1)
		VerifyToken, err := authService.VerifyToken(token)

		if err != nil || !VerifyToken.Valid {
			return helper.Response(c, 401, "unauthenticated", nil)
		}

		claim := VerifyToken.Claims.(jwt.MapClaims)
		userID, _ := claim["userID"].(float64)

		user, err := userService.FindByID(uint(userID))
		if err != nil || user.ID == 0 {
			return helper.Response(c, 401, "unauthenticated", nil)
		}

		c.Locals("userID", uint(userID))

		return c.Next()
	}
}
