package middleware

import (
	"context"
	"e-com/src/helper"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			helper.HandleHeader(w)
			w.WriteHeader(http.StatusOK)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.HandleHeader(w)
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			helper.HandleHeader(w)
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := helper.VerifyJWT(token)
		if err != nil {
			helper.HandleHeader(w)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
} 