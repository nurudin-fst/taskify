package repository

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	Db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		Db: db,
	}
}

type Task struct {
	Id          int       `gorm:"column:id"`
	ProjectId   int       `gorm:"column:project_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Status      string    `gorm:"column:status"`
	Deadline    time.Time `gorm:"column:deadline"`
}

func (*Task) TableName() string {
	return "tasks"
}

func (r *TaskRepository) Insert(task Task) (Task, error) {
	tx := r.Db.Create(&task)
	err := tx.Error

	return task, err
}

type TaskList struct {
	Id          int    `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Status      string `gorm:"column:status"`
	Deadline    string `gorm:"column:deadline"`
}

func (r *TaskRepository) SelectByProject(projectId int) ([]TaskList, error) {
	var tasks []TaskList
	err := r.Db.Select(`id, title, description, status, deadline`).
		Table("tasks").
		Where("project_id = ?", projectId).
		Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepository) Update(task Task) (Task, error) {
	err := r.Db.Omit("project_id").Save(&task).Error
	return task, err
}

func (r *TaskRepository) Delete(taskId int) (Task, error) {
	var task Task
	err := r.Db.
		Clauses(clause.Returning{}).
		Where("id = ?", taskId).
		Delete(&task).Error

	return task, err
}
