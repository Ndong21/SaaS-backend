-- materials
CREATE TABLE "materials" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "material_name" VARCHAR NOT NULL,
  "unit" VARCHAR NOT NULL,
  "description" VARCHAR
);

-- b_purchases (material purchases)
CREATE TABLE "b_purchases" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "material_id" VARCHAR(36) NOT NULL,
  "quantity" INT NOT NULL,
  "price" DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);

-- b_products (produced goods)
CREATE TABLE "b_products" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "product_name" VARCHAR NOT NULL,
  "description" TEXT
);

-- teams
CREATE TABLE "teams" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "team_name" VARCHAR NOT NULL,
  "phone_number" VARCHAR,
  "email" VARCHAR
);

-- sessions
CREATE TABLE "sessions" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "team_id" VARCHAR(36) NOT NULL,
  "start_time" TIMESTAMP NOT NULL,
  "end_time" TIMESTAMP NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);

-- session_materials
CREATE TABLE "session_materials" (
  "session_id" VARCHAR(36) NOT NULL,
  "material_id" VARCHAR(36) NOT NULL,
  "quantity" INT NOT NULL,
  PRIMARY KEY ("session_id", "material_id")
);

-- session_products
CREATE TABLE "session_products" (
  "session_id" VARCHAR(36) NOT NULL,
  "product_id" VARCHAR(36) NOT NULL,
  "quantity" INT NOT NULL,
  PRIMARY KEY ("session_id", "product_id")
);

-- sales
CREATE TABLE "b_sales" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "product_id" VARCHAR(36) NOT NULL,
  "quantity" INT NOT NULL,
  "selling_price" DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);
