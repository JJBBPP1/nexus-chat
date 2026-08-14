package handlers

import (
	"encoding/json"
	"net/http"

	"nexus-chat/internal/auth"
	"nexus-chat/internal/database"
	"nexus-chat/internal/models"
)

type LoginHandler struct {
	userRepository *database.UserRepository
}

func NewLoginHandler(userRepository *database.UserRepository) http.HandlerFunc {
	handler := &LoginHandler{
		userRepository: userRepository,
	}

	return handler.Handle
}

func (h *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON no válido", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		http.Error(w, "El correo electrónico es obligatorio", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		http.Error(w, "La contraseña es obligatoria", http.StatusBadRequest)
		return
	}

	userID, username, passwordHash, err := h.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	if err := auth.CheckPassword(req.Password, passwordHash); err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(userID, username)
	if err != nil {
		http.Error(w, "Error generando el token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Inicio de sesión correcto",
		"token":   token,
		"user": map[string]string{
			"username": username,
			"email":    req.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
