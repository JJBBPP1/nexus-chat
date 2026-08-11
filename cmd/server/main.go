package main

import (
	"context"
	"fmt"
	"net/http"

	"nexus-chat/internal/database"
	"nexus-chat/internal/handlers"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error conectando a PostgreSQL:", err)
		return
	}
	defer conn.Close(context.Background())

	userRepository := database.NewUserRepository(conn)

	http.HandleFunc("/api/health", handlers.HealthHandler)

	registerHandler := handlers.NewRegisterHandler(userRepository)
	http.HandleFunc("/api/register", registerHandler)

	fmt.Println("Servidor iniciado en http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
