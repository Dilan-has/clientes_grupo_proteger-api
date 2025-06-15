package service

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

func NewAffiliateDefault(rp internal.AffiliateRepository) *AffiliateDefault {
	return &AffiliateDefault{
		rp: rp,
	}
}

type AffiliateDefault struct {
	rp internal.AffiliateRepository
}

func (s *AffiliateDefault) FindAll() (affiliates []internal.Affiliate, err error) {
	affiliates, err = s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	return
}

func (s *AffiliateDefault) FindByID(id int) (affiliate internal.Affiliate, err error) {
	affiliate, err = s.rp.FindByID(id)
	if err != nil {
		return internal.Affiliate{}, err
	}
	return
}

func (s *AffiliateDefault) FindByCc(cc string) (affiliates []internal.Affiliate, err error) {
	affiliates, err = s.rp.FindByCc(cc)
	if err != nil {
		return []internal.Affiliate{}, err
	}
	return
}

func (s *AffiliateDefault) Create(affiliate *internal.Affiliate) (err error) {
	err = s.rp.Create(affiliate)
	if err != nil {
		return err
	}
	return
}

func (s *AffiliateDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return
}

func (s *AffiliateDefault) Update(affiliate *internal.Affiliate) (err error) {
	err = s.rp.Update(affiliate)
	if err != nil {
		return err
	}
	return
}

func (s *AffiliateDefault) FindByClientID(id int) (affiliates []internal.Affiliate, err error) {
	affiliates, err = s.rp.FindByClientID(id)
	if err != nil {
		return nil, err
	}
	return
}
