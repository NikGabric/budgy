// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32
	Username  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}
