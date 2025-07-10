package routes

import (
	"api/handlers"

	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(r *mux.Router) {
	customer := r.PathPrefix("/customer").Subrouter()
	customer.Use(handlers.JWTAuthMiddleware)
	customer.Use(handlers.CustomerOnly)
	customer.HandleFunc("/sum-balance", handlers.CustomerSumBalanceHandler).Methods("POST")
}
