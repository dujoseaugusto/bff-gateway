package handlers

import (
	"bff-gatway/internal/domain"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// @Summary Login User
// @Description Login
// @Tags user-login
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /auth/login [post]
func (h *Handler) HandleGetUserLogin(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.WithContext(r.Context())

	var requestData json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		logger.Printf("Error: Request body: %v", string(requestData))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	logger.Printf("Start request HandleGetUserLogin")

	userRequest := domain.User{
		Data: requestData,
	}

	user, err := h.Port.GetUserLogin(&userRequest)

	if err != nil {
		logger.Printf("Error request HandleGetUser: %v", err.Err)
		SendResponseText(w, err.Body, err.CodStatus)
		return
	}

	if *user != "" {
		SendResponseText(w, *user, http.StatusOK)
		return
	}

	_ = SendResponse(w, nil, http.StatusInternalServerError)

}

// @Summary Retorna todos os usuários
// @Description Get user
// @Tags user-content
// @Accept json
// @Produce json
// @Success 200 {object} json.RawMessage
// @Router /auth [get]
func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.WithContext(r.Context())

	logger.Printf("Start request HandleGetUser")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		_ = SendResponse(w, `{"error": "No Authorization header provided"}`, http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		_ = SendResponse(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "Token is empty"}`, http.StatusUnauthorized)
		return
	}

	user, err := h.Port.GetUser(&token)

	if err != nil {
		logger.Printf("Error request HandleGetUser: %v", err.Err)
		SendResponseText(w, err.Body, err.CodStatus)
		return
	}

	if *user != nil {
		_ = SendResponse(w, *user, http.StatusOK)
		return
	}

	_ = SendResponse(w, nil, http.StatusInternalServerError)

}

// @Summary Delete User
// @Description Delete user
// @Tags user-content
// @Accept json
// @Produce json
// @Success 200 {object} json.RawMessage
// @Router /auth/user/{token} [delete]
func (h *Handler) HandleGetUserRemove(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.WithContext(r.Context())

	logger.Printf("Start request HandleGetUserRemove")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		_ = SendResponse(w, `{"error": "No Authorization header provided"}`, http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		_ = SendResponse(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "Token is empty"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	tokenUser := vars["token"]
	if tokenUser == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "User is empty"}`, http.StatusBadRequest)
		return
	}

	user, err := h.Port.DeleteUser(&domain.InputUser{
		Token:     token,
		TokenUser: tokenUser,
	})

	if err != nil {
		logger.Printf("Error request HandleGetUser: %v", err.Err)
		SendResponseText(w, err.Body, err.CodStatus)
		return
	}

	if *user != nil {
		_ = SendResponse(w, *user, http.StatusOK)
		return
	}

	_ = SendResponse(w, nil, http.StatusInternalServerError)

}

// @Summary Delete User
// @Description get user
// @Tags user-content
// @Accept json
// @Produce json
// @Success 200 {object} json.RawMessage
// @Router /auth/user/{token} [get]
func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.WithContext(r.Context())

	logger.Printf("Start request HandleGetUserRemove")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		_ = SendResponse(w, `{"error": "No Authorization header provided"}`, http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		_ = SendResponse(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "Token is empty"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	tokenUser := vars["token"]
	if tokenUser == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "User is empty"}`, http.StatusBadRequest)
		return
	}

	user, err := h.Port.GetUserOne(&domain.InputUser{
		Token:     token,
		TokenUser: tokenUser,
	})

	if err != nil {
		logger.Printf("Error request HandleGetUser: %v", err.Err)
		SendResponseText(w, err.Body, err.CodStatus)
		return
	}

	if *user != nil {
		_ = SendResponse(w, *user, http.StatusOK)
		return
	}

	_ = SendResponse(w, nil, http.StatusInternalServerError)

}

// @Summary Delete User
// @Description get user
// @Tags user-content
// @Accept json
// @Produce json
// @Success 200 {object} json.RawMessage
// @Router /auth/user/{token} [get]
func (h *Handler) HandleGetUserRegister(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.WithContext(r.Context())

	logger.Printf("Start request HandleGetUserRemove")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		_ = SendResponse(w, `{"error": "No Authorization header provided"}`, http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		_ = SendResponse(w, `{"error": "Invalid Authorization header format"}`, http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		logger.Printf("Token is empty")
		_ = SendResponse(w, `{"error": "Token is empty"}`, http.StatusUnauthorized)
		return
	}

	var payload json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Printf("Body invalid")
		_ = SendResponse(w, `{"error": "Body invalid"}`, http.StatusBadRequest)
		return
	}

	user, err := h.Port.SetUser(&domain.InputUser{
		Token: token,
		Body:  payload,
	})

	if err != nil {
		logger.Printf("Error request HandleGetUser: %v", err.Err)
		SendResponseText(w, err.Body, err.CodStatus)
		return
	}

	if *user != nil {
		_ = SendResponse(w, *user, http.StatusOK)
		return
	}

	_ = SendResponse(w, nil, http.StatusInternalServerError)

}
