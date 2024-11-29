package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/repository"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	config := config.GetConfig()

	conn, err := pgx.Connect(ctx, config.DbConnString)
	if err != nil {
		panic("! Unable to connect to DB !")
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)

	router := http.NewServeMux()

	usersHandler := handlers.NewUsersHandler(queries)

	router.HandleFunc("/", handlers.GetStatus)
	router.HandleFunc("/user/{id}", usersHandler.GetUser)
	router.HandleFunc("/user", usersHandler.CreateUser)

	server := http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic("! Unable to start server !")
	}
}
