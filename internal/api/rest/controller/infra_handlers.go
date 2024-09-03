package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetPing(ctx *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
	}

	response := Response{Message: "pong"}
	return ctx.JSON(response)
}
