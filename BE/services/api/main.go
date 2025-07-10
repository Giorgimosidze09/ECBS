package main

import (
	"log"
	"net/http"
	"os"

	api_config "api/config"
	apiRoutes "api/routes"
	dbClient "client/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load config
	dir, _ := os.Getwd()
	log.Println("Running from:", dir)

	cfg := api_config.LoadFromEnv()
	api_config.Set(cfg)

	dbClient.SetupClient(cfg.NatsURL)

	r := mux.NewRouter()

	// Register route groups from modular routes
	apiRoutes.RegisterPublicRoutes(r)
	apiRoutes.RegisterAdminRoutes(r)
	apiRoutes.RegisterCustomerRoutes(r)

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // or ["*"] for all
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 API server running on :" + port)
	http.ListenAndServe(":"+port, handler)
}
