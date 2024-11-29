package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"net/http"
)

func main() {
	config := config.GetConfig()

	server := http.NewServeMux()

	server.HandleFunc("/", handlers.GetStatus)

	err := http.ListenAndServe(":"+config.ServerPort, server)
	if err != nil {
		panic("! Unable to start server !")
	}
}
