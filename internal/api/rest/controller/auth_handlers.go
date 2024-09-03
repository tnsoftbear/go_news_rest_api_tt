package controller

import (
	"fiber-reform-rest/internal/infra/security/jwt"

	"github.com/gofiber/fiber/v2"
)

func PostLogin(ctx *fiber.Ctx, jm *jwt.JWTManager) error {
	access_token := jm.Generate(&jwt.TokenPayload{ID: 1001})
	return ctx.JSON(fiber.Map{"Token": access_token})
}
