package repository

import (
	"database/sql"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (repo *UserRepository) FindByUsername(username string, password string) (bool, error) {
	var auth internal.Auth
	err := repo.db.QueryRow(
		"SELECT u.`user`, u.`password` FROM user AS u WHERE u.`user` = ? AND u.`password` = ?",
		username, password,
	).Scan(&auth.Username, &auth.Password)

	if err == sql.ErrNoRows {
		repo.logger.Warn("Usuario no encontrado o contraseña incorrecta",
			zap.String("username", username),
		)
		return false, nil
	}

	if err != nil {
		repo.logger.Error("Error al consultar usuario",
			zap.Error(err),
			zap.String("username", username),
		)
		return false, err
	}

	repo.logger.Info("Usuario autenticado correctamente",
		zap.String("username", username),
	)
	return true, nil
}
