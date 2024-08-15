package repository

import (
	"context"
	"errors"
	"log"
	"time"
	user_proto "user-service/userpb"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, req *user_proto.CreateUserRequest) (*user_proto.User, error) {
	var idResult int64

	query := `INSERT INTO users (username, email, password, role, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	now := time.Now()

	err := r.db.QueryRow(ctx, query, req.Username, req.Email, req.Password, req.Role, req.Status).Scan(&idResult)

	if err != nil {
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &user_proto.User{
		Id:        int64(idResult),
		Username:  req.Username,
		Email:     req.Email,
		Role:      req.Role,
		Status:    req.Status,
		CreatedAt: timestampProto,
		UpdatedAt: timestampProto,
	}, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error) {
	var results user_proto.GetAllUsersResponse

	offset := (payload.Page - 1) * payload.Limit
	query := `
	    SELECT id, username, email, role, created_at, updated_at
	    FROM users
	    ORDER BY id
	    LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, payload.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var total_count int

	queryCount := `SELECT COUNT(*) FROM users`
	err = r.db.QueryRow(ctx, queryCount).Scan(&total_count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result user_proto.GetUserResponse
		err := rows.Scan(&result.Id, &result.Username, &result.Email, &result.Role, &result.Status)
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

func (r *UserRepository) GetDetail(ctx context.Context, id int64) (*user_proto.GetUserResponse, error) {
	var result user_proto.GetUserResponse
	query := `SELECT id, username, email, role, status FROM users WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&result.Id, &result.Username, &result.Email, &result.Role, &result.Status)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, req *user_proto.UpdateUserRequest) (*user_proto.User, error) {

	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4, status = $5 WHERE id = $6`
	now := time.Now()

	result, err := r.db.Exec(ctx, query, req.Username, req.Email, req.Password, req.Role, req.Status, req.Id)

	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		errors.New("no rows affected during update operation")
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &user_proto.User{
		Id:        req.Id,
		Username:  req.Username,
		Email:     req.Email,
		Role:      req.Role,
		Status:    req.Status,
		CreatedAt: timestampProto,
		UpdatedAt: timestampProto,
	}, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error) {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(ctx, query, id)

	if err != nil {
		log.Println(err)
		return &user_proto.DeleteUserResponse{Success: false}, err
	}

	rowsNum := res.RowsAffected()

	if rowsNum == 0 {
		errors.New("no rows affected during delete operation")
		return &user_proto.DeleteUserResponse{Success: false}, err
	}

	return &user_proto.DeleteUserResponse{Success: true}, nil
}

func (r *UserRepository) GetDetailByEmail(ctx context.Context, email string) (*user_proto.User, error) {
	var result user_proto.User

	query := `SELECT id, username, email, password, role, status FROM users WHERE email = $1`

	err := r.db.QueryRow(ctx, query, email).Scan(&result.Id, &result.Username, &result.Email, &result.Password, &result.Role, &result.Status)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *UserRepository) Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.User, error) {
	var result user_proto.User

	query := `SELECT id, username, email, password, role, status FROM users WHERE email = $1 And password = $2`

	err := r.db.QueryRow(ctx, query, payload.Email, payload.Password).Scan(&result.Id, &result.Username, &result.Email, &result.Password, &result.Role, &result.Status)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
