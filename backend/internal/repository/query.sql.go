// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id, 
    transaction_type_id, 
    amount, 
    description,
    transaction_date
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, transaction_type_id, amount, description, transaction_date, created_at, updated_at
`

type CreateTransactionParams struct {
	UserID            int32              `json:"user_id"`
	TransactionTypeID int32              `json:"transaction_type_id"`
	Amount            pgtype.Numeric     `json:"amount"`
	Description       pgtype.Text        `json:"description"`
	TransactionDate   pgtype.Timestamptz `json:"transaction_date"`
}

// File: transactions.sql
// Create a new transaction
func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.UserID,
		arg.TransactionTypeID,
		arg.Amount,
		arg.Description,
		arg.TransactionDate,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransactionTypeID,
		&i.Amount,
		&i.Description,
		&i.TransactionDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTransactionType = `-- name: CreateTransactionType :one
INSERT INTO transaction_types (user_id, name, description)
VALUES ($1, $2, $3)
RETURNING id, user_id, name, description, created_at, updated_at
`

type CreateTransactionTypeParams struct {
	UserID      int32       `json:"user_id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
}

// File: transaction_types.sql
// Create a new transaction type
func (q *Queries) CreateTransactionType(ctx context.Context, arg CreateTransactionTypeParams) (TransactionType, error) {
	row := q.db.QueryRow(ctx, createTransactionType, arg.UserID, arg.Name, arg.Description)
	var i TransactionType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1, $2, $3
)
RETURNING id, username, email, password, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1 AND user_id = $2
`

type DeleteTransactionParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
}

// Delete a transaction
func (q *Queries) DeleteTransaction(ctx context.Context, arg DeleteTransactionParams) error {
	_, err := q.db.Exec(ctx, deleteTransaction, arg.ID, arg.UserID)
	return err
}

const deleteTransactionType = `-- name: DeleteTransactionType :exec
DELETE FROM transaction_types
WHERE id = $1 AND user_id = $2
`

type DeleteTransactionTypeParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
}

// Delete a transaction type
func (q *Queries) DeleteTransactionType(ctx context.Context, arg DeleteTransactionTypeParams) error {
	_, err := q.db.Exec(ctx, deleteTransactionType, arg.ID, arg.UserID)
	return err
}

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT 
    t.id, t.user_id, t.transaction_type_id, t.amount, t.description, t.transaction_date, t.created_at, t.updated_at,
    u.username,
    tt.name as transaction_type_name
FROM transactions t
JOIN users u ON t.user_id = u.id
JOIN transaction_types tt ON t.transaction_type_id = tt.id
WHERE t.id = $1
`

