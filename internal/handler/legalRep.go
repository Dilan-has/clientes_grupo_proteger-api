package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/legal_rep"
	"github.com/go-chi/chi/v5"
	resp "github.com/nicklaw5/go-respond"
)

func NewLegalRepDefault(rp internal.LegalRepRepository) *LegalRepDefault {
	return &LegalRepDefault{
		rp: rp,
	}
}

type LegalRepDefault struct {
	rp internal.LegalRepRepository
}

func (s *LegalRepDefault) FindAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		legalReps, err := s.rp.FindAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var legalRepsDTO []legal_rep.ResponseDto
		for _, c := range legalReps {
			dto := legal_rep.ResponseDto{}
			legalRepDTO := dto.Serialize(c)
			legalRepsDTO = append(legalRepsDTO, legalRepDTO)
		}

		resp.NewResponse(writer).Ok(legalReps)
	}

}

func (s *LegalRepDefault) FindByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		legalRep, err := s.rp.FindByID(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := legal_rep.ResponseDto{}
		legalRepDTO := dto.Serialize(legalRep)

		resp.NewResponse(writer).Ok(legalRepDTO)
	}
}

func (s *LegalRepDefault) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto legal_rep.RequestDTO
		err := json.NewDecoder(request.Body).Decode(&dto)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		legalRep := dto.Deserialize(dto)
		err = s.rp.Create(&legalRep)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(legalRep)
	}
}

func (s *LegalRepDefault) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.rp.Delete(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(nil)
	}
}

func (s *LegalRepDefault) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto legal_rep.RequestDTO
		err := json.NewDecoder(request.Body).Decode(&dto)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		legalRep := dto.Deserialize(dto)
		err = s.rp.Update(&legalRep)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(legalRep)
	}
}

func (s *LegalRepDefault) FindByCc() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cc := chi.URLParam(request, "cc")
		legalRep, err := s.rp.FindByCc(cc)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := legal_rep.ResponseDto{}
		legalRepDTO := dto.Serialize(legalRep)

		resp.NewResponse(writer).Ok(legalRepDTO)
	}
}
