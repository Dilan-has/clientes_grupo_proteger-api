package client

import (
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
)

type ResponseDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Nit     string `json:"nit"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	IdRep   int    `json:"id_rep"`
	Arl     string `json:"arl"`
}

func (r *ResponseDTO) Serialize(client internal.Client) ResponseDTO {
	clientResponse := ResponseDTO{
		Id:      client.ID,
		Name:    client.Name,
		Nit:     client.Nit,
		Address: client.Address,
		Phone:   client.Phone,
		Email:   client.Email,
		IdRep:   client.IdRep,
		Arl:     client.Arl,
	}

	return clientResponse
}
