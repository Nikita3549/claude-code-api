package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"claude-code-api/internal/claude"
	"claude-code-api/internal/configs"
	"claude-code-api/internal/requests"
	"claude-code-api/pkg/db"
)

func App() (http.Handler, *db.DB) {
	router := http.NewServeMux()
	conf := configs.LoadConfig()
	db := db.NewDB(conf)

	// Repositories
	requestRepository := requests.NewRequestRepository(db)

	// Services
	claudeClient := claude.NewClient(conf)
	claudeService := claude.NewClaudeService(
		&claude.ClaudeServiceDeps{
			RequestRepository: requestRepository,
			ClaudeClient:      claudeClient,
		},
	)

	// Handlers
	claude.NewClaudeHandler(router, claudeService)

	return router, db
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, db := App()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app,
	}

	log.Printf("Server started on port %v", port)
	go server.ListenAndServe()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx)
	db.Close()
	log.Println("Server stopped")
}
