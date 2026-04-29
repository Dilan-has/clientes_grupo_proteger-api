package organization

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type RequestDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (r *RequestDTO) Deserialize() (org internal.Organization) {
	org = internal.Organization{
		Id:   r.Id,
		Name: r.Name,
		Link: r.Link,
	}
	return
}
