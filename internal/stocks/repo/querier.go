// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package repo

import (
	"context"
)

type Querier interface {
	CountSalesTransactions(ctx context.Context) (int64, error)
	CreateBlockSale(ctx context.Context, arg CreateBlockSaleParams) (BSale, error)
	CreateBlocksProduct(ctx context.Context, arg CreateBlocksProductParams) (BProduct, error)
	CreateCatalog(ctx context.Context, arg CreateCatalogParams) (Catalog, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateMaterial(ctx context.Context, arg CreateMaterialParams) (Material, error)
	CreateMaterialPurchase(ctx context.Context, arg CreateMaterialPurchaseParams) (BPurchase, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (Purchase, error)
	CreateSale(ctx context.Context, arg CreateSaleParams) (Sale, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSessionMaterials(ctx context.Context, arg CreateSessionMaterialsParams) (SessionMaterial, error)
	CreateSessionProducts(ctx context.Context, arg CreateSessionProductsParams) (SessionProduct, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVendor(ctx context.Context, arg CreateVendorParams) (Vendor, error)
	DeleteCatalog(ctx context.Context, id string) error
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetAllProducts(ctx context.Context) ([]GetAllProductsRow, error)
	GetAllPurchases(ctx context.Context) ([]GetAllPurchasesRow, error)
	GetAllSales(ctx context.Context) ([]GetAllSalesRow, error)
	GetAllVendors(ctx context.Context) ([]Vendor, error)
	GetBlockSales(ctx context.Context) ([]GetBlockSalesRow, error)
	GetBlocksProducts(ctx context.Context) ([]GetBlocksProductsRow, error)
	GetCatalog(ctx context.Context) ([]GetCatalogRow, error)
	GetMaterialPurchases(ctx context.Context) ([]GetMaterialPurchasesRow, error)
	GetMaterials(ctx context.Context) ([]Material, error)
	GetSessionMaterials(ctx context.Context) ([]GetSessionMaterialsRow, error)
	GetSessionProducts(ctx context.Context) ([]GetSessionProductsRow, error)
	GetSessions(ctx context.Context) ([]Session, error)
	GetTeams(ctx context.Context) ([]Team, error)
	LoadTime(ctx context.Context) (interface{}, error)
	// -- name: LogIn :one
	// INSERT INTO "users" (email, password)
	// VALUES ($1, $2)
	// RETURNING *;
	SelectRequestedUser(ctx context.Context, email string) (SelectRequestedUserRow, error)
	Top5BestSellingProductsByRevenue(ctx context.Context) ([]Top5BestSellingProductsByRevenueRow, error)
	TotalSales(ctx context.Context) (int64, error)
	UpdateBlockSale(ctx context.Context, arg UpdateBlockSaleParams) (BSale, error)
}

var _ Querier = (*Queries)(nil)
