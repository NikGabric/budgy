package config

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"net/http"
)

func CreateRouter(queries *repository.Queries) *http.ServeMux {
	router := http.NewServeMux()
	userHandler := handlers.NewUserHandler(queries)

	router.HandleFunc("/", handlers.GetStatus)

	// user
	router.HandleFunc("POST /user/register", userHandler.CreateUser)
	router.HandleFunc("POST /user/login", userHandler.LoginUser)
	router.HandleFunc("GET /user/{id}", userHandler.GetUser)
	router.HandleFunc("POST /user", userHandler.CreateUser)

	return router
}
