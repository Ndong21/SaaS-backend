CREATE TABLE IF NOT EXISTS  "sales" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36) NOT NULL,
 "unit_price" INT NOT NULL,
 "quantity" int NOT NULL,
 "create_at" TIMESTAMP DEFAULT now(),
 "cashier_id" VARCHAR(36)
);