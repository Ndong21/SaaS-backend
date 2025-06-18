-- foreign key constraints
ALTER TABLE "b_purchases" 
ADD FOREIGN KEY ("material_id") REFERENCES "materials" ("id");

ALTER TABLE "session_materials" 
ADD FOREIGN KEY ("material_id") REFERENCES "materials" ("id");

ALTER TABLE "session_materials" 
ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

ALTER TABLE "session_materials" 
ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id");

ALTER TABLE "session_products" 
ADD FOREIGN KEY ("product_id") REFERENCES "b_products" ("id");

ALTER TABLE "session_products" 
ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

ALTER TABLE "session_products" 
ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id");

ALTER TABLE "b_sales" 
ADD FOREIGN KEY ("product_id") REFERENCES "b_products" ("id");