type GetTransactionByIDRow struct {
	ID                  int32              `json:"id"`
	UserID              int32              `json:"user_id"`
	TransactionTypeID   int32              `json:"transaction_type_id"`
	Amount              pgtype.Numeric     `json:"amount"`
	Description         pgtype.Text        `json:"description"`
	TransactionDate     pgtype.Timestamptz `json:"transaction_date"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	UpdatedAt           pgtype.Timestamptz `json:"updated_at"`
	Username            string             `json:"username"`
	TransactionTypeName string             `json:"transaction_type_name"`
}

// Get a transaction by ID
func (q *Queries) GetTransactionByID(ctx context.Context, id int32) (GetTransactionByIDRow, error) {
	row := q.db.QueryRow(ctx, getTransactionByID, id)
	var i GetTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransactionTypeID,
		&i.Amount,
		&i.Description,
		&i.TransactionDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.TransactionTypeName,
	)
	return i, err
}

const getTransactionTypeByID = `-- name: GetTransactionTypeByID :one
SELECT id, user_id, name, description, created_at, updated_at FROM transaction_types
WHERE id = $1
`

// Get a transaction type by ID
func (q *Queries) GetTransactionTypeByID(ctx context.Context, id int32) (TransactionType, error) {
	row := q.db.QueryRow(ctx, getTransactionTypeByID, id)
	var i TransactionType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password FROM users
WHERE username = $1 LIMIT 1
`

type GetUserByUsernameRow struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserTransactionTypes = `-- name: GetUserTransactionTypes :many
SELECT id, user_id, name, description, created_at, updated_at FROM transaction_types
WHERE user_id = $1
ORDER BY created_at DESC
`

// Get transaction types for a user
func (q *Queries) GetUserTransactionTypes(ctx context.Context, userID int32) ([]TransactionType, error) {
	rows, err := q.db.Query(ctx, getUserTransactionTypes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TransactionType
	for rows.Next() {
		var i TransactionType
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTransactions = `-- name: GetUserTransactions :many
SELECT 
    t.id, t.user_id, t.transaction_type_id, t.amount, t.description, t.transaction_date, t.created_at, t.updated_at,
    tt.name as transaction_type_name
FROM transactions t
JOIN transaction_types tt ON t.transaction_type_id = tt.id
WHERE t.user_id = $1
AND t.transaction_type_id = COALESCE($3, t.transaction_type_id)
ORDER BY t.transaction_date DESC
LIMIT COALESCE($4::int, 10) OFFSET $2
`

type GetUserTransactionsParams struct {
	UserID            int32       `json:"user_id"`
	Offset            int32       `json:"offset"`
	TransactionTypeID pgtype.Int4 `json:"transaction_type_id"`
	Limit             pgtype.Int4 `json:"limit"`
}

type GetUserTransactionsRow struct {
	ID                  int32              `json:"id"`
	UserID              int32              `json:"user_id"`
	TransactionTypeID   int32              `json:"transaction_type_id"`
	Amount              pgtype.Numeric     `json:"amount"`
	Description         pgtype.Text        `json:"description"`
	TransactionDate     pgtype.Timestamptz `json:"transaction_date"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	UpdatedAt           pgtype.Timestamptz `json:"updated_at"`
	TransactionTypeName string             `json:"transaction_type_name"`
}

// Get transactions for a user
func (q *Queries) GetUserTransactions(ctx context.Context, arg GetUserTransactionsParams) ([]GetUserTransactionsRow, error) {
	rows, err := q.db.Query(ctx, getUserTransactions,
		arg.UserID,
		arg.Offset,
		arg.TransactionTypeID,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTransactionsRow
	for rows.Next() {
		var i GetUserTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TransactionTypeID,
			&i.Amount,
			&i.Description,
			&i.TransactionDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TransactionTypeName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, created_at, updated_at FROM users
ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions
SET
    transaction_type_id = COALESCE($3, transaction_type_id),
    amount = COALESCE($4, amount),
    description = COALESCE($5, description),
    transaction_date = COALESCE($6, transaction_date)
WHERE id = $1 AND user_id = $2
RETURNING id, user_id, transaction_type_id, amount, description, transaction_date, created_at, updated_at
`

type UpdateTransactionParams struct {
	ID                int32              `json:"id"`
	UserID            int32              `json:"user_id"`
	TransactionTypeID pgtype.Int4        `json:"transaction_type_id"`
	Amount            pgtype.Numeric     `json:"amount"`
	Description       pgtype.Text        `json:"description"`
	TransactionDate   pgtype.Timestamptz `json:"transaction_date"`
}

// Update a transaction
func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction,
		arg.ID,
		arg.UserID,
		arg.TransactionTypeID,
		arg.Amount,
		arg.Description,
		arg.TransactionDate,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransactionTypeID,
		&i.Amount,
		&i.Description,
		&i.TransactionDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTransactionType = `-- name: UpdateTransactionType :one
UPDATE transaction_types
SET name = $2, description = $3
WHERE id = $1 AND user_id = $4
RETURNING id, user_id, name, description, created_at, updated_at
`

type UpdateTransactionTypeParams struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	UserID      int32       `json:"user_id"`
}

// Update a transaction type
func (q *Queries) UpdateTransactionType(ctx context.Context, arg UpdateTransactionTypeParams) (TransactionType, error) {
	row := q.db.QueryRow(ctx, updateTransactionType,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.UserID,
	)
	var i TransactionType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $1,
email = $2
WHERE id = $3
`

type UpdateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Username, arg.Email, arg.ID)
	return err
}
