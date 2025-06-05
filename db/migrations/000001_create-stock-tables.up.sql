-- categories
CREATE TABLE "categories" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "category_name" VARCHAR NOT NULL,
 "category_description" VARCHAR NOT NULL
);

-- products
CREATE TABLE "products" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "category_id" VARCHAR(36) NOT NULL,
 "product_name" VARCHAR NOT NULL
);

-- vendors
CREATE TABLE "vendors" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "vendor_name" VARCHAR NOT NULL,
 "vendor_location" VARCHAR NOT NULL
);

-- purchases
CREATE TABLE "purchases" ( 
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36) NOT NULL, 
 "total_price" INT NOT NULL, 
 "quantity" INT NOT NULL, 
 "vendor_id" VARCHAR(36), 
 "created_at" TIMESTAMP DEFAULT now()
);

-- sales
CREATE TABLE "sales" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36) NOT NULL,
 "unit_price" INT NOT NULL,
 "quantity" INT NOT NULL,
 "created_at" TIMESTAMP DEFAULT now(),
 "cashier_id" VARCHAR(36)
);

-- catalog
CREATE TABLE "catalog" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
 "product_id" VARCHAR(36) NOT NULL,
 "unit_price" INT NOT NULL
);


-- foreign key constraints
ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "sales" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "catalog" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");