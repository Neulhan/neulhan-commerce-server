package middleware

import (
	"github.com/gofiber/fiber/v2"
	"neulhan-commerce-server/src/jwt"
)

func NewUserMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("accessToken")
		if accessToken == "" {
			return c.Next()
		}
		var err error
		claims := &jwt.Claims{}
		claims, err = jwt.ParseToken(accessToken)
		if err != nil {
			return err
		}
		c.Locals("UserID", claims.UserID)
		return c.Next()
	}
}
