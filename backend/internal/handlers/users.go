package handlers

import (
	"backend/internal/helpers"
	"backend/internal/repository"
	"context"
	"encoding/json"
	"log"
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

	userDto.Password, err = helpers.HashPassword(userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

type LoginUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userDto LoginUserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.queries.GetUserByUsername(context.Background(), userDto.Username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	matchPass := helpers.CheckPasswordHash(user.Password, userDto.Password)
	if !matchPass {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: token
	w.WriteHeader(http.StatusOK)
}
