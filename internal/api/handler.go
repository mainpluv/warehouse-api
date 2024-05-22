package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"warehouse-api/internal/models"

	"github.com/gorilla/mux"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) ReserveItems(w http.ResponseWriter, r *http.Request) {
	var req models.ReserveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ReserveItems(h.DB, req.WarehouseID, req.Codes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ReleaseItems(w http.ResponseWriter, r *http.Request) {
	var req models.ReleaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ReleaseItems(h.DB, req.WarehouseID, req.Codes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetRemainingItems(w http.ResponseWriter, r *http.Request) {
	var req models.RemainingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items, err := GetRemainingItems(h.DB, req.WarehouseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func SetupRouter(db *sql.DB) *mux.Router {
	handler := &Handler{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/reserve", handler.ReserveItems).Methods("POST")
	router.HandleFunc("/release", handler.ReleaseItems).Methods("POST")
	router.HandleFunc("/remaining", handler.GetRemainingItems).Methods("POST")
	return router
}
