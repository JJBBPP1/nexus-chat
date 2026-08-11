package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "12345678"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("error generando hash: %v", err)
	}

	if hash == password {
		t.Fatal("la contraseña no debe almacenarse en texto plano")
	}

	if err := CheckPassword(password, hash); err != nil {
		t.Fatalf("la contraseña correcta no fue validada: %v", err)
	}

	if err := CheckPassword("contraseña_incorrecta", hash); err == nil {
		t.Fatal("una contraseña incorrecta fue aceptada")
	}
}
