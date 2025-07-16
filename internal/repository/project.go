package repository

import (
	"gorm.io/gorm"
)

type ProjectRepository struct {
	Db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		Db: db,
	}
}

type Project struct {
	Id          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	CreatedBy   int    `gorm:"column:created_by"`
}

func (*Project) TableName() string {
	return "projects"
}

func (r *ProjectRepository) Insert(project Project) (Project, error) {
	tx := r.Db.Create(&project)
	err := tx.Error

	return project, err
}

type ProjectList struct {
	Id          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (r *ProjectRepository) List(userId int) ([]ProjectList, error) {
	var projects []ProjectList
	err := r.Db.Select(`id, name, description`).
		Table("projects").
		Find(&projects, "created_by = ?", userId).Error

	return projects, err
}

func (r *ProjectRepository) Update(project Project) (Project, error) {
	err := r.Db.Omit("created_by").Save(&project).Error
	return project, err
}
