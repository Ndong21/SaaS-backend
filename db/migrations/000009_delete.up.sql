-- Delete from child/dependent tables first
DELETE FROM session_products;
DELETE FROM session_materials;
DELETE FROM sessions;

DELETE FROM b_sales;
DELETE FROM b_purchases;

DELETE FROM b_products;
DELETE FROM materials;
DELETE FROM teams;

