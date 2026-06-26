package main

import (
	"fmt"
	"log"
	"net/http"

	"claude-code-api/internal/configs"
)

func App() http.Handler {
	router := http.NewServeMux()

	return router
}

func main() {
	app := App()
	conf := configs.LoadConfig()
	port := conf.Port

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: app,
	}

	log.Printf("Server started on port %v", port)
	server.ListenAndServe()
}
