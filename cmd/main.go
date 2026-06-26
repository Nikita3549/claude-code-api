package main

import (
	"fmt"
	"log"
	"net/http"
)

func App() http.Handler {
	router := http.NewServeMux()

	return router
}

func main() {
	app := App()
	port := 8080

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: app,
	}

	log.Printf("Server started on port %v", port)
	server.ListenAndServe()
}
