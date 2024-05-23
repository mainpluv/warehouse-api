package api

import (
    "github.com/gorilla/mux"
    "net/http"
)

type Router struct {
    handler *Handler
}

func NewRouter(handler *Handler) *Router {
    return &Router{handler: handler}
}

func (r *Router) Start(addr string) error {
    router := mux.NewRouter()
    router.HandleFunc("/reserve", r.handler.ReserveProducts).Methods("POST")
    router.HandleFunc("/release", r.handler.ReleaseProducts).Methods("POST")
    router.HandleFunc("/warehouse/{id}/stock", r.handler.GetWarehouseStock).Methods("GET")

    return http.ListenAndServe(addr, router)
}
