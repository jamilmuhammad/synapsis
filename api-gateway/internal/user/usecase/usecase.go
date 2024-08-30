package usecase

import (
	"context"

	user_proto "user-service/userpb"
)

type userUseCase struct {
	userClient user_proto.UserServiceClient
}

func NewUserUseCase(userClient user_proto.UserServiceClient) *userUseCase {
	return &userUseCase{userClient}
}

func (u *userUseCase) GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error) {

	result, err := u.userClient.GetAllUsers(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, payload *user_proto.CreateUserRequest) (*user_proto.User, error) {

	result, err := u.userClient.CreateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) GetUserById(ctx context.Context, payload *user_proto.GetDetailUserRequest) (*user_proto.GetUserResponse, error) {

	result, err := uc.userClient.GetUserById(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) GetDetailByEmail(ctx context.Context, payload *user_proto.GetDetailUserByEmailRequest) (*user_proto.User, error) {

	result, err := uc.userClient.GetUserByEmail(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, payload *user_proto.UpdateUserRequest) (*user_proto.User, error) {

	result, err := u.userClient.UpdateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, payload *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error) {

	result, err := u.userClient.DeleteUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	var result = &user_proto.LoginResponse{}

	result, err := uc.userClient.Login(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) RefreshToken(ctx context.Context, payload *user_proto.RefreshTokenRequest) (*user_proto.RefreshTokenResponse, error) {
	var result = &user_proto.RefreshTokenResponse{}

	result, err := uc.userClient.RefreshToken(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}
