package main

import (
	"log"
	"net/http"

	"go_final_project/internal/config"
	"go_final_project/pkg/router"
)

func main() {
	config.MustLoad()

	port := ":" + config.Port

	router := router.SetupRouter()

	log.Printf("Server is running at %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Fatalf("server stopped")
}
