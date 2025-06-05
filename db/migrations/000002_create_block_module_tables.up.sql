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

-- sessions: Predefined time slots (no team info here)
CREATE TABLE "sessions" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "session" VARCHAR NOT NULL  -- e.g., '7am–9am'
);

-- session_materials: Materials a team used during a session
CREATE TABLE "session_materials" (
  "session_id" VARCHAR(36) NOT NULL,
  "team_id" VARCHAR(36) NOT NULL,
  "material_id" VARCHAR(36) NOT NULL,
  "date" DATE NOT NULL,
  "quantity" INT NOT NULL,
  PRIMARY KEY ("session_id", "team_id", "material_id", "date")
);
-- session_products: Products a team made during a session
CREATE TABLE "session_products" (
  "session_id" VARCHAR(36) NOT NULL,
  "team_id" VARCHAR(36) NOT NULL,
  "product_id" VARCHAR(36) NOT NULL,
  "date" DATE NOT NULL,
  "quantity" INT NOT NULL,
  PRIMARY KEY ("session_id", "team_id", "product_id", "date")
);

-- sales
CREATE TABLE "b_sales" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "product_id" VARCHAR(36) NOT NULL,
  "quantity" INT NOT NULL,
  "selling_price" DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);
