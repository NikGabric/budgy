package main

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/repository"
	"context"
	"fmt"
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
