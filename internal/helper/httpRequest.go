package helper

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func ReadBody[T any | struct{}](c *fiber.Ctx) (out T, err error) {
	err = c.BodyParser(&out)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = ValidateStruct(out)
	if err != nil {
		return
	}

	return
}
