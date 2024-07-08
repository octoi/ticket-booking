package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/octoi/ticket-booking/models"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			log.Warnf("empty authorization header")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthoriized",
			})
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Warnf("Invalid token parts")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthoriized",
			})
		}

		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

		if err != nil || !token.Valid {
			log.Warnf("Invalid token")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthoriized",
			})
		}

		userId := token.Claims.(jwt.MapClaims)["id"]

		if err := db.Model(&models.User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("User not found")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthoriized",
			})
		}

		ctx.Locals("userId", userId)

		return ctx.Next()
	}
}
