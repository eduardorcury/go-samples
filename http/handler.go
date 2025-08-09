package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpHandler struct {
	Service UserService
}

type userRequest struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

func (s *HttpHandler) HandleNewUser(w http.ResponseWriter, r *http.Request) {
	var reqBody userRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	user := &User{
		ID:    reqBody.UserID,
		Email: reqBody.Email,
		Name:  reqBody.Name,
	}

	ctx := r.Context()

	t, err := s.Service.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
