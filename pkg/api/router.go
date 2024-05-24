package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// маршрутизатор для апи
type Router struct {
	handler *Handler
}

// новый экземплр маршрутизатора
func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

// запуск http сервера
func (r *Router) Start(addr string) error {
	router := mux.NewRouter()
	router.HandleFunc("/reserve", r.handler.ReserveProducts).Methods("POST")
	router.HandleFunc("/release", r.handler.ReleaseProducts).Methods("POST")
	router.HandleFunc("/warehouse/{id}/stock", r.handler.GetWarehouseStock).Methods("GET")
	router.HandleFunc("/warehouse/{id}/reserved", r.handler.GetReservedStock).Methods("GET")

	return http.ListenAndServe(addr, router)
}
