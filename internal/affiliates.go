package internal

import "errors"

type Affiliate struct {
	ID        int
	Name      string
	Cc        string
	Eps       string
	Status    string
	IdClient  int
	Pension   string
	Risk      string
	Caja      string
	Income    string
	Birthdate string
}

type AffiliateService interface {
	FindAll() ([]Affiliate, error)
	FindByID(id int) (Affiliate, error)
	FindByCc(cc string) ([]Affiliate, error)
	FindByClientID(id int) ([]Affiliate, error)
	Create(affiliate *Affiliate) error
	Delete(id int) error
	Update(affiliate *Affiliate) error
}

type AffiliateRepository interface {
	FindAll() ([]Affiliate, error)
	FindByID(id int) (Affiliate, error)
	FindByCc(cc string) ([]Affiliate, error)
	FindByClientID(id int) ([]Affiliate, error)
	Create(affiliate *Affiliate) error
	Delete(id int) error
	Update(affiliate *Affiliate) error
}

type AffiliateField string

const (
	AffiliateID   AffiliateField = "id"
	AffiliateName AffiliateField = "name"
	AffiliateCc   AffiliateField = "cc"
	AffiliateEps  AffiliateField = "eps"
)

var (
	ErrAffiliateRepositoryNotFound   = errors.New("repository: affiliate not found")
	ErrAffiliateRepositoryDuplicated = errors.New("repository: affiliate already exists")
	ErrAffiliateServiceNotFound      = errors.New("service: affiliate not found")
)
