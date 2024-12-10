package handlers

import (
	"backend/internal/repository"
	"context"
	"encoding/json"
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

func (h *TransactionsHandler) GetTransactionsForUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int32)
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := repository.GetUserTransactionsParams{UserID: userId, Limit: int32(limit)}

	transactions, err := h.q.GetUserTransactions(context.Background(), params)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	transactionsJson, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(transactionsJson)
}

func (h *TransactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tDto repository.CreateTransactionParams
	err := json.NewDecoder(r.Body).Decode(&tDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := h.q.CreateTransaction(context.Background(), tDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ttJson, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(ttJson)
}
