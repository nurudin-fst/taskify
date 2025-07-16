package dto

type ProjectInsertIn struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ProjectUpdateIn struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
