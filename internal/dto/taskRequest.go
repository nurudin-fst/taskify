package dto

type TaskInsertIn struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

type TaskUpdateIn struct {
	Title       string `json:"title"`
	Status      string `json:"status" validate:"oneof=todo in_progress done" enums:"todo,in_progress,done"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}
