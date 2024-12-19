package main

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/repository"
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()

	dbpool, err := pgxpool.New(ctx, cfg.DbConnString)
	if err != nil {
		panic("! Unable to create connection pool !")
	}
	defer dbpool.Close()

	queries := repository.New(dbpool)
	router := config.CreateRouter(queries)
	mwStack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mwStack(router),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		panic("! Unable to start server !")
	}
}
