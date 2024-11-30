package config

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"net/http"
)

func CreateRouter(queries *repository.Queries) *http.ServeMux {
	usersHandler := handlers.NewUsersHandler(queries)
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.GetStatus)
	router.HandleFunc("GET /user/{id}", usersHandler.GetUser)
	router.HandleFunc("POST /user", usersHandler.CreateUser)

	return router
}
