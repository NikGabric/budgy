package config

import (
	"backend/internal/handler"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", handler.StatusHandler)
	router.HandleFunc("/test", handler.TestHandler)

	return router
}
