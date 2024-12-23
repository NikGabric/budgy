package handlers

import (
	"backend/internal/helpers"
	"backend/internal/repository"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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
	params := repository.GetUserTransactionsParams{
		UserID: userId,
	}
	totalCountParams := repository.GetTransactionsCountParams{
		UserID: userId,
	}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		limit, err := strconv.ParseInt(limitParam, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params.Limit = pgtype.Int4{Int32: int32(limit), Valid: true}
	}

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		page, err := strconv.ParseInt(pageParam, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params.Offset = int32((page - 1) * int64(params.Limit.Int32))
	}

	if ttIdParam := r.URL.Query().Get("transaction_type_id"); ttIdParam != "" {
		ttId, err := strconv.ParseInt(ttIdParam, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params.TransactionTypeID = pgtype.Int4{Int32: int32(ttId), Valid: true}
		totalCountParams.TransactionTypeID = pgtype.Int4{Int32: int32(ttId), Valid: true}
	}

	if ttFromDate := r.URL.Query().Get("from_date"); ttFromDate != "" {
		fromDate, err := time.Parse("2006-01-02", ttFromDate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params.FromDate = pgtype.Timestamptz{Time: fromDate, Valid: true}
		totalCountParams.FromDate = pgtype.Timestamptz{Time: fromDate, Valid: true}
	}

	if ttToDate := r.URL.Query().Get("to_date"); ttToDate != "" {
		toDate, err := time.Parse("2006-01-02", ttToDate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params.ToDate = pgtype.Timestamptz{Time: toDate, Valid: true}
		totalCountParams.ToDate = pgtype.Timestamptz{Time: toDate, Valid: true}
	}

	transactions, err := h.q.GetUserTransactions(context.Background(), params)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if transactions == nil {
		transactions = make([]repository.GetUserTransactionsRow, 0)
	}

	totalCount, err := h.q.GetTransactionsCount(context.Background(), totalCountParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := helpers.CreatePaginationResponse(transactions, int(totalCount[0]))

	respJson, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (h *TransactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tDto repository.CreateTransactionParams
	err := json.NewDecoder(r.Body).Decode(&tDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tDto.UserID = r.Context().Value("userId").(int32)

	t, err := h.q.CreateTransaction(context.Background(), tDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tJson, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(tJson)
}

func (h *TransactionsHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var tDto repository.UpdateTransactionParams
	err := json.NewDecoder(r.Body).Decode(&tDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tDto.UserID = r.Context().Value("userId").(int32)

	t, err := h.q.UpdateTransaction(context.Background(), tDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tJson, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tJson)
}

func (h *TransactionsHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var tDto repository.DeleteTransactionParams
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tDto.ID = int32(id)
	tDto.UserID = r.Context().Value("userId").(int32)

	err = h.q.DeleteTransaction(context.Background(), tDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success"))
}
