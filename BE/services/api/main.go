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
	r.Group(func(admin chi.Router) {
		admin.Use(handlers.JWTAuthMiddleware)
		admin.Use(handlers.AdminOnly)
		admin.Post("/users", handlers.CreateUserHandler)
		admin.Post("/users/list", handlers.GetUserList)
		admin.Post("/cards/assign", handlers.AssignCardHandler)
		admin.Post("/balances/topup", handlers.TopUpBalanceHandler)
		admin.Get("/stats/users", handlers.GetUserStatsHandler)
		admin.Get("/stats/cards", handlers.GetCardStatsHandler)
		admin.Get("/stats/total-balance", handlers.GetTotalBalanceHandler)
		admin.Post("/cards/list", handlers.GetCardsList)
		admin.Post("/charges/list", handlers.GetCharges)
		admin.Post("/balances/ride-cost", handlers.ChangeRideCost)
		admin.Post("/balances/list", handlers.BalanceList)
		admin.Post("/devices", handlers.CreateDevices)
		admin.Post("/devices/list", handlers.DevicesList)
		admin.Post("/cards/activate", handlers.AddCardActivationHandler)
		// RESTful user endpoints
		admin.Get("/users/{id}", handlers.GetUserByIDHandler)
		admin.Put("/users/{id}", handlers.UpdateUserHandler)
		admin.Delete("/users/{id}", handlers.SoftDeleteUserHandler)
		// RESTful card endpoints
		admin.Get("/cards/{id}", handlers.GetCardByIDHandler)
		admin.Put("/cards/{id}", handlers.UpdateCardHandler)
		admin.Delete("/cards/{id}", handlers.SoftDeleteCardHandler)
		// RESTful device endpoints
		admin.Get("/devices/{id}", handlers.GetDeviceByIDHandler)
		admin.Put("/devices/{id}", handlers.UpdateDeviceHandler)
		admin.Delete("/devices/{id}", handlers.SoftDeleteDeviceHandler)
	})

	// Public endpoints
	r.Post("/auth/register", handlers.RegisterAuthUserHandler)
	r.Post("/auth/login", handlers.LoginAuthUserHandler)

	// Webhook and validation endpoints (if you want to restrict, move to admin group)
	r.Post("/cards/validate", handlers.ValidateCardHandler)
	r.Post("/webhook/card-scan", handlers.HandleCardScanWebhook)
	r.Get("/devices/{device_id}/authorized-access", handlers.SyncAuthorizedAccessHandler)
	r.Post("/access-logs/sync", handlers.SyncAccessLogs)

	r.Group(func(customer chi.Router) {
		customer.Use(handlers.JWTAuthMiddleware)
		customer.Use(handlers.CustomerOnly) // You’ll need to add this middleware, similar to AdminOnly
		customer.Post("/customer/sum-balance", handlers.CustomerSumBalanceHandler)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 API server running on :" + port)
	http.ListenAndServe(":"+port, r)
}
