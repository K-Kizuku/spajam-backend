package middleware

import (
	"context"
	"net/http"
)

type Key string

const UserIDKey Key = "userID"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// authHeader := r.Header.Get("Authorization")
		// token := strings.TrimPrefix(authHeader, "Bearer ")
		email := r.Header.Get("X-User-Email")

		// id, err := jwt.VerifyToken(token)
		// if err != nil {
		// 	http.Error(w, "Invalid or token", http.StatusUnauthorized)
		// 	slog.Error("Invalid or expired ID token", "error", err.Error())
		// 	return
		// }

		ctx = context.WithValue(ctx, UserIDKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
