package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"claude-code-api/internal/claude"
	"claude-code-api/internal/configs"
	"claude-code-api/internal/requests"
	"claude-code-api/pkg/db"
)

func App() http.Handler {
	router := http.NewServeMux()

	return router
}

func main() {
	app := App()
	conf := configs.LoadConfig()
	port := conf.Port
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

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: app,
	}
	res, _ := claudeService.Ask(context.Background(), claude.AskClaudeRequest{
		Question: "Привет! придумай стихотворение короткое",
		Prompt:   "В каждой строке стихотворения должно быть слово 'солнце'",
	})

	fmt.Println(res)

	log.Printf("Server started on port %v", port)
	server.ListenAndServe()
}
