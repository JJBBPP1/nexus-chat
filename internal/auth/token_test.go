package auth

import (
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	token, err := GenerateToken(1, "pepe")
	if err != nil {
		t.Fatalf("error generando token: %v", err)
	}

	if token == "" {
		t.Fatal("el token generado está vacío")
	}

	validatedToken, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("error validando token: %v", err)
	}

	if !validatedToken.Valid {
		t.Fatal("el token debería ser válido")
	}
}
