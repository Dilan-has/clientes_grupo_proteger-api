package legal_rep

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type ResponseDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cc   string `json:"cc"`
}

func (r *ResponseDto) Serialize(legalRep internal.LegalRep) ResponseDto {
	legalRepResponse := ResponseDto{
		ID:   legalRep.ID,
		Name: legalRep.Name,
		Cc:   legalRep.Cc,
	}

	return legalRepResponse
}
