package handler

import (
	"github.com/gofiber/fiber/v2"
)

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseHttpError(c *fiber.Ctx, code int, messageErr string) error {
	return c.Status(code).JSON(GeneralResponse{
		Code:    code,
		Message: messageErr,
	})
}
