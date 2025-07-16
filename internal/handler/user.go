package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/helper"
	"github.com/nurudin-fst/taskify/internal/usecase"
)

type UserHandler struct {
	UserUC usecase.UserUC
}

func NewUserHandler(userUc *usecase.UserUC) *UserHandler {
	return &UserHandler{
		UserUC: *userUc,
	}
}

func (h *UserHandler) Router(app *fiber.App) {
	app.Post(`/register`, h.Register)
	app.Post(`/login`, h.Login)
}

type UserRegisterResponse struct {
	GeneralResponse
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	in, err := helper.ReadBody[dto.UserRegisterIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}
	code, err := h.UserUC.Register(in)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}

	var res UserRegisterResponse
	res.Code = fiber.StatusCreated
	res.Message = "User registered"

	return c.Status(fiber.StatusCreated).JSON(res)
}

type UserLoginResponse struct {
	GeneralResponse
	Data string `json:"data"`
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	in, err := helper.ReadBody[dto.UserLoginIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	token, code, err := h.UserUC.Login(in)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res UserLoginResponse
	res.Data = token
	res.Code = fiber.StatusCreated
	res.Message = "User logged in"

	return c.Status(fiber.StatusOK).JSON(res)
}
