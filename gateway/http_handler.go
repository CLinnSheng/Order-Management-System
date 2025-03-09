package main

import (
	"net/http"
)

type handler struct {
	// gateway

}

func NewHandler() *handler {
	return &handler{}
}

// a function that create routing for all the api endpoint registered
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandlerCreateOrder)
}

// handler to handle http request of create order with api endpoint /api/customers/{customerID}/orders
func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
}
