CREATE TABLE "catalog" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36),
 "unit_price" INT NOT NULL
);