package repository

import (
	"context"
	"errors"
	"log"
	"time"
	
	category_proto "category-service/categorypb"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, req *category_proto.CreateCategoryRequest) (*category_proto.Category, error) {
	var idResult int64

	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`

	now := time.Now()

	err := r.db.QueryRow(ctx, query, req.Name).Scan(&idResult)

	if err != nil {
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &category_proto.Category{
		Id:        int64(idResult),
		Name:      req.Name,
		CreatedAt: timestampProto,
		UpdatedAt: timestampProto,
	}, nil
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context, payload *category_proto.GetAllCategoriesRequest) (*category_proto.GetAllCategoriesResponse, error) {
	var results category_proto.GetAllCategoriesResponse

	offset := (payload.Page - 1) * payload.Limit
	query := `
	    SELECT id, name
	    FROM categories
	    ORDER BY id
	    LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, payload.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var total_count int

	queryCount := `SELECT COUNT(*) FROM categories`
	err = r.db.QueryRow(ctx, queryCount).Scan(&total_count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result category_proto.GetCategoryResponse
		err := rows.Scan(&result.Id, &result.Name)
		if err != nil {
			return nil, err
		}
		results.Items = append(results.Items, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *CategoryRepository) GetCategoryById(ctx context.Context, payload *category_proto.GetDetailCategoryRequest) (*category_proto.GetCategoryResponse, error) {
	var result category_proto.GetCategoryResponse
	query := `SELECT id, name FROM categories WHERE id = $1`
	err := r.db.QueryRow(ctx, query, payload.Id).Scan(&result.Id, &result.Name)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, req *category_proto.UpdateCategoryRequest) (*category_proto.Category, error) {

	query := `UPDATE categories SET name = $1 WHERE id = $2`
	now := time.Now()

	result, err := r.db.Exec(ctx, query, req.Name, req.Id)

	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		errors.New("no rows affected during update operation")
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &category_proto.Category{
		Id:        req.Id,
		Name:      req.Name,
		CreatedAt: timestampProto,
		UpdatedAt: timestampProto,
	}, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id *category_proto.DeleteCategoryRequest) (*category_proto.DeleteCategoryResponse, error) {
	query := `DELETE FROM categories WHERE id = $1`
	res, err := r.db.Exec(ctx, query, id)

	if err != nil {
		log.Println(err)
		return &category_proto.DeleteCategoryResponse{Success: false}, err
	}

	rowsNum := res.RowsAffected()

	if rowsNum == 0 {
		errors.New("no rows affected during delete operation")
		return &category_proto.DeleteCategoryResponse{Success: false}, err
	}

	return &category_proto.DeleteCategoryResponse{Success: true}, nil
}
