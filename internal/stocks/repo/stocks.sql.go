// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: stocks.sql

package repo

import (
	"context"
)

const createCatalog = `-- name: CreateCatalog :one
INSErT INTO "catalog" (product_id, unit_price)
VALUES ($1, $2)
RETURNING id, product_id, unit_price
`

type CreateCatalogParams struct {
	ProductID string `json:"product_id"`
	UnitPrice int32  `json:"unit_price"`
}

func (q *Queries) CreateCatalog(ctx context.Context, arg CreateCatalogParams) (Catalog, error) {
	row := q.db.QueryRow(ctx, createCatalog, arg.ProductID, arg.UnitPrice)
	var i Catalog
	err := row.Scan(&i.ID, &i.ProductID, &i.UnitPrice)
	return i, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO "categories" (category_name, category_description)
VALUES ($1,$2)
RETURNING id, category_name, category_description
`

type CreateCategoryParams struct {
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, arg.CategoryName, arg.CategoryDescription)
	var i Category
	err := row.Scan(&i.ID, &i.CategoryName, &i.CategoryDescription)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO "products" (category_id, product_name)
VALUES ($1,$2)
RETURNING id, category_id, product_name
`

type CreateProductParams struct {
	CategoryID  string `json:"category_id"`
	ProductName string `json:"product_name"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.CategoryID, arg.ProductName)
	var i Product
	err := row.Scan(&i.ID, &i.CategoryID, &i.ProductName)
	return i, err
}

const createPurchase = `-- name: CreatePurchase :one
INSErT INTO "purchases" (product_id, total_price, quantity, vendor_id)
VALUES ($1, $2, $3, $4)
RETURNING id, product_id, total_price, quantity, vendor_id, created_at
`

type CreatePurchaseParams struct {
	ProductID  string  `json:"product_id"`
	TotalPrice int32   `json:"total_price"`
	Quantity   int32   `json:"quantity"`
	VendorID   *string `json:"vendor_id"`
}

func (q *Queries) CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (Purchase, error) {
	row := q.db.QueryRow(ctx, createPurchase,
		arg.ProductID,
		arg.TotalPrice,
		arg.Quantity,
		arg.VendorID,
	)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.TotalPrice,
		&i.Quantity,
		&i.VendorID,
		&i.CreatedAt,
	)
	return i, err
}

const createSale = `-- name: CreateSale :one
INSErT INTO "sales" (product_id, unit_price, quantity)
VALUES ($1, $2, $3)
RETURNING id, product_id, unit_price, quantity, created_at, cashier_id
`

type CreateSaleParams struct {
	ProductID string `json:"product_id"`
	UnitPrice int32  `json:"unit_price"`
	Quantity  int32  `json:"quantity"`
}

func (q *Queries) CreateSale(ctx context.Context, arg CreateSaleParams) (Sale, error) {
	row := q.db.QueryRow(ctx, createSale, arg.ProductID, arg.UnitPrice, arg.Quantity)
	var i Sale
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Quantity,
		&i.CreatedAt,
		&i.CashierID,
	)
	return i, err
}

const createVendor = `-- name: CreateVendor :one
INSErT INTO "vendors" (vendor_name, vendor_location)
VALUES ($1, $2)
RETURNING id, vendor_name, vendor_location
`

type CreateVendorParams struct {
	VendorName     string `json:"vendor_name"`
	VendorLocation string `json:"vendor_location"`
}

func (q *Queries) CreateVendor(ctx context.Context, arg CreateVendorParams) (Vendor, error) {
	row := q.db.QueryRow(ctx, createVendor, arg.VendorName, arg.VendorLocation)
	var i Vendor
	err := row.Scan(&i.ID, &i.VendorName, &i.VendorLocation)
	return i, err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT 
category_name,
category_description
FROM categories
`

type GetAllCategoriesRow struct {
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

func (q *Queries) GetAllCategories(ctx context.Context) ([]GetAllCategoriesRow, error) {
	rows, err := q.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllCategoriesRow{}
	for rows.Next() {
		var i GetAllCategoriesRow
		if err := rows.Scan(&i.CategoryName, &i.CategoryDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT 
p.product_name,
c.category_name
FROM products p 
JOIN categories c ON p.category_id = c.id
`

type GetAllProductsRow struct {
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
}

func (q *Queries) GetAllProducts(ctx context.Context) ([]GetAllProductsRow, error) {
	rows, err := q.db.Query(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllProductsRow{}
	for rows.Next() {
		var i GetAllProductsRow
		if err := rows.Scan(&i.ProductName, &i.CategoryName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPurchases = `-- name: GetAllPurchases :many
SELECT 
  p.id,
  pr.product_name,
  p.total_price,
  p.quantity,
  TO_CHAR(p.created_at, 'DD-MM-YYYY') AS "purchase_date",
  v.vendor_name
FROM 
  purchases p
JOIN 
  products pr ON p.product_id = pr.id
LEFT JOIN 
  vendors v ON p.vendor_id = v.id
`

type GetAllPurchasesRow struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	TotalPrice   int32   `json:"total_price"`
	Quantity     int32   `json:"quantity"`
	PurchaseDate string  `json:"purchase_date"`
	VendorName   *string `json:"vendor_name"`
}

func (q *Queries) GetAllPurchases(ctx context.Context) ([]GetAllPurchasesRow, error) {
	rows, err := q.db.Query(ctx, getAllPurchases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllPurchasesRow{}
	for rows.Next() {
		var i GetAllPurchasesRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.TotalPrice,
			&i.Quantity,
			&i.PurchaseDate,
			&i.VendorName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSales = `-- name: GetAllSales :many
SELECT 
  s.id,
  pr.product_name,
  s.unit_price,
  s.quantity,
  s.unit_price * s.quantity AS total_price,
  TO_CHAR(s.created_at, 'DD-MM-YYYY') AS "Sale_date"
FROM 
  sales s
JOIN 
  products pr ON s.product_id = pr.id
`

type GetAllSalesRow struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	UnitPrice   int32  `json:"unit_price"`
	Quantity    int32  `json:"quantity"`
	TotalPrice  int32  `json:"total_price"`
	SaleDate    string `json:"Sale_date"`
}

func (q *Queries) GetAllSales(ctx context.Context) ([]GetAllSalesRow, error) {
	rows, err := q.db.Query(ctx, getAllSales)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllSalesRow{}
	for rows.Next() {
		var i GetAllSalesRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.UnitPrice,
			&i.Quantity,
			&i.TotalPrice,
			&i.SaleDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllVendors = `-- name: GetAllVendors :many
SELECT 
id,
vendor_name,
vendor_location
FROM vendors
`

func (q *Queries) GetAllVendors(ctx context.Context) ([]Vendor, error) {
	rows, err := q.db.Query(ctx, getAllVendors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vendor{}
	for rows.Next() {
		var i Vendor
		if err := rows.Scan(&i.ID, &i.VendorName, &i.VendorLocation); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCatalog = `-- name: GetCatalog :many
SELECT 
p.product_name,
c.unit_price
FROM products p 
JOIN catalog c ON p.id = c.product_id
`

type GetCatalogRow struct {
	ProductName string `json:"product_name"`
	UnitPrice   int32  `json:"unit_price"`
}

func (q *Queries) GetCatalog(ctx context.Context) ([]GetCatalogRow, error) {
	rows, err := q.db.Query(ctx, getCatalog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCatalogRow{}
	for rows.Next() {
		var i GetCatalogRow
		if err := rows.Scan(&i.ProductName, &i.UnitPrice); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const loadTime = `-- name: LoadTime :one
SELECT NOW()
`

func (q *Queries) LoadTime(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRow(ctx, loadTime)
	var now interface{}
	err := row.Scan(&now)
	return now, err
}
