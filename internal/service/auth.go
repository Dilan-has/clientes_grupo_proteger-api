package service

import (
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/auth"
)

func NewAuthDefault(rp internal.AuthRepository) *AuthDefault {
	return &AuthDefault{
		rp: rp,
	}
}

type AuthDefault struct {
	rp internal.AuthRepository
}

func (s *AuthDefault) FindByUsername(username string, password string) (bool, error) {
	auth, err := s.rp.FindByUsername(username, password)
	if err != nil {
		return false, err
	}
	return auth, nil
}

func (s *AuthDefault) GenerateToken(username string) (string, error) {
	token, err := auth.NewAuthService().GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthDefault) ValidateToken(token string) (bool, error) {
	valid, err := auth.NewAuthService().ValidateToken(token)
	if err != nil {
		return false, err
	}
	return valid, nil
}
