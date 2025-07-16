package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/repository"
)

type ProjectUC struct {
	ProjectRepo repository.ProjectRepository
}

func NewProjectUC(projectRepo *repository.ProjectRepository) *ProjectUC {
	return &ProjectUC{
		ProjectRepo: *projectRepo,
	}
}

func (uc *ProjectUC) Insert(in dto.ProjectInsertIn, userId int) (out dto.ProjectInsert, code int, err error) {
	project := repository.Project{
		Name:        in.Name,
		Description: in.Description,
		CreatedBy:   userId,
	}
	project, err = uc.ProjectRepo.Insert(project)
	if err != nil {
		return dto.ProjectInsert{}, fiber.StatusInternalServerError, err
	}
	out.Id = project.Id
	out.Name = project.Name

	return out, code, err
}

func (uc *ProjectUC) List(userId int) (out []repository.ProjectList, code int, err error) {
	out, err = uc.ProjectRepo.List(userId)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return
}

func (uc *ProjectUC) Update(in dto.ProjectUpdateIn, projectId int) (out dto.ProjectUpdate, code int, err error) {
	project := repository.Project{
		Id:          projectId,
		Name:        in.Name,
		Description: in.Description,
	}
	project, err = uc.ProjectRepo.Update(project)
	out.Id = project.Id
	out.Name = project.Name
	out.Description = project.Description

	return
}
