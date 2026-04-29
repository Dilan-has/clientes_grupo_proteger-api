package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/organization"
	"github.com/go-chi/chi/v5"
	resp "github.com/nicklaw5/go-respond"
)

func NewOrganizationHandler(svc internal.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{svc: svc}
}

type OrganizationHandler struct {
	svc internal.OrganizationService
}

func (h *OrganizationHandler) FindAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		orgs, err := h.svc.FindAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var dtos []organization.ResponseDTO
		for _, o := range orgs {
			dto := organization.ResponseDTO{}
			dtos = append(dtos, dto.Serialize(o))
		}

		resp.NewResponse(writer).Ok(dtos)
	}
}

func (h *OrganizationHandler) FindByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		org, err := h.svc.FindByID(id)
		if err != nil {
			if err == internal.ErrOrganizationNotFound {
				writer.WriteHeader(http.StatusNotFound)
				return
			}
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := organization.ResponseDTO{}
		resp.NewResponse(writer).Ok(dto.Serialize(org))
	}
}

func (h *OrganizationHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto organization.RequestDTO
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		org := dto.Deserialize()
		if err := h.svc.Create(&org); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseDto := organization.ResponseDTO{}
		resp.NewResponse(writer).Created(responseDto.Serialize(org))
	}
}

func (h *OrganizationHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto organization.RequestDTO
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		org := dto.Deserialize()
		if err := h.svc.Update(&org); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseDto := organization.ResponseDTO{}
		resp.NewResponse(writer).Ok(responseDto.Serialize(org))
	}
}

func (h *OrganizationHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.svc.Delete(id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(nil)
	}
}
