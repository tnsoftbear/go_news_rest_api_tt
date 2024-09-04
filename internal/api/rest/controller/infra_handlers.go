package controller

import (
	"github.com/gofiber/fiber/v2"
)

type GetPingResponse struct{
	Message string
}

// GetPing godoc
// @Summary      Ping
// @Description  Check service health by ping http request
// @Tags         infra
// @Produce      json
// @Success      200  		{object}  controller.GetPingResponse
// @Failure		 500
// @Router       /ping [get]
func GetPing(ctx *fiber.Ctx) error {
	response := GetPingResponse{Message: "pong"}
	return ctx.JSON(response)
}
