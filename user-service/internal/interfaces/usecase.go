package interfaces

import (
	"context"
	user_proto "user-service/userpb"
)

type UserInterfaceUseCase interface {
	GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error)
	CreateUser(ctx context.Context, post *user_proto.CreateUserRequest) (*user_proto.User, error)
	UpdateUser(ctx context.Context, post *user_proto.UpdateUserRequest) (*user_proto.User, error)
	GetUserById(ctx context.Context, payload *user_proto.GetDetailUserRequest) (*user_proto.GetUserResponse, error)
	GetUserByEmail(ctx context.Context, payload *user_proto.GetDetailUserByEmailRequest) (*user_proto.User, error)
	DeleteUser(ctx context.Context, id *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error)
	Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.User, error)

}
