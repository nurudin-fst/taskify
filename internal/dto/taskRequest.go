package dto

type TaskInsertIn struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"  validate:"required"`
	Deadline    string `json:"deadline" validate:"required"`
}

type TaskUpdateIn struct {
	Title       string `json:"title"  validate:"required"`
	Status      string `json:"status" validate:"required,oneof=todo in_progress done" enums:"todo,in_progress,done"`
	Description string `json:"description" validate:"required"`
	Deadline    string `json:"deadline" validate:"required"`
}
