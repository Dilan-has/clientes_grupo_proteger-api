package service

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

func NewLegalRepDefault(rp internal.LegalRepRepository) *LegalRepDefault {
	return &LegalRepDefault{
		rp: rp,
	}
}

type LegalRepDefault struct {
	rp internal.LegalRepRepository
}

func (s *LegalRepDefault) FindAll() (legalReps []internal.LegalRep, err error) {
	legalReps, err = s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	return
}

func (s *LegalRepDefault) FindByID(id int) (legalRep internal.LegalRep, err error) {
	legalRep, err = s.rp.FindByID(id)
	if err != nil {
		return internal.LegalRep{}, err
	}
	return
}

func (s *LegalRepDefault) Create(legalRep *internal.LegalRep) (err error) {
	err = s.rp.Create(legalRep)
	if err != nil {
		return err
	}
	return
}

func (s *LegalRepDefault) Update(legalRep *internal.LegalRep) (err error) {
	err = s.rp.Update(legalRep)
	if err != nil {
		return err
	}
	return
}

func (s *LegalRepDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return
}

func (s *LegalRepDefault) FindByCc(cc string) (legalRep internal.LegalRep, err error) {
	legalRep, err = s.rp.FindByCc(cc)
	if err != nil {
		return internal.LegalRep{}, err
	}
	return
}
