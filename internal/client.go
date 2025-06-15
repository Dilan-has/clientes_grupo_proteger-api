package internal

import "errors"

type Client struct {
	ID      int
	Name    string
	Nit     string
	Address string
	Phone   string
	Email   string
	IdRep   int
	Arl     string
}

type ClientService interface {
	FindAll() ([]Client, error)
	FindByID(id int) (Client, error)
	FindByNit(nit string) (Client, error)
	FindByLegalRepID(id int) (Client, error)
	Create(client *Client) error
	Update(client *Client) error
	Delete(id int) error
}

type ClientRepository interface {
	FindAll() ([]Client, error)
	FindByID(id int) (Client, error)
	FindByNit(nit string) (Client, error)
	FindByLegalRepID(id int) (Client, error)
	Create(client *Client) error
	Update(client *Client) error
	Delete(id int) error
}

type ClientField string

var (
	ErrClientRepositoryNotFound   = errors.New("repository: client not found")
	ErrClientRepositoryDuplicated = errors.New("repository: client already exists")
)
