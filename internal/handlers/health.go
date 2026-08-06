package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	respuesta := HealthResponse{
		Status:  "ok",
		Service: "Nexus Chat",
		Version: "0.1.0",
	}

	json.NewEncoder(w).Encode(respuesta)
}
