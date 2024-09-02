package repository

import (
	"context"
	"errors"
	// "fmt"
	"log"
	"time"

	author_proto "author-service/authorpb"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthorRepository struct {
	db *pgxpool.Pool
}

func NewAuthorRepository(db *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, req *author_proto.CreateAuthorRequest) (*author_proto.Author, error) {
	var idResult int64

	query := `INSERT INTO authors (name, bio, date_of_birth) VALUES ($1, $2, $3) RETURNING id`

	now := time.Now()

	err := r.db.QueryRow(ctx, query, req.Name, req.Bio, req.DateOfBirth.AsTime()).Scan(&idResult)

	if err != nil {
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &author_proto.Author{
		Id:          int64(idResult),
		Name:        req.Name,
		Bio:         req.Bio,
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   timestampProto,
		UpdatedAt:   timestampProto,
	}, nil
}

func (r *AuthorRepository) GetAllAuthors(ctx context.Context, payload *author_proto.GetAllAuthorsRequest) (*author_proto.GetAllAuthorsResponse, error) {
	var results author_proto.GetAllAuthorsResponse

	offset := (payload.Page - 1) * payload.Limit
	query := `
	    SELECT id, name, bio, date_of_birth
	    FROM authors
	    ORDER BY id
	    LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, payload.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var total_count int

	queryCount := `SELECT COUNT(*) FROM authors`
	err = r.db.QueryRow(ctx, queryCount).Scan(&total_count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result author_proto.GetAuthorResponse
		var dateOfBirth time.Time
		err := rows.Scan(&result.Id, &result.Name, &result.Bio, &dateOfBirth)
		if err != nil {
			return nil, err
		}

		result.DateOfBirth = dateOfBirth.Format(time.RFC3339)

		results.Items = append(results.Items, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *AuthorRepository) GetAuthorById(ctx context.Context, payload *author_proto.GetDetailAuthorRequest) (*author_proto.GetAuthorResponse, error) {
	var result author_proto.GetAuthorResponse
	var dateOfBirth time.Time
	query := `SELECT id, name, bio, date_of_birth FROM authors WHERE id = $1`
	err := r.db.QueryRow(ctx, query, payload.Id).Scan(&result.Id, &result.Name, &result.Bio, &dateOfBirth)
	if err != nil {
		return nil, err
	}

	result.DateOfBirth = dateOfBirth.Format(time.RFC3339)

	return &result, nil
}

func (r *AuthorRepository) UpdateAuthor(ctx context.Context, req *author_proto.UpdateAuthorRequest) (*author_proto.Author, error) {

	query := `UPDATE authors SET name = $1, bio = $2, date_of_birth = $3 WHERE id = $4`
	now := time.Now()

	result, err := r.db.Exec(ctx, query, req.Name, req.Bio, req.DateOfBirth.AsTime(), req.Id)

	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		errors.New("no rows affected during update operation")
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &author_proto.Author{
		Id:          req.Id,
		Name:        req.Name,
		Bio:         req.Bio,
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   timestampProto,
		UpdatedAt:   timestampProto,
	}, nil
}

func (r *AuthorRepository) DeleteAuthor(ctx context.Context, id *author_proto.DeleteAuthorRequest) (*author_proto.DeleteAuthorResponse, error) {
	query := `DELETE FROM authors WHERE id = $1`
	res, err := r.db.Exec(ctx, query, id)

	if err != nil {
		log.Println(err)
		return &author_proto.DeleteAuthorResponse{Success: false}, err
	}

	rowsNum := res.RowsAffected()

	if rowsNum == 0 {
		errors.New("no rows affected during delete operation")
		return &author_proto.DeleteAuthorResponse{Success: false}, err
	}

	return &author_proto.DeleteAuthorResponse{Success: true}, nil
}
