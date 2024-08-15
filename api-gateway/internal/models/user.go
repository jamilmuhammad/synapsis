package models

import (
	"time"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role string `json:"role"`
	Status string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserRequest struct {
	Username    string `json:"username" validate:"required,min=20"`
	Email  string `json:"email" validate:"required,min=20"`
	Password string `json:"password" validate:"required,min=6"`
	Role   string `json:"role" validate:"required,oneof=member librarian admin"`
	Status   string `json:"status" validate:"required,oneof=pending verified rejected"`
}

type GetDetailUserRequest struct {
	ID       string `json:"id" validate:"required"`
}

type UpdateUserRequest struct {
	ID       string `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required,min=20"`
	Email  string `json:"email" validate:"required,min=20"`
	Password string `json:"password" validate:"required,min=6"`
	Role   string `json:"role" validate:"required,oneof=member librarian admin"`
	Status   string `json:"status" validate:"required,oneof=pending verified rejected"`
}

type UserResponse struct {
	Username    string `json:"username"`
	Email  string `json:"email"`
	Password string `json:"-"`
	Role   string `json:"role"`
	Status   string `json:"status"`
}
