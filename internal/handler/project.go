package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/helper"
	"github.com/nurudin-fst/taskify/internal/repository"
	"github.com/nurudin-fst/taskify/internal/usecase"
)

type ProjectHandler struct {
	ProjectUc usecase.ProjectUC
}

func NewProjectHandler(projectUc *usecase.ProjectUC) *ProjectHandler {
	return &ProjectHandler{
		ProjectUc: *projectUc,
	}
}

func (h *ProjectHandler) Router(app *fiber.App) {
	tokenAuth := JWTAuthMiddleware()
	app.Get(`/projects`, tokenAuth, h.List)
	app.Post(`/projects`, tokenAuth, h.Insert)
	app.Put(`/project/:id`, tokenAuth, h.UpdateProject)
}

type ProjectInsertResponse struct {
	GeneralResponse
	Data dto.ProjectInsert `json:"data"`
}

func (h *ProjectHandler) Insert(c *fiber.Ctx) error {
	claims := c.Locals("user").(jwt.MapClaims)
	userId := (claims["user_id"]).(float64)
	in, err := helper.ReadBody[dto.ProjectInsertIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.ProjectUc.Insert(in, int(userId))
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res ProjectInsertResponse
	res.Data = out
	res.Message = "Project created"
	res.Code = fiber.StatusCreated

	return c.Status(fiber.StatusCreated).JSON(res)
}

type ProjectListResponse struct {
	GeneralResponse
	Data []repository.ProjectList `json:"data"`
}

func (h *ProjectHandler) List(c *fiber.Ctx) error {
	claims := c.Locals("user").(jwt.MapClaims)
	userId := (claims["user_id"]).(float64)
	out, code, err := h.ProjectUc.List(int(userId))
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res ProjectListResponse
	res.Data = out
	res.Code = fiber.StatusOK
	res.Message = "Project list obtained"

	return c.Status(fiber.StatusOK).JSON(res)
}

type ProjectUpdateResponse struct {
	GeneralResponse
	Data dto.ProjectUpdate `json:"data"`
}

func (h *ProjectHandler) UpdateProject(c *fiber.Ctx) error {
	projectId, err := c.ParamsInt("id")
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}
	in, err := helper.ReadBody[dto.ProjectUpdateIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.ProjectUc.Update(in, projectId)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res ProjectUpdateResponse
	res.Data = out
	res.Code = fiber.StatusOK
	res.Message = "Project updated"

	return c.Status(fiber.StatusOK).JSON(res)
}
