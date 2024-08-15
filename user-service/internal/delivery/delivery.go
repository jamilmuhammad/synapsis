package delivery

import (
	"context"
	"log"
	"time"
	auth "user-service/cmd/auth"
	"user-service/internal/interfaces"
	"user-service/userpb"
	user_proto "user-service/userpb"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// type UserUseCase interface {
// 	GetAllUsers(ctx context.Context, limit int, offset int) (*user_proto.GetAllUsersResponse, error)
// 	CreateUser(ctx context.Context, post *user_proto.CreateUserRequest) (*user_proto.User, error)
// 	UpdateUser(ctx context.Context, post *user_proto.UpdateUserRequest) (*user_proto.User, error)
// 	GetDetail(ctx context.Context, id int64) (*user_proto.GetUserResponse, error)
// 	GetDetailByEmail(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.User, error)
// 	DeleteUser(ctx context.Context, id int64) error
// }

type UserDelivery struct {
	userpb.UnimplementedUserServiceServer
	repoUser interfaces.UserInterfaceRepository
}

func NewDelivery(repoUser interfaces.UserInterfaceRepository) *UserDelivery {
	return &UserDelivery{repoUser: repoUser}
}

func (uc *UserDelivery) GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error) {
	result, err := uc.repoUser.GetAllUsers(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) CreateUser(ctx context.Context, payload *user_proto.CreateUserRequest) (*user_proto.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	payload.Password = string(hashedPassword)

	result, err := uc.repoUser.CreateUser(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) GetDetail(ctx context.Context, id int64) (*user_proto.GetUserResponse, error) {
	result, err := uc.repoUser.GetDetail(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) GetDetailByEmail(ctx context.Context, email string) (*user_proto.User, error) {
	result, err := uc.repoUser.GetDetailByEmail(ctx, email)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) UpdateUser(ctx context.Context, user *user_proto.UpdateUserRequest) (*user_proto.User, error) {
	result, err := uc.repoUser.UpdateUser(ctx, user)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) DeleteUser(ctx context.Context, id *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error) {
	result, err := uc.repoUser.DeleteUser(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserDelivery) Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	var result *user_proto.LoginResponse

	user, err := uc.repoUser.GetDetailByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return nil, err
	}

	userData := &user_proto.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
	}

	result.Data = userData

	userClaims := auth.UserClaims{
		Id:       string(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
		MapClaims: jwt.MapClaims{
			"user_id": string(user.Id),
			"iat":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	access_token, err := auth.GenerateAccessToken(userClaims)

	if err != nil {
		return nil, err
	}

	result.AccessToken = access_token

	refresh_token, err := auth.GenerateRefreshToken(userClaims)

	if err != nil {
		return nil, err
	}

	result.RefreshToken = refresh_token

	return result, nil
}
