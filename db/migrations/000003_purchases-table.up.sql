CREATE TABLE "purchases" ( 
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36) NOT NULL, 
 "total_price" INT NOT NULL, 
 "quantity" INT NOT NULL, 
 "vendor_id" VARCHAR(36), 
 "created_at" TIMESTAMP DEFAULT now()
);