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
func PostLogin(ctx *fiber.Ctx, jm *jwt.JWTManager) error {
	access_token := jm.Generate(&jwt.TokenPayload{ID: 1001})
	return ctx.JSON(auth.AccessToken{Token: access_token})
}
