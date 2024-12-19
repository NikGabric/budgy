package handlers

import (
	"backend/internal/helpers"
	"backend/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	queries *repository.Queries
}

func NewUserHandler(queries *repository.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

type LoginUserResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userDto LoginUserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(userDto)

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

	token, err := helpers.CreateJwt(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := LoginUserResponse{Token: token}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResp))
}
