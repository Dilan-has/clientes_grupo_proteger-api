package repository

import (
	"database/sql"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindByUsername(username string, password string) (bool, error) {
	var auth internal.Auth
	err := repo.db.QueryRow("SELECT u.`user`, u.`password` FROM user AS `u`  WHERE u.`user` = ? AND u.`password` = ?", username, password).
		Scan(&auth.Username, &auth.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}
