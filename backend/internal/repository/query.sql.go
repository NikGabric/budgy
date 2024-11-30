// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"
)

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
	Username string
	Email    string
	Password string
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
	ID       int32
	Username string
	Email    string
	Password string
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $1,
email = $2
WHERE id = $3
`

type UpdateUserParams struct {
	Username string
	Email    string
	ID       int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Username, arg.Email, arg.ID)
	return err
}
