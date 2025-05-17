CREATE TABLE "categories" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "category_name" varchar NOT NULL,
 "category_description" varchar NOT NULL
);