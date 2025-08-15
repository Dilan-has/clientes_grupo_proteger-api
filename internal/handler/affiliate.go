package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/dto/affiliate"
	"github.com/go-chi/chi/v5"
	resp "github.com/nicklaw5/go-respond"
	"go.uber.org/zap"
)

func NewAffiliateHandler(affiliateService internal.AffiliateService, historyService internal.HistoryService, logger zap.Logger) *AffiliateHandler {
	return &AffiliateHandler{affiliateService: affiliateService, historyService: historyService, logger: logger}
}

type AffiliateHandler struct {
	affiliateService internal.AffiliateService
	historyService   internal.HistoryService
	logger           zap.Logger
}

func (h *AffiliateHandler) FindAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		affiliates, err := h.affiliateService.FindAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var affiliatesDTO []affiliate.ResponseDTO
		for _, a := range affiliates {
			dto := affiliate.ResponseDTO{}
			affiliateDTO := dto.Serialize(a)
			affiliatesDTO = append(affiliatesDTO, affiliateDTO)
		}

		resp.NewResponse(writer).Ok(affiliates)
	}
}

func (h *AffiliateHandler) FindByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		affiliates, err := h.affiliateService.FindByID(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto := affiliate.ResponseDTO{}
		affiliateDTO := dto.Serialize(affiliates)

		resp.NewResponse(writer).Ok(affiliateDTO)
	}
}

func (h *AffiliateHandler) FindByCc() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cc := chi.URLParam(request, "cc")
		affiliates, err := h.affiliateService.FindByCc(cc)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var affiliatesDTO []affiliate.ResponseDTO
		for _, a := range affiliates {
			dto := affiliate.ResponseDTO{}
			affiliateDTO := dto.Serialize(a)
			affiliatesDTO = append(affiliatesDTO, affiliateDTO)
		}

		resp.NewResponse(writer).Ok(affiliates)
	}
}

func (h *AffiliateHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := h.logger

		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Error("Error leyendo el body: ", zap.Error(err))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		logger.Info("Request Body: ", zap.String("body", string(bodyBytes)))

		request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var dto affiliate.RequestDTO
		err = json.NewDecoder(request.Body).Decode(&dto)
		if err != nil {
			logger.Error("Error deserializando el JSON", zap.Error(err))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		affiliate_ := dto.Deserialize(dto)

		err = h.affiliateService.Create(&affiliate_)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		history := []internal.DateHistory{
			{
				Entry_date: affiliate_.Income,
			},
		}

		payload := &internal.History{
			Id:      affiliate_.ID,
			Name:    affiliate_.Name,
			Cc:      affiliate_.Cc,
			History: history,
		}

		err = h.historyService.SaveHistory(request.Context(), payload)

		responseDTO := &affiliate.ResponseDTO{}
		response := responseDTO.Serialize(affiliate_)

		resp.NewResponse(writer).Ok(response)
	}
}

func (h *AffiliateHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.affiliateService.Delete(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.NewResponse(writer).Ok(nil)
	}
}

func (h *AffiliateHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := h.logger

		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Error("Error leyendo el body: ", zap.Error(err))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		logger.Info("Request Body: ", zap.String("body", string(bodyBytes)))

		request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var dto affiliate.RequestDTO
		err = json.NewDecoder(request.Body).Decode(&dto)
		if err != nil {
			logger.Error("Error deserializando el JSON", zap.Error(err))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		affiliate_ := dto.Deserialize(dto)

		err = h.affiliateService.Update(&affiliate_)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseDTO := &affiliate.ResponseDTO{}
		response := responseDTO.Serialize(affiliate_)

		resp.NewResponse(writer).Ok(response)
	}
}

func (h *AffiliateHandler) FindByClientId() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		clientIdStr := chi.URLParam(request, "clientId")
		clientId, err := strconv.Atoi(clientIdStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		affiliates, err := h.affiliateService.FindByClientID(clientId)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var affiliatesDTO []affiliate.ResponseDTO
		for _, a := range affiliates {
			dto := affiliate.ResponseDTO{}
			affiliateDTO := dto.Serialize(a)
			affiliatesDTO = append(affiliatesDTO, affiliateDTO)
		}

		resp.NewResponse(writer).Ok(affiliates)
	}
}
