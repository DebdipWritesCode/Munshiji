// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, email, password_hash
) VALUES (
  $1, $2, $3
)
RETURNING id, name, email, password_hash, created_at, is_email_verified
`

type CreateUserParams struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, created_at, is_email_verified FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, password_hash, created_at, is_email_verified FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = COALESCE($1, name),
    email = COALESCE($2, email),
    password_hash = COALESCE($3, password_hash)
WHERE id = $4
RETURNING id, name, email, password_hash, created_at, is_email_verified
`

type UpdateUserParams struct {
	Name         sql.NullString `json:"name"`
	Email        sql.NullString `json:"email"`
	PasswordHash sql.NullString `json:"password_hash"`
	ID           int32          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Email,
		arg.PasswordHash,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const verifyUserEmail = `-- name: VerifyUserEmail :exec
UPDATE users
SET is_email_verified = TRUE
WHERE id = $1
`

func (q *Queries) VerifyUserEmail(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.verifyUserEmailStmt, verifyUserEmail, id)
	return err
}
