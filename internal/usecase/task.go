package usecase

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/helper"
	"github.com/nurudin-fst/taskify/internal/repository"
)

type TaskUC struct {
	TaskRepo repository.TaskRepository
}

func NewTaskUC(taskRepo *repository.TaskRepository) *TaskUC {
	return &TaskUC{
		TaskRepo: *taskRepo,
	}
}

func (uc *TaskUC) Insert(in dto.TaskInsertIn, projectId int) (out dto.TaskInsert, code int, err error) {
	deadline, err := helper.ParseDate(in.Deadline, time.RFC3339)
	if err != nil {
		return dto.TaskInsert{}, fiber.StatusBadRequest, err
	}
	task := repository.Task{
		ProjectId:   projectId,
		Title:       in.Title,
		Description: in.Description,
		Status:      "todo",
		Deadline:    deadline,
	}
	task, err = uc.TaskRepo.Insert(task)
	if err != nil {
		return dto.TaskInsert{}, fiber.StatusInternalServerError, err
	}
	out.Id = task.Id
	out.Title = task.Title
	out.Status = task.Status
	out.Deadline = task.Deadline.Format(time.DateTime)

	return
}

func (uc *TaskUC) List(projectId int) (out []repository.TaskList, code int, err error) {
	out, err = uc.TaskRepo.SelectByProject(projectId)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return
}

func (uc *TaskUC) Update(in dto.TaskUpdateIn, taskId int) (out dto.TaskUpdate, code int, err error) {
	deadline, err := helper.ParseDate(in.Deadline, time.RFC3339)
	if err != nil {
		return dto.TaskUpdate{}, fiber.StatusBadRequest, err
	}

	task := repository.Task{
		Id:          taskId,
		Title:       in.Title,
		Description: in.Description,
		Status:      in.Status,
		Deadline:    deadline,
	}
	task, err = uc.TaskRepo.Update(task)
	if err != nil {
		return dto.TaskUpdate{}, fiber.StatusInternalServerError, err
	}
	out.Id = task.Id
	out.Title = task.Title
	out.Status = task.Status
	out.Description = task.Description
	out.Deadline = task.Deadline.Format(time.DateTime)

	return
}

func (uc *TaskUC) Delete(taskId int) (out dto.TaskDelete, code int, err error) {
	task, err := uc.TaskRepo.Delete(taskId)
	if err != nil {
		return dto.TaskDelete{}, fiber.StatusInternalServerError, err
	}
	out.Id = task.Id
	out.Title = task.Title
	out.Status = task.Status
	out.Description = task.Description
	out.Deadline = task.Deadline.Format(time.DateTime)

	return
}
