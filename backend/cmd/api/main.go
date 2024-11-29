package main

import (
	"backend/internal/handlers"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", handlers.GetStatus)

	err := http.ListenAndServe(":3030", server)
	if err != nil {
		panic("! Unable to start server !")
	}
}
