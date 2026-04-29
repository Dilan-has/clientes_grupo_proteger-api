package organization

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type ResponseDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (r *ResponseDTO) Serialize(org internal.Organization) ResponseDTO {
	return ResponseDTO{
		Id:   org.Id,
		Name: org.Name,
		Link: org.Link,
	}
}
