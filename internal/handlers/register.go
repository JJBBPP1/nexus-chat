package handlers

import (
	"encoding/json"
	"net/http"

	"nexus-chat/internal/auth"
	"nexus-chat/internal/database"
	"nexus-chat/internal/models"
	"nexus-chat/internal/validator"
)

type RegisterHandler struct {
	userRepository *database.UserRepository
}

func NewRegisterHandler(userRepository *database.UserRepository) http.HandlerFunc {
	handler := &RegisterHandler{
		userRepository: userRepository,
	}

	return handler.Handle
}

func (h *RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON no válido", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateUsername(req.Username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateEmail(req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error procesando la contraseña", http.StatusInternalServerError)
		return
	}

	err = h.userRepository.CreateUser(
		req.Username,
		req.Email,
		passwordHash,
	)

	if err != nil {
		http.Error(w, "Error creando usuario", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Usuario registrado correctamente",
		"user": map[string]string{
			"username": req.Username,
			"email":    req.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
