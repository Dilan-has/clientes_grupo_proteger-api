package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/user"
	resp "github.com/nicklaw5/go-respond"
)

// Constructor para el AuthHandler
func NewAuthHandler(authService internal.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type AuthHandler struct {
	authService internal.AuthService
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var creds user.User
		err := json.NewDecoder(request.Body).Decode(&creds)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		valid, err := h.authService.FindByUsername(creds.Username, creds.Password)
		if err != nil || !valid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := h.authService.GenerateToken(creds.Username)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(map[string]string{"token": token})
	}
}
