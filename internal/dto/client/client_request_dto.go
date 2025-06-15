package client

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type RequestDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Nit     string `json:"nit"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	IdRep   int    `json:"id_rep"`
	Arl     string `json:"arl"`
}

func (r *RequestDTO) Deserialize(dto RequestDTO) internal.Client {
	client := internal.Client{
		ID:      dto.Id,
		Name:    dto.Name,
		Nit:     dto.Nit,
		Address: dto.Address,
		Phone:   dto.Phone,
		Email:   dto.Email,
		IdRep:   dto.IdRep,
		Arl:     dto.Arl,
	}

	return client
}
