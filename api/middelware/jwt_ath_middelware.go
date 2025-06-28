package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/omaraliali1010/go_template/internal/jwtservice"
)

func JwtAuthMiddleware(service *jwtservice.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) != 2 || strings.ToLower(t[0]) != "bearer" {
				http.Error(w, jsonError("Authorization header format must be Bearer {token}"), http.StatusUnauthorized)
				return
			}

			tokenStr := t[1]
			log.Println("JwtAuthMiddleware tokenStr", tokenStr)
			claims, err := service.ParseAccessToken(tokenStr)
			log.Println("JwtAuthMiddleware claims", claims)
			if err != nil {
				http.Error(w, jsonError("Invalid token: "+err.Error()), http.StatusUnauthorized)
				return
			}

			// Inject user ID into context
			ctx := context.WithValue(r.Context(), "x-user-id", claims.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func jsonError(message string) string {
	return `{"message": "` + message + `"}`
}
