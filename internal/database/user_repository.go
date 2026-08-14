package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) CreateUser(
	username string,
	email string,
	passwordHash string,
) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
	`

	_, err := r.conn.Exec(
		context.Background(),
		query,
		username,
		email,
		passwordHash,
	)

	if err != nil {
		return fmt.Errorf("error creando usuario: %w", err)
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (int64, string, string, error) {
	query := `
		SELECT id, username, password_hash
		FROM users
		WHERE email = $1
	`

	var userID int64
	var username string
	var passwordHash string

	err := r.conn.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(&userID, &username, &passwordHash)

	if err != nil {
		return 0, "", "", fmt.Errorf("error buscando usuario: %w", err)
	}

	return userID, username, passwordHash, nil
}
