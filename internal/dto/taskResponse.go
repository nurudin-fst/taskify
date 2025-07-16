package dto

type TaskInsert struct {
	Id       int    `json:"task_id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Deadline string `json:"deadline"`
}

type TaskUpdate struct {
	Id          int    `json:"task_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

type TaskDelete struct {
	Id          int    `json:"task_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}
