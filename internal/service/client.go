package service

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

func NewClientDefault(rp internal.ClientRepository) *ClientDefault {
	return &ClientDefault{
		rp: rp,
	}
}

type ClientDefault struct {
	rp internal.ClientRepository
}

func (s *ClientDefault) FindAll() (clients []internal.Client, err error) {
	clients, err = s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	return
}

func (s *ClientDefault) FindByID(id int) (client internal.Client, err error) {
	client, err = s.rp.FindByID(id)
	if err != nil {
		return internal.Client{}, err
	}
	return
}

func (s *ClientDefault) FindByNit(nit string) (client internal.Client, err error) {
	client, err = s.rp.FindByNit(nit)
	if err != nil {
		return internal.Client{}, err
	}
	return
}

func (s *ClientDefault) Create(client *internal.Client) (err error) {
	err = s.rp.Create(client)
	if err != nil {
		return err
	}
	return
}

func (s *ClientDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return
}

func (s *ClientDefault) Update(client *internal.Client) (err error) {
	err = s.rp.Update(client)
	if err != nil {
		return err
	}
	return
}

func (s *ClientDefault) FindByLegalRepID(id int) (clients internal.Client, err error) {
	clients, err = s.rp.FindByLegalRepID(id)
	if err != nil {
		return internal.Client{}, err
	}
	return
}
