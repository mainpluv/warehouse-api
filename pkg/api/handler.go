package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"warehouse-api/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.ServiceInterface
}

// создаем новый экземпляр обработчика апи
func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{service: service}
}

// запрос на резервацию
func (h *Handler) ReserveProducts(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Codes []string `json:"codes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.ReserveProducts(r.Context(), request.Codes); err != nil {
		http.Error(w, "Failed to reserve products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// запрос на освобождение из резервации
func (h *Handler) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Codes []string `json:"codes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.ReleaseProducts(r.Context(), request.Codes); err != nil {
		http.Error(w, "Failed to release products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// запрос на получение оставшихся товаров
func (h *Handler) GetWarehouseStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	stock, err := h.service.GetWarehouseStock(r.Context(), warehouseID)
	if err != nil {
		http.Error(w, "Failed to get warehouse stock", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// запрос на получение зарезервированных товаров
func (h *Handler) GetReservedStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	stock, err := h.service.GetReservedStock(r.Context(), warehouseID)
	if err != nil {
		http.Error(w, "Failed to get reserved stock", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}
