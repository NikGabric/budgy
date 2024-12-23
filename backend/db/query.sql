-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET username = $1,
email = $2
WHERE id = $3;

-- name: GetUserByUsername :one
SELECT id, username, email, password FROM users
WHERE username = $1 LIMIT 1;

-- File: transaction_types.sql
-- Create a new transaction type
-- name: CreateTransactionType :one
INSERT INTO transaction_types (user_id, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- Get a transaction type by ID
-- name: GetTransactionTypeByID :one
SELECT * FROM transaction_types
WHERE id = $1;

-- Get transaction types for a user
-- name: GetUserTransactionTypes :many
SELECT * FROM transaction_types
WHERE user_id = $1
ORDER BY created_at DESC;

-- Update a transaction type
-- name: UpdateTransactionType :one
UPDATE transaction_types
SET name = $2, description = $3
WHERE id = $1 AND user_id = $4
RETURNING *;

-- Delete a transaction type
-- name: DeleteTransactionType :exec
DELETE FROM transaction_types
WHERE id = $1 AND user_id = $2;

-- File: transactions.sql
-- Create a new transaction
-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id, 
    transaction_type_id, 
    amount, 
    description,
    transaction_date
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- Get a transaction by ID
-- name: GetTransactionByID :one
SELECT 
    t.*,
    u.username,
    tt.name as transaction_type_name
FROM transactions t
JOIN users u ON t.user_id = u.id
JOIN transaction_types tt ON t.transaction_type_id = tt.id
WHERE t.id = $1;

-- Get transactions for a user
-- name: GetUserTransactions :many
SELECT 
    t.*,
    tt.name as transaction_type_name
FROM transactions t
JOIN transaction_types tt ON t.transaction_type_id = tt.id
WHERE t.user_id = $1
AND t.transaction_type_id = COALESCE(sqlc.narg('transaction_type_id'), t.transaction_type_id)
AND t.transaction_date <= COALESCE(sqlc.narg('to_date'), t.transaction_date)
AND t.transaction_date >= COALESCE(sqlc.narg('from_date'), t.transaction_date)
ORDER BY t.transaction_date DESC
LIMIT COALESCE(sqlc.narg('limit')::int, 10) OFFSET $2;

-- Get transactions count
-- name: GetTransactionsCount :many
SELECT COUNT(*)
FROM transactions t
WHERE t.user_id = $1
AND t.transaction_type_id = COALESCE(sqlc.narg('transaction_type_id'), t.transaction_type_id)
AND t.transaction_date <= COALESCE(sqlc.narg('to_date'), t.transaction_date)
AND t.transaction_date >= COALESCE(sqlc.narg('from_date'), t.transaction_date);

-- Update a transaction
-- name: UpdateTransaction :one
UPDATE transactions
SET
    transaction_type_id = COALESCE(sqlc.narg('transaction_type_id'), transaction_type_id),
    amount = COALESCE(sqlc.narg('amount'), amount),
    description = COALESCE(sqlc.narg('description'), description),
    transaction_date = COALESCE(sqlc.narg('transaction_date'), transaction_date)
WHERE id = $1 AND user_id = $2
RETURNING *;

-- Delete a transaction
-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1 AND user_id = $2;