package handlers

import (
	"backend/internal/repository"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type TransactionTypesHandler struct {
	q *repository.Queries
}

func NewTransactionTypesHandler(q *repository.Queries) *TransactionTypesHandler {
	return &TransactionTypesHandler{q: q}
}

func (h *TransactionTypesHandler) GetTransactionTypeById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transactionType, err := h.q.GetTransactionTypeByID(context.Background(), int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ttJson, err := json.Marshal(transactionType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(ttJson)
}

func (h *TransactionTypesHandler) CreateTransactionType(w http.ResponseWriter, r *http.Request) {
	var ttDto repository.CreateTransactionTypeParams
	err := json.NewDecoder(r.Body).Decode(&ttDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tt, err := h.q.CreateTransactionType(context.Background(), ttDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ttJson, err := json.Marshal(tt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(ttJson)
}
