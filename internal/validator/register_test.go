package validator

import "testing"

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{
			name:     "usuario válido",
			username: "pepe123",
			wantErr:  false,
		},
		{
			name:     "usuario vacío",
			username: "",
			wantErr:  true,
		},
		{
			name:     "usuario demasiado corto",
			username: "ab",
			wantErr:  true,
		},
		{
			name:     "usuario con caracteres no permitidos",
			username: "pepe@123",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.username)

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ValidateUsername() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "email válido",
			email:   "pepe@gmail.com",
			wantErr: false,
		},
		{
			name:    "email vacío",
			email:   "",
			wantErr: true,
		},
		{
			name:    "email inválido",
			email:   "hola",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ValidateEmail() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "contraseña válida",
			password: "12345678",
			wantErr:  false,
		},
		{
			name:     "contraseña vacía",
			password: "",
			wantErr:  true,
		},
		{
			name:     "contraseña demasiado corta",
			password: "123",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ValidatePassword() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
