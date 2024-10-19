// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package query

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  user_id, mail, name, code, hashed_password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING user_id, mail, name, code, hashed_password
`

type CreateUserParams struct {
	UserID         string
	Mail           string
	Name           string
	Code           string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Mail,
		arg.Name,
		arg.Code,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.Code,
		&i.HashedPassword,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, mail, name, code, hashed_password FROM users
WHERE mail = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, mail string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, mail)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.Code,
		&i.HashedPassword,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, mail, name, code, hashed_password FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.Code,
		&i.HashedPassword,
	)
	return i, err
}

const updateCode = `-- name: UpdateCode :exec
UPDATE users
  set code = $2
WHERE user_id = $1
`

type UpdateCodeParams struct {
	UserID string
	Code   string
}

func (q *Queries) UpdateCode(ctx context.Context, arg UpdateCodeParams) error {
	_, err := q.db.Exec(ctx, updateCode, arg.UserID, arg.Code)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
  set hashed_password = $2
WHERE user_id = $1
`

type UpdatePasswordParams struct {
	UserID         string
	HashedPassword string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.UserID, arg.HashedPassword)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
  set mail = $2,
  name = $3,
  hashed_password = $4
WHERE user_id = $1
`

type UpdateUserParams struct {
	UserID         string
	Mail           string
	Name           string
	HashedPassword string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.UserID,
		arg.Mail,
		arg.Name,
		arg.HashedPassword,
	)
	return err
}
