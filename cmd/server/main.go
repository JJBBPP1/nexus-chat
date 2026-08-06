package main

import (
	"fmt"
	"net/http"

	"nexus-chat/internal/handlers"
)

func main() {

	http.HandleFunc("/api/health", handlers.HealthHandler)

	fmt.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
