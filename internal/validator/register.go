package validator

import (
	"errors"
	"net/mail"
	"regexp"
	"unicode/utf8"
)

func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("el nombre de usuario es obligatorio")
	}

	if utf8.RuneCountInString(username) < 3 {
		return errors.New("el nombre de usuario debe tener al menos 3 caracteres")
	}

	if utf8.RuneCountInString(username) > 20 {
		return errors.New("el nombre de usuario no puede superar los 20 caracteres")
	}

	match, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)

	if !match {
		return errors.New("el nombre de usuario solo puede contener letras, números y _")
	}

	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("el correo electrónico es obligatorio")
	}

	_, err := mail.ParseAddress(email)

	if err != nil {
		return errors.New("el correo electrónico no es válido")
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("la contraseña es obligatoria")
	}

	if utf8.RuneCountInString(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}

	if utf8.RuneCountInString(password) > 64 {
		return errors.New("la contraseña no puede superar los 64 caracteres")
	}

	return nil
}
