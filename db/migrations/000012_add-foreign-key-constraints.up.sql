ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "sales" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "catalog" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");