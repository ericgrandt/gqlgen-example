package middleware

import (
	"context"
	"net/http"
)

var userCtxKey = contextKey{name: "user"}

type contextKey struct {
	name string
}

type User struct {
	ID string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: For example purposes, just hardcode a user with id of 1
		ctx := context.WithValue(r.Context(), userCtxKey, User{ID: "1"})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserContext(ctx context.Context) User {
	user, _ := ctx.Value(userCtxKey).(User)
	return user
}
