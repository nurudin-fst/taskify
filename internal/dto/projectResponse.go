package dto

type ProjectInsert struct {
	Id   int    `json:"project_id"`
	Name string `json:"name"`
}

type ProjectUpdate struct {
	Id          int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
