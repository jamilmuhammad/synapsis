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
	Username    string `json:"username" validate:"required,min=6"`
	Email  string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role   string `json:"role" validate:"required,oneof=member librarian admin"`
	Status   string `json:"status" validate:"required,oneof=pending verified rejected"`
}

type GetAllUsersRequest struct {
	Page string `json:"page" default:"1"`
	Limit string `json:"limit" default:"10"`
}

type GetDetailUserRequest struct {
	ID       string `json:"id" validate:"required"`
}

type UpdateUserRequest struct {
	ID       string `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required,min=6"`
	Email  string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role   string `json:"role" validate:"required,oneof=member librarian admin"`
	Status   string `json:"status" validate:"required,oneof=pending verified rejected"`
}

type DeleteUserRequest struct {
	ID       string `json:"id" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken    string `json:"refresh_token" validate:"required"`
}

type UserResponse struct {
	Username    string `json:"username"`
	Email  string `json:"email"`
	Password string `json:"-"`
	Role   string `json:"role"`
	Status   string `json:"status"`
}
