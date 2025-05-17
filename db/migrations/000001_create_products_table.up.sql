CREATE TABLE "products" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "category_id" VARCHAR(36) DEFAULT gen_random_uuid()::varchar(36) NOT NULL,
 "product_name" varchar NOT NULL,
 "quantity" integer NOT NULL 
);