-- name: CreateMaterial :one
INSERT INTO "materials" (material_name, unit, description)
VALUES ($1,$2,$3)
RETURNING *;

-- name: CreateMaterialPurchase :one
INSERT INTO "b_purchases" (material_id, quantity, price)
VALUES ($1,$2,$3)
RETURNING *;

-- name: CreateBlocksProduct :one
INSERT INTO "b_products" (product_name, description)
VALUES ($1,$2)
RETURNING *;

-- name: CreateTeam :one
INSErT INTO "teams" (team_name, phone_number, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateSession :one
INSErT INTO "sessions" (session, description)
VALUES ($1, $2)
RETURNING *;

-- name: CreateBlockSale :one
INSERT INTO "b_sales" (product_id, selling_price, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateSessionMaterials :one
INSERT INTO "session_materials" (session_id, team_id,material_id, date,quantity)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: CreateSessionProducts :one
INSERT INTO "session_products" (session_id, team_id,product_id, date,quantity)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;