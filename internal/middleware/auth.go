package middleware

import (
	"net/http"
	"strings"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	resp "github.com/nicklaw5/go-respond"
)

func AuthMiddleware(authService internal.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				resp.NewResponse(w).Unauthorized("Missing Authorization header")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				resp.NewResponse(w).Unauthorized("Invalid Authorization header format")
				return
			}

			valid, err := authService.ValidateToken(token)
			if err != nil || !valid {
				resp.NewResponse(w).Unauthorized("Invalid or expired token")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
