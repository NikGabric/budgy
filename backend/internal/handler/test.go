package handler

import (
	"io"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET received!"))
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
