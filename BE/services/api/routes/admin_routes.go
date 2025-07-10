package routes

import (
	"api/handlers"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(r *mux.Router) {
	admin := r.PathPrefix("/").Subrouter()
	admin.Use(handlers.JWTAuthMiddleware)
	admin.Use(handlers.AdminOnly)

	admin.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	admin.HandleFunc("/users/list", handlers.GetUserList).Methods("POST")
	admin.HandleFunc("/cards/assign", handlers.AssignCardHandler).Methods("POST")
	admin.HandleFunc("/balances/topup", handlers.TopUpBalanceHandler).Methods("POST")
	admin.HandleFunc("/stats/users", handlers.GetUserStatsHandler).Methods("GET")
	admin.HandleFunc("/stats/cards", handlers.GetCardStatsHandler).Methods("GET")
	admin.HandleFunc("/stats/total-balance", handlers.GetTotalBalanceHandler).Methods("GET")
	admin.HandleFunc("/cards/list", handlers.GetCardsList).Methods("POST")
	admin.HandleFunc("/charges/list", handlers.GetCharges).Methods("POST")
	admin.HandleFunc("/balances/ride-cost", handlers.ChangeRideCost).Methods("POST")
	admin.HandleFunc("/balances/list", handlers.BalanceList).Methods("POST")
	admin.HandleFunc("/devices", handlers.CreateDevices).Methods("POST")
	admin.HandleFunc("/devices/list", handlers.DevicesList).Methods("POST")
	admin.HandleFunc("/cards/activate", handlers.AddCardActivationHandler).Methods("POST")
	// RESTful user endpoints
	admin.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	admin.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	admin.HandleFunc("/users/{id}", handlers.SoftDeleteUserHandler).Methods("DELETE")
	// RESTful card endpoints
	admin.HandleFunc("/cards/{id}", handlers.GetCardByIDHandler).Methods("GET")
	admin.HandleFunc("/cards/{id}", handlers.UpdateCardHandler).Methods("PUT")
	admin.HandleFunc("/cards/{id}", handlers.SoftDeleteCardHandler).Methods("DELETE")
	// RESTful device endpoints
	admin.HandleFunc("/devices/{id}", handlers.GetDeviceByIDHandler).Methods("GET")
	admin.HandleFunc("/devices/{id}", handlers.UpdateDeviceHandler).Methods("PUT")
	admin.HandleFunc("/devices/{id}", handlers.SoftDeleteDeviceHandler).Methods("DELETE")
}
