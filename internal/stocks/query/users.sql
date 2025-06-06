-- name: CreateUser :one
INSERT INTO "users" (name, email, phone_number, password, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- -- name: LogIn :one
-- INSERT INTO "users" (email, password)
-- VALUES ($1, $2)
-- RETURNING *;

-- name: SelectRequestedUser :one
SELECT id, email, password,role FROM "users" 
WHERE email = $1;