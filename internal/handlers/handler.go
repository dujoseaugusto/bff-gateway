package handlers

import (
	"bff-gatway/internal/domain"
	"bff-gatway/internal/helpers"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Port   domain.PortService
	Logger *helpers.Logger
}

func NewHandler(repository domain.PortService, logger *helpers.Logger) *Handler {
	return &Handler{
		Port:   repository,
		Logger: logger,
	}
}

func SendResponse(w http.ResponseWriter, response interface{}, statusCode int) error {
	w.Header().Set("Content-security-policy", "default-src 'none")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)

}

func SendResponseText(w http.ResponseWriter, response string, statusCode int) {
	w.Header().Set("Content-security-policy", "default-src 'none")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}
