-- name: CreateCategory :one
INSERT INTO "categories" (category_name, category_description)
VALUES ($1,$2)
RETURNING *;

-- name: CreateProduct :one
INSERT INTO "products" (category_id, product_name)
VALUES ($1,$2)
RETURNING *;

-- name: LoadTime :one
SELECT NOW();

-- name: CreatePurchase :one
INSErT INTO "purchases" (product_id, total_price, quantity, vendor_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateSale :one
INSErT INTO "sales" (product_id, unit_price, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateVendor :one
INSErT INTO "vendors" (vendor_name, vendor_location)
VALUES ($1, $2)
RETURNING *;

-- name: CreateCatalog :one
INSErT INTO "catalog" (product_id, unit_price)
VALUES ($1, $2)
RETURNING *;

-- name: GetAllProducts :many
SELECT 
p.product_name,
c.category_name
FROM products p 
JOIN categories c ON p.category_id = c.id;

-- name: GetAllCategories :many
SELECT 
category_name,
category_description
FROM categories;

-- name: GetAllVendors :many
SELECT 
id,
vendor_name,
vendor_location
FROM vendors;


-- name: GetCatalog :many
SELECT 
p.product_name,
c.unit_price
FROM products p 
JOIN catalog c ON p.id = c.product_id;