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
INSErT INTO "sales" (product_id, unit_price, quantity, cashier_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateVendor :one
INSErT INTO "vendors" (vendor_name, vendor_location, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateCatalog :one
INSErT INTO "catalog" (product_id, unit_price)
VALUES ($1, $2)
RETURNING *;

-- name: GetAllProducts :many
WITH purchase_totals AS (
  SELECT 
    product_id,
    SUM(quantity) AS total_purchased
  FROM purchases
  GROUP BY product_id
),
sales_totals AS (
  SELECT 
    product_id,
    SUM(quantity) AS total_sold
  FROM sales
  GROUP BY product_id
)
SELECT 
  pr.id AS product_id,
  pr.product_name,
  c.category_name,
  COALESCE(pur.total_purchased, 0) - COALESCE(sal.total_sold, 0) AS quantity_left
FROM 
  products pr
JOIN 
  categories c ON pr.category_id = c.id
LEFT JOIN 
  purchase_totals pur ON pr.id = pur.product_id
LEFT JOIN 
  sales_totals sal ON pr.id = sal.product_id;



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
vendor_location,
COALESCE(description, '') AS description
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
  v.vendor_name
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
  TO_CHAR(s.created_at, 'DD-MM-YYYY') AS "Sale_date"
  --u.name AS cashier
FROM 
  sales s
JOIN 
  products pr ON s.product_id = pr.id
LEFT JOIN users u ON u.id = s.cashier_id;

-- name: DeleteCatalog :exec
DELETE FROM catalog
WHERE id = $1;


-- name: TotalSales :one
SELECT SUM(unit_price * quantity) AS total_sales
FROM sales;

-- name: CountSalesTransactions :one
SELECT COUNT(*) AS transaction_count
FROM sales;

-- name: Top5BestSellingProductsByRevenue :many
SELECT 
  p.id AS product_id,
  p.product_name,
  SUM(s.quantity) AS total_quantity_sold,
  SUM(s.quantity * s.unit_price) AS total_revenue
FROM sales s
JOIN products p ON s.product_id = p.id
GROUP BY p.id, p.product_name
ORDER BY total_revenue DESC
LIMIT 5;

-- name: UpdateSales :one
UPDATE sales
SET unit_price = $1, quantity = $2
WHERE id = $3
RETURNING *;

-- name: DeleteSale :exec
DELETE FROM sales
WHERE id = $1;

-- name: UpdatePurchase :one
UPDATE purchases
SET total_price = $1, quantity = $2
WHERE id = $3
RETURNING *;

-- name: DeletePurchase :exec
DELETE FROM purchases
WHERE id = $1;

-- name: UpdateProduct :one
UPDATE products
SET category_id = $1, product_name = $2
WHERE id = $3
RETURNING *;

-- name: Deleteproduct :exec
DELETE FROM products
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories
SET category_name = $1, category_description = $2
WHERE id = $3
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

