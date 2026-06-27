package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/v1/health" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "authorization header is missing", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		userId, err := getUserIDFromToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		r.Header.Set("X-USER-ID", userId)

		next.ServeHTTP(w, r)
	})
}

func getUserIDFromToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}

	if _, _, err := jwt.NewParser().ParseUnverified(tokenString, claims); err != nil {
		return "", err
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", errors.New("Sub claim is missing")
	}

	return sub, nil
}
