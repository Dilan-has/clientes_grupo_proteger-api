package internal

type LegalRep struct {
	ID   int
	Name string
	Cc   string
}

type LegalRepService interface {
	FindAll() ([]LegalRep, error)
	FindByID(id int) (LegalRep, error)
	Create(legalRep *LegalRep) error
	Update(legalRep *LegalRep) error
	Delete(id int) error
	FindByCc(cc string) (LegalRep, error)
}

type LegalRepRepository interface {
	FindAll() ([]LegalRep, error)
	FindByID(id int) (LegalRep, error)
	Create(legalRep *LegalRep) error
	Update(legalRep *LegalRep) error
	Delete(id int) error
	FindByCc(cc string) (LegalRep, error)
}
