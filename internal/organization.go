package internal

import "errors"

type Organization struct {
	Id   int
	Name string
	Link string
}

type OrganizationService interface {
	FindAll() ([]Organization, error)
	FindByID(id int) (Organization, error)
	Create(organization *Organization) error
	Update(organization *Organization) error
	Delete(id int) error
}

type OrganizationRepository interface {
	FindAll() ([]Organization, error)
	FindByID(id int) (Organization, error)
	Create(organization *Organization) error
	Update(organization *Organization) error
	Delete(id int) error
}

var (
	ErrOrganizationNotFound   = errors.New("repository: organization not found")
	ErrOrganizationDuplicated = errors.New("repository: organization already exists")
)
