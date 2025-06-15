package internal

type Auth struct {
	Username string
	Password string
}

type AuthService interface {
	FindByUsername(username string, password string) (bool, error)
	GenerateToken(username string) (string, error)

	ValidateToken(token string) (bool, error)
}

type AuthRepository interface {
	FindByUsername(username string, password string) (bool, error)
}
