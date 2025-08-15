package affiliate

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type RequestDTO struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Cc              string `json:"cc"`
	Eps             string `json:"eps"`
	Status          string `json:"status"`
	IdClient        int    `json:"id_client"`
	Pension         string `json:"pension"`
	Risk            string `json:"risk"`
	Caja            string `json:"caja"`
	Income          string `json:"income"`
	Birthdate       string `json:"birthdate"`
	LastPaymentDate string `json:"last_payment_date"`
}

func (request *RequestDTO) Deserialize(dto RequestDTO) internal.Affiliate {
	affiliate := internal.Affiliate{
		ID:              dto.Id,
		Name:            dto.Name,
		Cc:              dto.Cc,
		Eps:             dto.Eps,
		Status:          dto.Status,
		IdClient:        dto.IdClient,
		Pension:         dto.Pension,
		Risk:            dto.Risk,
		Caja:            dto.Caja,
		Income:          dto.Income,
		Birthdate:       dto.Birthdate,
		LastPaymentDate: dto.LastPaymentDate,
	}

	return affiliate
}
