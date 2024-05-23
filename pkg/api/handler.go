package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"warehouse-api/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ReserveProducts(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Codes []string `json:"codes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ReserveProducts(r.Context(), request.Codes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Codes []string `json:"codes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ReleaseProducts(r.Context(), request.Codes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetWarehouseStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	stock, err := h.service.GetWarehouseStock(r.Context(), warehouseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stock)
}
