package internal

import "errors"

type Credentials struct {
	Id           int
	IdClient     int
	Organization string
	User         string
	Password     string
}

type CredentialsService interface {
	FindAll() ([]Credentials, error)
	FindByID(id int) (Credentials, error)
	Create(credentials *Credentials) error
	Update(credentials *Credentials) error
	Delete(id int) error
	FindByClient(idClient int) ([]Credentials, error)
}

type CredentialsRepository interface {
	FindAll() ([]Credentials, error)
	FindByID(id int) (Credentials, error)
	Create(credentials *Credentials) error
	Update(credentials *Credentials) error
	Delete(id int) error
	FindByClient(idClient int) ([]Credentials, error)
}

type CredentialsField string

var (
	ErrCredentialsRepositoryNotFound   = errors.New("repository: credentials not found")
	ErrCredentialsRepositoryDuplicated = errors.New("repository: credentials already exists")
)
