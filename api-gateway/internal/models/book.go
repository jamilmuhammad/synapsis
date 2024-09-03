package models

import (
	"time"
)

type Book struct {
	ID                int64      `json:"id"`
	Title             string     `json:"title"`
	Isbn              string     `json:"isbn"`
	DateOfPublication CustomTime `json:"date_of_publication"`
	Copies            int32      `json:"copies"`
	AuthorId          int64      `json:"author_id"`
	CategoryId        int64      `json:"category_id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CreateBookRequest struct {
	Title             string     `json:"title" validate:"required,min=3"`
	Isbn              string     `json:"isbn" validate:"max=20"`
	DateOfPublication CustomTime `json:"date_of_publication"`
	Copies            int32      `json:"copies"`
	AuthorId          int64      `json:"author_id" validate:"required"`
	CategoryId        int64      `json:"category_id" validate:"required"`
}

type GetAllBooksRequest struct {
	Page  string `json:"page" default:"1"`
	Limit string `json:"limit" default:"10"`
}

type GetDetailBookRequest struct {
	ID string `json:"id" validate:"required"`
}

type UpdateBookRequest struct {
	ID                string     `json:"id" validate:"required"`
	Title             string     `json:"title" validate:"required,min=3"`
	Isbn              string     `json:"isbn" validate:"max=20"`
	DateOfPublication CustomTime `json:"date_of_publication"`
	Copies            int32      `json:"copies"`
	AuthorId          int64      `json:"author_id" validate:"required"`
	CategoryId        int64      `json:"category_id" validate:"required"`
}

type DeleteBookRequest struct {
	ID string `json:"id" validate:"required"`
}

type BookResponse struct {
	Title             string     `json:"title"`
	Isbn              string     `json:"isbn"`
	DateOfPublication CustomTime `json:"date_of_publication"`
	Copies            int32      `json:"copies"`
	AuthorId          int32      `json:"author_id"`
	CategoryId        int32      `json:"category_id"`
}
