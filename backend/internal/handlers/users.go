package handlers

import (
	"backend/internal/repository"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type UsersHandler struct {
	queries *repository.Queries
}

func NewUsersHandler(queries *repository.Queries) *UsersHandler {
	return &UsersHandler{queries: queries}
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.queries.GetUser(context.Background(), int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto repository.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.queries.CreateUser(context.Background(), userDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(userJson)
}
