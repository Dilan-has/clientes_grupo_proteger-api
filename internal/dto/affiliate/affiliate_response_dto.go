package affiliate

import "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"

type ResponseDTO struct {
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

func (r *ResponseDTO) Serialize(affiliate internal.Affiliate) ResponseDTO {
	affiliateResponse := ResponseDTO{
		Id:              affiliate.ID,
		Name:            affiliate.Name,
		Cc:              affiliate.Cc,
		Eps:             affiliate.Eps,
		Status:          affiliate.Status,
		IdClient:        affiliate.IdClient,
		Pension:         affiliate.Pension,
		Risk:            affiliate.Risk,
		Caja:            affiliate.Caja,
		Income:          affiliate.Income,
		Birthdate:       affiliate.Birthdate,
		LastPaymentDate: affiliate.LastPaymentDate,
	}

	return affiliateResponse
}
