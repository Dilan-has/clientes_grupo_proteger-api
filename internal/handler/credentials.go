package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/credentials"
	"github.com/go-chi/chi/v5"
	resp "github.com/nicklaw5/go-respond"
)

func NewCredentialsHandler(credentialsService internal.CredentialsService) *CredentialsHandler {
	return &CredentialsHandler{credentialsService: credentialsService}
}

type CredentialsHandler struct {
	credentialsService internal.CredentialsService
}

func (h *CredentialsHandler) FindAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		credential, err := h.credentialsService.FindAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var credentialsDTO []credentials.ResponseDTO
		for _, c := range credential {
			dto := credentials.ResponseDTO{}
			credentialDTO := dto.Serialize(c)
			credentialsDTO = append(credentialsDTO, credentialDTO)
		}

		resp.NewResponse(writer).Ok(credential)

	}

}

func (h *CredentialsHandler) FindByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		credential, err := h.credentialsService.FindByID(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := credentials.ResponseDTO{}
		credentialDTO := dto.Serialize(credential)

		resp.NewResponse(writer).Ok(credentialDTO)
	}
}

func (h *CredentialsHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto credentials.RequestDTO
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		credential := dto.Deserialize()
		if err := h.credentialsService.Create(&credential); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(credential)
	}
}

func (h *CredentialsHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.credentialsService.Delete(id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(nil)
	}
}

func (h *CredentialsHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto credentials.RequestDTO
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		credential := dto.Deserialize()
		if err := h.credentialsService.Update(&credential); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(credential)
	}
}

func (h *CredentialsHandler) FindByClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "idClient")
		id, err := strconv.Atoi(idStr)
		credential, err := h.credentialsService.FindByClient(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var credentialsDTO []credentials.ResponseDTO
		for _, c := range credential {
			dto := credentials.ResponseDTO{}
			credentialDTO := dto.Serialize(c)
			credentialsDTO = append(credentialsDTO, credentialDTO)
		}

		resp.NewResponse(writer).Ok(credential)
	}
}
