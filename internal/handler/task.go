package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/helper"
	"github.com/nurudin-fst/taskify/internal/repository"
	"github.com/nurudin-fst/taskify/internal/usecase"
)

type TaskHandler struct {
	TaskUC usecase.TaskUC
}

func NewTaskHandler(taskUc *usecase.TaskUC) *TaskHandler {
	return &TaskHandler{
		TaskUC: *taskUc,
	}
}

func (h *TaskHandler) Router(app *fiber.App) {
	tokenAuth := JWTAuthMiddleware()
	app.Post(`/projects/:id/tasks`, tokenAuth, h.Insert)
	app.Get(`/projects/:id/tasks`, tokenAuth, h.List)
	app.Delete(`/task/:id`, tokenAuth, h.Delete)
	app.Put(`/task/:id`, tokenAuth, h.Update)
}

type TaskInsertResponse struct {
	GeneralResponse
	Data dto.TaskInsert `json:"data"`
}

func (h *TaskHandler) Insert(c *fiber.Ctx) error {
	projectId, err := c.ParamsInt("id")
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}
	in, err := helper.ReadBody[dto.TaskInsertIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.TaskUC.Insert(in, projectId)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res TaskInsertResponse
	res.Data = out
	res.Code = fiber.StatusCreated
	res.Message = "Task created"

	return c.Status(fiber.StatusCreated).JSON(res)
}

type TaskListResponse struct {
	GeneralResponse
	Data []repository.TaskList `json:"data"`
}

func (h *TaskHandler) List(c *fiber.Ctx) error {
	projectId, err := c.ParamsInt("id")
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.TaskUC.List(projectId)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res TaskListResponse
	res.Data = out
	res.Code = fiber.StatusOK
	res.Message = "Tasks list obtained"

	return c.Status(fiber.StatusOK).JSON(res)
}

type TaskUpdateResponse struct {
	GeneralResponse
	Data dto.TaskUpdate `json:"data"`
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	taskId, err := c.ParamsInt("id")
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}
	in, err := helper.ReadBody[dto.TaskUpdateIn](c)
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.TaskUC.Update(in, taskId)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}

	var res TaskUpdateResponse
	res.Data = out
	res.Code = fiber.StatusOK
	res.Message = "Task updated"

	return c.Status(fiber.StatusOK).JSON(res)
}

type TaskDeleteResponse struct {
	GeneralResponse
	Data dto.TaskDelete `json:"data"`
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	taskId, err := c.ParamsInt("id")
	if err != nil {
		return ResponseHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	out, code, err := h.TaskUC.Delete(taskId)
	if err != nil {
		return ResponseHttpError(c, code, err.Error())
	}
	var res TaskDeleteResponse
	res.Data = out
	res.Code = fiber.StatusOK
	res.Message = "Task deleted"

	return c.Status(fiber.StatusOK).JSON(res)
}
