package handlers

import (
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("Server healthy!"))
	w.WriteHeader(http.StatusOK)
	return
}
