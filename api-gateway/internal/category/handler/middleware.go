package handler

import (
	"context"
	"lib"
	"net/http"
	"strings"
	
	auth "user-service/cmd/auth"
)

type RoleEnum string

const (
	Member    RoleEnum = "member"
	Librarian RoleEnum = "librarian"
	Admin     RoleEnum = "admin"
)

type StatusEnum string

const (
	Pending  StatusEnum = "pending"
	Verified StatusEnum = "verified"
	Rejected StatusEnum = "rejected"
)

type ContextKey string

func AuthAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sp := lib.CreateRootSpan(r, "GetAllUsers")
		defer sp.Finish()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]

		lib.LogRequest(sp, token)

		// Verify token with user-service
		claims, err := auth.ValidateAccessToken(token)

		if err != nil {
			lib.WriteResponse(sp, w, lib.NewErrUnauthorized(err.Error()), nil)
			return
		}

		// Set claims in context
		ctx := context.WithValue(r.Context(), ContextKey("claims"), claims)

		// Validate claims role
		// if string(claims.Role) != string(Admin) {
		// 	// Token is valid and user is an admin, call the next handler
		// 	lib.WriteResponse(sp, w, lib.NewErrForbidden("Invalid user role access"), nil)
		// 	return
		// }

		switch claims.Role {
		case string(Admin):
			// Token is valid and user is an admin, call the next handler
			// lib.WriteResponse(sp, w, lib.NewErrForbidden("Invalid user role access for admin"), nil)

			// Call the next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		case string(Librarian):
			// Token is valid and user is a librarian, call the next handler

			// Call the next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		case string(Member):
			// Token is valid and user is a member, call the next handler

			// Call the next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		default:
			// Token is valid but user role is not recognized, return forbidden error

			// Call the next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

	}
}
