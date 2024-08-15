package usecase

import (
	"context"
	"log"
	"api_gateway/internal/models"
	user_proto "user-service/userpb"
	"strconv"
)

type userUseCase struct {
	userClient user_proto.UserServiceClient
}

func NewUserUseCase(userClient user_proto.UserServiceClient) *userUseCase {
	return &userUseCase{userClient}
}

func (u *userUseCase) GetUsers(ctx context.Context, limit string, offset string) (*user_proto.GetUserRequest, error) {

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := u.userClient.GetUsers(ctx, &user_proto.GetUsersRequest{Limit: int32(limitInt), Offset: int32(offsetInt)})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, payload models.CreateUserRequest) (*user_proto.Post, error) {

	req := &user_proto.CreateUserRequest{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   payload.Status,
	}

	result, err := u.userClient.CreateUser(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, payload models.UpdateUserRequest) (*user_proto.Post, error) {
	id, err := strconv.Atoi(payload.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req := &user_proto.UpdateUserRequest{
		Id:       int32(id),
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   payload.Status,
	}

	result, err := u.userClient.UpdateUserById(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id string) (*user_proto.GetUserResponse, error) {
	result, err := u.userClient.GetUserById(ctx, &user_proto.GetUserByIdRequest{Id: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	_, err := u.userClient.DeleteUserById(ctx, &user_proto.DeleteUserRequest{Id: id})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
