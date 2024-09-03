package repository

import (
	"context"
	"errors"
	"fmt"
	"lib"

	"log"
	"time"

	author_proto "author-service/authorpb"
	book_proto "book-service/bookpb/book"
	category_proto "category-service/categorypb"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookRepository struct {
	db             *pgxpool.Pool
	authorClient   author_proto.AuthorServiceClient
	categoryClient category_proto.CategoryServiceClient
}

func NewBookRepository(db *pgxpool.Pool, authorClient author_proto.AuthorServiceClient, categoryClient category_proto.CategoryServiceClient) *BookRepository {
	return &BookRepository{
		db:             db,
		authorClient:   authorClient,
		categoryClient: categoryClient,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, req *book_proto.CreateBookRequest) (*book_proto.Book, error) {
	var idResult int64

	author, err := r.authorClient.GetAuthorById(ctx, &author_proto.GetDetailAuthorRequest{Id: req.AuthorId})

	if err != nil {
		return nil, err
	}

	if author == nil {
		return &book_proto.Book{}, lib.NewErrNotFound(fmt.Sprintf("author id %v not found", req.AuthorId))
	}

	category, err := r.categoryClient.GetCategoryById(ctx, &category_proto.GetDetailCategoryRequest{Id: req.CategoryId})

	if err != nil {
		return nil, err
	}

	if category == nil {
		return &book_proto.Book{}, lib.NewErrNotFound(fmt.Sprintf("category id %v not found", req.CategoryId))
	}

	query := `INSERT INTO books (title, isbn, date_of_publication, copies, author_id, category_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	now := time.Now()

	err = r.db.QueryRow(ctx, query, req.Title, req.Isbn, req.DateOfPublication.AsTime(), req.Copies, req.AuthorId, req.CategoryId).Scan(&idResult)

	if err != nil {
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &book_proto.Book{
		Id:                int64(idResult),
		Title:             req.Title,
		Isbn:              req.Isbn,
		DateOfPublication: req.DateOfPublication,
		AuthorId:          req.AuthorId,
		CategoryId:        req.CategoryId,
		CreatedAt:         timestampProto,
		UpdatedAt:         timestampProto,
	}, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context, payload *book_proto.GetAllBooksRequest) (*book_proto.GetAllBooksResponse, error) {
	var results book_proto.GetAllBooksResponse

	offset := (payload.Page - 1) * payload.Limit
	query := `
	    SELECT id, title, isbn, date_of_publication, copies, author_id, category_id
	    FROM books
	    ORDER BY id
	    LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, payload.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var total_count int

	queryCount := `SELECT COUNT(*) FROM books`
	err = r.db.QueryRow(ctx, queryCount).Scan(&total_count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result book_proto.GetBookResponse
		var dateOfPublication time.Time
		var authorIdResult int64
		var categoryIdResult int64
		err := rows.Scan(&result.Id, &result.Title, &result.Isbn, &dateOfPublication, &result.Copies, &authorIdResult, &categoryIdResult)
		if err != nil {
			return nil, err
		}

		author, err := r.authorClient.GetAuthorById(ctx, &author_proto.GetDetailAuthorRequest{Id: authorIdResult})

		if err != nil {
			return nil, err
		}

		if author != nil {
			result.Author = &book_proto.GetAuthorResponse{
				Id:          author.Id,
				Name:        author.Name,
				Bio:         author.Bio,
				DateOfBirth: author.DateOfBirth,
			}
		} else {
			result.Author = nil
		}

		category, err := r.categoryClient.GetCategoryById(ctx, &category_proto.GetDetailCategoryRequest{Id: categoryIdResult})

		if err != nil {
			return nil, err
		}

		if category != nil {
			result.Category = &book_proto.GetCategoryResponse{
				Id:   category.Id,
				Name: category.Name,
			}
		} else {
			result.Category = nil
		}

		result.DateOfPublication = dateOfPublication.Format(time.RFC3339)

		results.Items = append(results.Items, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *BookRepository) GetBookById(ctx context.Context, payload *book_proto.GetDetailBookRequest) (*book_proto.GetBookResponse, error) {
	var result book_proto.GetBookResponse
	var dateOfPublication time.Time
	var authorIdResult int64
	var categoryIdResult int64
	query := `SELECT id, title, isbn, date_of_publication, copies, author_id, category_id FROM books WHERE id = $1`
	err := r.db.QueryRow(ctx, query, payload.Id).Scan(&result.Id, &result.Title, &result.Isbn, &dateOfPublication, &result.Copies, &authorIdResult, &categoryIdResult)
	if err != nil {
		return nil, err
	}

	author, err := r.authorClient.GetAuthorById(ctx, &author_proto.GetDetailAuthorRequest{Id: authorIdResult})

	if err != nil {
		return nil, err
	}

	if author != nil {
		result.Author = &book_proto.GetAuthorResponse{
			Id:          author.Id,
			Name:        author.Name,
			Bio:         author.Bio,
			DateOfBirth: author.DateOfBirth,
		}
	} else {
		result.Author = nil
	}

	category, err := r.categoryClient.GetCategoryById(ctx, &category_proto.GetDetailCategoryRequest{Id: categoryIdResult})

	if err != nil {
		return nil, err
	}

	if category != nil {
		result.Category = &book_proto.GetCategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		}
	} else {
		result.Category = nil
	}

	result.DateOfPublication = dateOfPublication.Format(time.RFC3339)

	return &result, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, req *book_proto.UpdateBookRequest) (*book_proto.Book, error) {

	query := `UPDATE books SET title = $1, isbn = $2, date_of_publication = $3, copies = $4, author_id = $5, category_id = $6 WHERE id = $7`
	now := time.Now()

	author, err := r.authorClient.GetAuthorById(ctx, &author_proto.GetDetailAuthorRequest{Id: req.AuthorId})

	if err != nil {
		return nil, err
	}

	if author == nil {
		return &book_proto.Book{}, lib.NewErrNotFound(fmt.Sprintf("author id %v not found", req.AuthorId))
	}

	category, err := r.categoryClient.GetCategoryById(ctx, &category_proto.GetDetailCategoryRequest{Id: req.CategoryId})

	if err != nil {
		return nil, err
	}

	if category == nil {
		return &book_proto.Book{}, lib.NewErrNotFound(fmt.Sprintf("category id %v not found", req.CategoryId))
	}

	result, err := r.db.Exec(ctx, query, req.Title, req.Isbn, req.DateOfPublication.AsTime(), req.Copies, req.AuthorId, req.CategoryId, req.Id)

	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		errors.New("no rows affected during update operation")
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &book_proto.Book{
		Id:                req.Id,
		Title:             req.Title,
		Isbn:              req.Isbn,
		DateOfPublication: req.DateOfPublication,
		AuthorId:          req.AuthorId,
		CategoryId:        req.CategoryId,
		CreatedAt:         timestampProto,
		UpdatedAt:         timestampProto,
	}, nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id *book_proto.DeleteBookRequest) (*book_proto.DeleteBookResponse, error) {
	query := `DELETE FROM books WHERE id = $1`
	res, err := r.db.Exec(ctx, query, id)

	if err != nil {
		log.Println(err)
		return &book_proto.DeleteBookResponse{Success: false}, err
	}

	rowsNum := res.RowsAffected()

	if rowsNum == 0 {
		errors.New("no rows affected during delete operation")
		return &book_proto.DeleteBookResponse{Success: false}, err
	}

	return &book_proto.DeleteBookResponse{Success: true}, nil
}
