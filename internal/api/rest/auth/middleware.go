package auth

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"frr-news/internal/infra/config"
	"frr-news/internal/infra/security/jwt"

	"github.com/gofiber/fiber/v2"
)

// Handler is the authentication middleware
func Handler(authCfg *config.Auth) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			logrus.Debug("Authorization header missed")
			return fiber.ErrUnauthorized
		}

		chunks := strings.Split(authHeader, " ")
		if chunks[0] != "Bearer" {
			logrus.Debug(fmt.Sprintf("Authorization header format must be \"Baerer <token>\", got: %s", authHeader))
			return fiber.ErrUnauthorized
		}

		if len(chunks) < 2 {
			return fiber.ErrUnauthorized
		}

		jm := jwt.NewJWTManager(&authCfg.Jwt)
		user, err := jm.Verify(chunks[1])
		if err != nil {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("USER", user.ID)
		return ctx.Next()
	}
}
