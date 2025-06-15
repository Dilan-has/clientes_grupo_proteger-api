package credentials

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type RequestDTO struct {
	Id           int    `json:"id"`
	IdClient     int    `json:"id_client"`
	Organization string `json:"organization"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

func (r *RequestDTO) Deserialize() (credential internal.Credentials) {
	credential = internal.Credentials{
		Id:           r.Id,
		IdClient:     r.IdClient,
		Organization: r.Organization,
		User:         r.User,
		Password:     r.Password,
	}

	return
}
