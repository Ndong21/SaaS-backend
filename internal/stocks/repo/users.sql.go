// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package repo

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (name, email, phone_number, password, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, phone_number, password, role
`

type CreateUserParams struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Password    string  `json:"password"`
	Role        string  `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.PhoneNumber,
		arg.Password,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const selectRequestedUser = `-- name: SelectRequestedUser :one

SELECT id, email, password,role FROM "users" 
WHERE email = $1
`

type SelectRequestedUserRow struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// -- name: LogIn :one
// INSERT INTO "users" (email, password)
// VALUES ($1, $2)
// RETURNING *;
func (q *Queries) SelectRequestedUser(ctx context.Context, email string) (SelectRequestedUserRow, error) {
	row := q.db.QueryRow(ctx, selectRequestedUser, email)
	var i SelectRequestedUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
	)
	return i, err
}
