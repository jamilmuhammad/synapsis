package models

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type Author struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Bio         string     `json:"bio"`
	DateOfBirth CustomTime `json:"date_of_birth"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateAuthorRequest struct {
	Name        string     `json:"name" validate:"required,min=3"`
	Bio         string     `json:"bio" validate:"max=200"`
	DateOfBirth CustomTime `json:"date_of_birth"`
}

type GetAllAuthorsRequest struct {
	Page  string `json:"page" default:"1"`
	Limit string `json:"limit" default:"10"`
}

type GetDetailAuthorRequest struct {
	ID string `json:"id" validate:"required"`
}

type UpdateAuthorRequest struct {
	ID          string     `json:"id" validate:"required"`
	Name        string     `json:"name" validate:"required,min=3"`
	Bio         string     `json:"bio" validate:"max=200"`
	DateOfBirth CustomTime `json:"date_of_birth"`
}

type DeleteAuthorRequest struct {
	ID string `json:"id" validate:"required"`
}

type AuthorResponse struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	DateOfBirth string `json:"date_of_birth"`
}
