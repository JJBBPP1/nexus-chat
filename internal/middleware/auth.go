package middleware

import (
	"net/http"
	"strings"

	"nexus-chat/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token no válido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		_, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Token no válido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
