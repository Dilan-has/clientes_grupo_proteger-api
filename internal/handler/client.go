package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/client"
	"github.com/go-chi/chi/v5"
	resp "github.com/nicklaw5/go-respond"
)

func NewClientHandler(clientService internal.ClientService) *ClientHandler {
	return &ClientHandler{clientService: clientService}
}

type ClientHandler struct {
	clientService internal.ClientService
}

func (h *ClientHandler) FindAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		clients, err := h.clientService.FindAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var clientsDTO []client.ResponseDTO
		for _, c := range clients {
			dto := client.ResponseDTO{}
			clientDTO := dto.Serialize(c)
			clientsDTO = append(clientsDTO, clientDTO)
		}

		resp.NewResponse(writer).Ok(clients)

	}
}

func (h *ClientHandler) FindByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		clients, err := h.clientService.FindByID(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := client.ResponseDTO{}
		clientDTO := dto.Serialize(clients)

		resp.NewResponse(writer).Ok(clientDTO)
	}
}

func (h *ClientHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto client.RequestDTO
		err := json.NewDecoder(request.Body).Decode(&dto)

		client_ := dto.Deserialize(dto)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.clientService.Create(&client_)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(client_)
	}
}

func (h *ClientHandler) FindByNit() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		nit := chi.URLParam(request, "nit")
		clients, err := h.clientService.FindByNit(nit)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := client.ResponseDTO{}
		clientDTO := dto.Serialize(clients)

		resp.NewResponse(writer).Ok(clientDTO)
	}
}

func (h *ClientHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var dto client.RequestDTO
		err := json.NewDecoder(request.Body).Decode(&dto)

		client_ := dto.Deserialize(dto)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.clientService.Update(&client_)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseDTO := &client.ResponseDTO{}
		response := responseDTO.Serialize(client_)

		resp.NewResponse(writer).Ok(response)
	}
}

func (h *ClientHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.clientService.Delete(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(nil)
	}
}

func (h *ClientHandler) FindByLegalRepID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "idLegalRep")
		id, err := strconv.Atoi(idStr)
		clients, err := h.clientService.FindByLegalRepID(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := client.ResponseDTO{}
		clientDTO := dto.Serialize(clients)

		resp.NewResponse(writer).Ok(clientDTO)
	}
}
