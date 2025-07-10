package routes

import (
	"api/handlers"

	"github.com/gorilla/mux"
)

func RegisterPublicRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", handlers.RegisterAuthUserHandler).Methods("POST")
	r.HandleFunc("/auth/login", handlers.LoginAuthUserHandler).Methods("POST")
	r.HandleFunc("/cards/validate", handlers.ValidateCardHandler).Methods("POST")
	// this 4 is not for frontend
	r.HandleFunc("/webhook/card-scan", handlers.HandleCardScanWebhook).Methods("POST")
	r.HandleFunc("/devices/{device_id}/authorized-access", handlers.SyncAuthorizedAccessHandler).Methods("GET")
	r.HandleFunc("/access-logs/sync", handlers.SyncAccessLogs).Methods("POST")
	r.HandleFunc("/api/paybox/topup", handlers.PayboxTopupHandler).Methods("POST")
}
