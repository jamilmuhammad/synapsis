package auth

import (
	"context"
	"fmt"
	"lib"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var JWT_SIGNATURE_KEY = []byte("secret")

type UserClaims struct {
	Id    string `json:"id"`
	Username  string `json:"username"`
	Email string `json:"email"`
	Role string `json:"role"`
	Status string `json:"status"`
	jwt.MapClaims
}

func GenerateAccessToken(claims UserClaims) (string, error) {
	cfg := lib.LoadConfigByFile("../cmd", "config", "yaml")

	// get config access token string
	accessToken := []byte(cfg.Token.AccessToken)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	tokenString, err := token.SignedString(accessToken)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(claims UserClaims) (string, error) {
	cfg := lib.LoadConfigByFile("../cmd", "config", "yaml")

	// get config refresh token string
	refreshToken := []byte(cfg.Token.RefreshToken)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	tokenString, err := token.SignedString(refreshToken)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(accessToken string) (*UserClaims, error) {
	cfg := lib.LoadConfigByFile("../cmd", "config", "yaml")

	// get config access token string
	accessConfigToken := []byte(cfg.Token.AccessToken)

	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return accessConfigToken, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unexpected signing method: %v", err)
	}

}

func ValidateRefreshToken(refreshToken string) (*UserClaims, error) {
	cfg := lib.LoadConfigByFile("../cmd", "config", "yaml")

	// get config refresh token string
	refreshConfigToken := []byte(cfg.Token.RefreshToken)

	token, err := jwt.ParseWithClaims(refreshToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return refreshConfigToken, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unexpected signing method: %v", err)
	}

}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	claims, err := ValidateAccessToken(token[0])
	if err != nil {
		return nil, err
	}

	type contextKey string

	newCtx := context.WithValue(ctx, contextKey("user_id"), claims.MapClaims["user_id"])
	return newCtx, nil
}

func UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		newCtx, err := AuthInterceptor(ctx)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
