package main

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/repository"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()

	conn, err := pgx.Connect(ctx, cfg.DbConnString)
	if err != nil {
		panic("! Unable to connect to DB !")
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)
	router := config.CreateRouter(queries)
	mwStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mwStack(router),
	}

	err = server.ListenAndServe()
	if err != nil {
		panic("! Unable to start server !")
	}
}
