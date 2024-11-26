package main

import (
	"backend/internal/config"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := config.Router()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.ServerPort),
		Handler: router,
	}
	server.ListenAndServe()
}
