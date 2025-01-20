package controller

import (
	"frr-news/internal/api/rest/auth"
	"frr-news/internal/infra/security/jwt"

	"github.com/gofiber/fiber/v2"
)

// PostLogin godoc
// @Summary      Authentication
// @Description  Authenticate user and provide access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  auth.AccessToken
// @Failure      500  {object}  error
// @Router       /login [post]
func PostLogin(jm *jwt.JWTManager) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := jm.Generate(&jwt.TokenPayload{ID: 1001})
		if err != nil {
			return ctx.JSON(struct {
				Error string
			}{
				Error: err.Error(),
			})
		}
		return ctx.JSON(auth.AccessToken{Token: accessToken})
	}
}
