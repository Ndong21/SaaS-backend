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
INSErT INTO "purchases" (product_id, total_price, quantity, vendor_id, cashier_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateSale :one
INSErT INTO "sales" (product_id, unit_price, quantity, cashier_id)
VALUES ($1, $2, $3, $4)
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
WITH purchase_totals AS (
  SELECT 
    pr.id AS product_id,
    pr.product_name,
    c.category_name,
    SUM(p.quantity) AS total_purchased
  FROM purchases p
  JOIN products pr ON p.product_id = pr.id
  JOIN categories c ON pr.category_id = c.id
  GROUP BY pr.id, pr.product_name, c.category_name
),
sales_totals AS (
  SELECT 
    s.product_id,
    SUM(s.quantity) AS total_sold
  FROM sales s
  GROUP BY s.product_id
)
SELECT 
  pt.product_name,
  pt.category_name,
  COALESCE(pt.total_purchased, 0) - COALESCE(st.total_sold, 0) AS quantity_left
FROM 
  purchase_totals pt
LEFT JOIN 
  sales_totals st ON pt.product_id = st.product_id;


-- name: GetAllCategories :many
SELECT 
id,
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

-- name: GetAllPurchases :many
SELECT 
  p.id,
  pr.product_name,
  p.total_price,
  p.quantity,
  TO_CHAR(p.created_at, 'DD-MM-YYYY') AS "purchase_date",
  v.vendor_name,
  u.name AS cashier
FROM 
  purchases p
JOIN 
  products pr ON p.product_id = pr.id
LEFT JOIN 
  vendors v ON p.vendor_id = v.id
LEFT JOIN users u ON u.id = p.cashier_id;

-- name: GetAllSales :many
SELECT 
  s.id,
  pr.product_name,
  s.unit_price,
  s.quantity,
  s.unit_price * s.quantity AS total_price,
  TO_CHAR(s.created_at, 'DD-MM-YYYY') AS "Sale_date",
  u.name AS cashier
FROM 
  sales s
JOIN 
  products pr ON s.product_id = pr.id
LEFT JOIN users u ON u.id = s.cashier_id;



