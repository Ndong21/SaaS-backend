CREATE TABLE "vendors" (
 "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
 "vendor_name" varchar NOT NULL
);