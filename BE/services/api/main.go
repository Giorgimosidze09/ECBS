package main

import (
	dbClient "client/database"
	"log"
	"net/http"
	"os"

	api_config "api/config"
	handlers "api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Load config
	dir, _ := os.Getwd()
	log.Println("Running from:", dir)

	cfg := api_config.LoadFromEnv()
	api_config.Set(cfg)

	dbClient.SetupClient(cfg.NatsURL)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // or ["*"] for all
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Admin
	r.Post("/users", handlers.CreateUserHandler)
	r.Post("/users/list", handlers.GetUserList)
	r.Post("/cards/assign", handlers.AssignCardHandler)
	r.Post("/balances/topup", handlers.TopUpBalanceHandler)
	r.Get("/stats/users", handlers.GetUserStatsHandler)
	r.Get("/stats/cards", handlers.GetCardStatsHandler)
	r.Get("/stats/total-balance", handlers.GetTotalBalanceHandler)
	r.Post("/cards/list", handlers.GetCardsList)
	r.Post("/charges/list", handlers.GetCharges)
	r.Post("/balances/ride-cost", handlers.ChangeRideCost)
	r.Post("/balances/list", handlers.BalanceList)
	r.Post("/devices", handlers.CreateDevices)
	r.Post("/devices/list", handlers.DevicesList)

	// webhook validation
	r.Post("/cards/validate", handlers.ValidateCardHandler)
	r.Post("/webhook/card-scan", handlers.HandleCardScanWebhook)
	r.Get("/devices/{device_id}/authorized-access", handlers.SyncAuthorizedAccessHandler)
	r.Post("/access-logs/sync", handlers.SyncAccessLogs)

	r.Post("/cards/activate", handlers.AddCardActivationHandler)

	// RESTful user endpoints
	r.Get("/users/{id}", handlers.GetUserByIDHandler)
	r.Put("/users/{id}", handlers.UpdateUserHandler)
	r.Delete("/users/{id}", handlers.SoftDeleteUserHandler)

	// RESTful card endpoints
	r.Get("/cards/{id}", handlers.GetCardByIDHandler)
	r.Put("/cards/{id}", handlers.UpdateCardHandler)
	r.Delete("/cards/{id}", handlers.SoftDeleteCardHandler)

	// RESTful device endpoints
	r.Get("/devices/{id}", handlers.GetDeviceByIDHandler)
	r.Put("/devices/{id}", handlers.UpdateDeviceHandler)
	r.Delete("/devices/{id}", handlers.SoftDeleteDeviceHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 API server running on :" + port)
	http.ListenAndServe(":"+port, r)
}
