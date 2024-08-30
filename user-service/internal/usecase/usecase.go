package usecase

import (
	"context"
	"fmt"
	"lib"
	"log"
	"time"
	auth "user-service/cmd/auth"
	"user-service/internal/interfaces"

	// "user-service/internal/repository"
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

type UserUseCase struct {
	// repo repository.UserRepository
	repoUser interfaces.UserInterfaceUseCase
}

func NewUseCase(repoUser interfaces.UserInterfaceUseCase) *UserUseCase {
	return &UserUseCase{
		repoUser,
	}
}

func (uc *UserUseCase) GetAllUsers(ctx context.Context, payload *user_proto.GetAllUsersRequest) (*user_proto.GetAllUsersResponse, error) {
	result, err := uc.repoUser.GetAllUsers(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserUseCase) CreateUser(ctx context.Context, payload *user_proto.CreateUserRequest) (*user_proto.User, error) {
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

func (uc *UserUseCase) GetUserById(ctx context.Context, payload *user_proto.GetDetailUserRequest) (*user_proto.GetUserResponse, error) {

	result, err := uc.repoUser.GetUserById(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *UserUseCase) GetUserByEmail(ctx context.Context, payload *user_proto.GetDetailUserByEmailRequest) (*user_proto.User, error) {
	result, err := uc.repoUser.GetUserByEmail(ctx, payload)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return &user_proto.User{}, lib.NewErrNotFound(fmt.Sprintf("user email %s not found", payload.Email))
	}

	if result.Status != string(lib.Verified) {
		return &user_proto.User{}, lib.NewErrForbidden("user is not verified")
	}

	return result, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, payload *user_proto.UpdateUserRequest) (*user_proto.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)

	result, err := uc.repoUser.UpdateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id *user_proto.DeleteUserRequest) (*user_proto.DeleteUserResponse, error) {
	result, err := uc.repoUser.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *UserUseCase) Login(ctx context.Context, payload *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	var result = &user_proto.LoginResponse{}

	user, err := uc.repoUser.GetUserByEmail(ctx, &user_proto.GetDetailUserByEmailRequest{Email: payload.Email})

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

	result.User = userData

	Id := fmt.Sprintf("%c", user.Id)

	userClaims := auth.UserClaims{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
		MapClaims: jwt.MapClaims{
			"user_id": Id,
			"iat":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	access_token, err := auth.GenerateAccessToken(userClaims)

	if err != nil {
		return nil, err
	}

	result.AccessToken = access_token

	userClaims.MapClaims = jwt.MapClaims{
		"user_id": Id,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refresh_token, err := auth.GenerateRefreshToken(userClaims)

	if err != nil {
		return nil, err
	}

	result.RefreshToken = refresh_token

	return result, nil
}

func (uc *UserUseCase) RefreshToken(ctx context.Context, payload *user_proto.RefreshTokenRequest) (*user_proto.RefreshTokenResponse, error) {
	var result = &user_proto.RefreshTokenResponse{}

	claims, err := auth.ValidateRefreshToken(payload.RefreshToken)

	if err != nil {
		return nil, err
	}

	claims.MapClaims = jwt.MapClaims{
		"user_id": claims.Id,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	access_token, err := auth.GenerateAccessToken(*claims)

	if err != nil {
		return nil, err
	}

	result.AccessToken = access_token

	result.RefreshToken = payload.RefreshToken

	return result, nil
}
