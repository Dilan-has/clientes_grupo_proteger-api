package legal_rep

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type RequestDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cc   string `json:"cc"`
}

func (r *RequestDTO) Deserialize(request RequestDTO) internal.LegalRep {
	legalRep := internal.LegalRep{
		ID:   request.ID,
		Name: request.Name,
		Cc:   request.Cc,
	}

	return legalRep
}
