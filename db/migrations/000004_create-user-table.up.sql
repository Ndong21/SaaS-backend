CREATE TABLE "users" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "phone_number" VARCHAR,
  "password" VARCHAR NOT NULL,
  "role" VARCHAR NOT NULL
);
