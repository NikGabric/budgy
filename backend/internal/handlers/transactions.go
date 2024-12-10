package handlers

import (
	"backend/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type TransactionsHandler struct {
	q *repository.Queries
}

func NewTransactionsHandler(q *repository.Queries) *TransactionsHandler {
	return &TransactionsHandler{q: q}
}

func (h *TransactionsHandler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := h.q.GetTransactionByID(context.Background(), int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tJson, err := json.Marshal(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tJson)
}

func (h *TransactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tDto repository.CreateTransactionParams
	err := json.NewDecoder(r.Body).Decode(&tDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(tDto)

	t, err := h.q.CreateTransaction(context.Background(), tDto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ttJson, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(ttJson)
}
