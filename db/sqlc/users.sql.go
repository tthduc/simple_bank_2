// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
  $1, $2, $3, $4
)
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE($1, hashed_password),
    full_name = COALESCE($2, full_name),
    email = COALESCE($3, email),
    password_changed_at = COALESCE($4, password_changed_at)
WHERE
    username = $5
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashedPassword"`
	FullName          sql.NullString `json:"fullName"`
	Email             sql.NullString `json:"email"`
	PasswordChangedAt sql.NullTime   `json:"passwordChangedAt"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.PasswordChangedAt,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
