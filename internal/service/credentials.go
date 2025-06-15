package service

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

func NewCredentialsDefault(rp internal.CredentialsRepository) *CredentialsDefault {
	return &CredentialsDefault{
		rp: rp,
	}
}

type CredentialsDefault struct {
	rp internal.CredentialsRepository
}

func (s *CredentialsDefault) FindAll() (credentials []internal.Credentials, err error) {
	credentials, err = s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	return
}

func (s *CredentialsDefault) FindByID(id int) (credential internal.Credentials, err error) {
	credential, err = s.rp.FindByID(id)
	if err != nil {
		return internal.Credentials{}, err
	}
	return
}

func (s *CredentialsDefault) Create(credential *internal.Credentials) (err error) {
	err = s.rp.Create(credential)
	if err != nil {
		return err
	}
	return
}

func (s *CredentialsDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return
}

func (s *CredentialsDefault) Update(credential *internal.Credentials) (err error) {
	err = s.rp.Update(credential)
	if err != nil {
		return err
	}
	return
}

func (s *CredentialsDefault) FindByClient(idClient int) (credentials []internal.Credentials, err error) {
	credentials, err = s.rp.FindByClient(idClient)
	if err != nil {
		return nil, err
	}
	return
}
