package main

import (
	"context"
	"fmt"
	"log"

	"nexus-chat/internal/auth"
	"nexus-chat/internal/database"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	repository := database.NewUserRepository(conn)

	password := "12345678"

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	err = repository.CreateUser(
		"test_user",
		"test@nexus.local",
		passwordHash,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuario creado correctamente")
	fmt.Println("Hash almacenado:")
	fmt.Println(passwordHash)
}
