package models

import (
	"time"
)

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type GetAllCategoriesRequest struct {
	Page  string `json:"page" default:"1"`
	Limit string `json:"limit" default:"10"`
}

type GetDetailCategoryRequest struct {
	ID string `json:"id" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id" validate:"required"`
}

type CategoryResponse struct {
	Name string `json:"name"`
}
