package handler

import (
	"context"
	"log"

	"user-service/internal/interfaces"
	user_proto "user-service/userpb"

)

type UserHandler struct {
	user_proto.UnimplementedUserServiceServer
	repoUser interfaces.UserInterfaceHandler
}

func NewHandler(repoUser interfaces.UserInterfaceHandler) *UserHandler {
	return &UserHandler{repoUser: repoUser}
}

func (uc *UserHandler) GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error) {
	result, err := uc.repoUser.GetAllUsers(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserHandler) CreateUser(ctx context.Context, payload *user_proto.CreateUserRequest) (*user_proto.User, error) {

	result, err := uc.repoUser.CreateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) GetUserById(ctx context.Context, payload *user_proto.GetDetailUserRequest) (*user_proto.GetUserResponse, error) {

	result, err := uc.repoUser.GetUserById(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) GetUserByEmail(ctx context.Context, payload *user_proto.GetDetailUserByEmailRequest) (*user_proto.User, error) {

	result, err := uc.repoUser.GetUserByEmail(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) UpdateUser(ctx context.Context, payload *user_proto.UpdateUserRequest) (*user_proto.User, error) {

	result, err := uc.repoUser.UpdateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) DeleteUser(ctx context.Context, id *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error) {

	result, err := uc.repoUser.DeleteUser(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {

	result, err := uc.repoUser.Login(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *UserHandler) RefreshToken(ctx context.Context, payload *user_proto.RefreshTokenRequest) (*user_proto.RefreshTokenResponse, error) {

	result, err := uc.repoUser.RefreshToken(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}
