package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"claude-code-api/internal/claude"
	"claude-code-api/internal/configs"
	"claude-code-api/internal/requests"
	"claude-code-api/pkg/db"
)

func App() http.Handler {
	router := http.NewServeMux()
	conf := configs.LoadConfig()
	db := db.NewDB(conf)

	// Repositories
	requestRepository := requests.NewRequestRepository(db)

	// Services
	claudeClient := claude.NewClient()
	claudeService := claude.NewClaudeService(
		&claude.ClaudeServiceDeps{
			RequestRepository: requestRepository,
			ClaudeClient:      claudeClient,
		},
	)

	// Handlers
	claude.NewClaudeHandler(router, claudeService)

	return router
}

func main() {
	app := App()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app,
	}

	log.Printf("Server started on port %v", port)
	server.ListenAndServe()
}
