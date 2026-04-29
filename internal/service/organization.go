package service

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

func NewOrganizationDefault(rp internal.OrganizationRepository) *OrganizationDefault {
	return &OrganizationDefault{rp: rp}
}

type OrganizationDefault struct {
	rp internal.OrganizationRepository
}

func (s *OrganizationDefault) FindAll() ([]internal.Organization, error) {
	return s.rp.FindAll()
}

func (s *OrganizationDefault) FindByID(id int) (internal.Organization, error) {
	return s.rp.FindByID(id)
}

func (s *OrganizationDefault) Create(organization *internal.Organization) error {
	return s.rp.Create(organization)
}

func (s *OrganizationDefault) Update(organization *internal.Organization) error {
	return s.rp.Update(organization)
}

func (s *OrganizationDefault) Delete(id int) error {
	return s.rp.Delete(id)
}
