package credentials

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type ResponseDTO struct {
	Id             int    `json:"id"`
	IdClient       int    `json:"id_client"`
	OrganizationId int    `json:"organization_id"`
	User           string `json:"user"`
	Password       string `json:"password"`
}

func (r *ResponseDTO) Serialize(credential internal.Credentials) ResponseDTO {
	credentialResponse := ResponseDTO{
		Id:             credential.Id,
		IdClient:       credential.IdClient,
		OrganizationId: credential.OrganizationId,
		User:           credential.User,
		Password:       credential.Password,
	}

	return credentialResponse
}
