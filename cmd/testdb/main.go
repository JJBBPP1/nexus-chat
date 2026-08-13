package main

import (
	"fmt"

	"nexus-chat/internal/database"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Error conectando a PostgreSQL:", err)
		return
	}

	userRepository := database.NewUserRepository(conn)

	username, passwordHash, err := userRepository.GetUserByEmail("pepe@gmail.com")
	if err != nil {
		fmt.Println("Error buscando usuario:", err)
		return
	}

	fmt.Println("Usuario encontrado:", username)
	fmt.Println("Hash almacenado:", passwordHash)
}
