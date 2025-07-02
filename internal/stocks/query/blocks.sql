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
INSERT INTO "b_sales" (product_id, selling_price, quantity, cashier_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateSessionMaterials :one
INSERT INTO "session_materials" (session_id, team_id,material_id, date,quantity)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: CreateSessionProducts :one
INSERT INTO "session_products" (session_id, team_id,product_id, date,quantity)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetMaterials :many
SELECT 
  id,
  material_name,
  unit,
  description
FROM materials;

-- name: GetMaterialPurchases :many
SELECT 
  bp.id,
  m.material_name,
  bp.quantity,
  bp.price,
  m.unit,
  TO_CHAR(bp.created_at, 'MM-DD-YYYY') AS purchase_date
FROM b_purchases bp
JOIN materials m ON bp.material_id = m.id;


-- name: GetBlocksProducts :many
WITH produced_totals AS (
  SELECT 
    sp.product_id,
    SUM(sp.quantity) AS quantity_produced
  FROM session_products sp
  GROUP BY sp.product_id
),
sales_totals AS (
  SELECT 
    bs.product_id,
    SUM(bs.quantity) AS quantity_sold
  FROM b_sales bs
  GROUP BY bs.product_id
)
SELECT 
  bp.id,
  bp.product_name,
  bp.description,
  --COALESCE(pt.quantity_produced, 0) AS quantity_produced,
  --COALESCE(st.quantity_sold, 0) AS quantity_sold,
  COALESCE(pt.quantity_produced, 0) - COALESCE(st.quantity_sold, 0) AS quantity_left
FROM b_products bp
LEFT JOIN produced_totals pt ON bp.id = pt.product_id
LEFT JOIN sales_totals st ON bp.id = st.product_id;

-- name: GetTeams :many
SELECT 
  id,
  team_name,
  phone_number,
  email
FROM teams;

-- name: GetSessions :many
SELECT 
  id,
  session,
  description
FROM sessions;

-- name: GetBlockSales :many
SELECT 
  bs.id,
  bp.product_name,
  bs.quantity,
  bs.selling_price,
  TO_CHAR(bs.created_at, 'MM-DD-YYYY') AS sale_date,
  bs.cashier_id
FROM b_sales bs
JOIN b_products bp ON bs.product_id = bp.id;


-- name: GetSessionMaterials :many
SELECT 
  -- Composite ID
  sm.session_id || '_' || sm.team_id || '_' || sm.material_id || '_' || TO_CHAR(sm.date, 'YYYYMMDD') AS id,
  s.session,
  t.team_name,
  m.material_name,
  sm.quantity,
  m.unit,
  TO_CHAR(sm.date, 'MM-DD-YYYY') AS used_date
FROM session_materials sm
JOIN sessions s ON sm.session_id = s.id
JOIN teams t ON sm.team_id = t.id
JOIN materials m ON sm.material_id = m.id;

-- name: GetSessionProducts :many
SELECT 
  -- Composite ID
  sp.session_id || '_' || sp.team_id || '_' || sp.product_id || '_' || TO_CHAR(sp.date, 'YYYYMMDD') AS id,
  s.session,
  t.team_name,
  bp.product_name,
  sp.quantity,
  TO_CHAR(sp.date, 'MM-DD-YYYY') AS production_date
FROM session_products sp
JOIN sessions s ON sp.session_id = s.id
JOIN teams t ON sp.team_id = t.id
JOIN b_products bp ON sp.product_id = bp.id;

-- name: UpdateBlockSale :one
UPDATE "b_sales"
SET
  product_id = $2,
  selling_price = $3,
  quantity = $4,
  cashier_id = $5
WHERE id = $1
RETURNING *;
