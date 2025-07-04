// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BProduct struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}

type BPurchase struct {
	ID         string           `json:"id"`
	MaterialID string           `json:"material_id"`
	Quantity   int32            `json:"quantity"`
	Price      pgtype.Numeric   `json:"price"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type BSale struct {
	ID           string           `json:"id"`
	ProductID    string           `json:"product_id"`
	Quantity     int32            `json:"quantity"`
	SellingPrice pgtype.Numeric   `json:"selling_price"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	CashierID    *string          `json:"cashier_id"`
}

type Catalog struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	UnitPrice int32  `json:"unit_price"`
}

type Category struct {
	ID                  string `json:"id"`
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

type Material struct {
	ID           string  `json:"id"`
	MaterialName string  `json:"material_name"`
	Unit         string  `json:"unit"`
	Description  *string `json:"description"`
}

type Product struct {
	ID          string `json:"id"`
	CategoryID  string `json:"category_id"`
	ProductName string `json:"product_name"`
}

type Purchase struct {
	ID         string           `json:"id"`
	ProductID  string           `json:"product_id"`
	TotalPrice int32            `json:"total_price"`
	Quantity   int32            `json:"quantity"`
	VendorID   *string          `json:"vendor_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	CashierID  *string          `json:"cashier_id"`
}

type Sale struct {
	ID        string           `json:"id"`
	ProductID string           `json:"product_id"`
	UnitPrice int32            `json:"unit_price"`
	Quantity  int32            `json:"quantity"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	CashierID *string          `json:"cashier_id"`
}

type Session struct {
	ID          string `json:"id"`
	Session     string `json:"session"`
	Description string `json:"description"`
}

type SessionMaterial struct {
	SessionID  string      `json:"session_id"`
	TeamID     string      `json:"team_id"`
	MaterialID string      `json:"material_id"`
	Date       pgtype.Date `json:"date"`
	Quantity   int32       `json:"quantity"`
}

type SessionProduct struct {
	SessionID string      `json:"session_id"`
	TeamID    string      `json:"team_id"`
	ProductID string      `json:"product_id"`
	Date      pgtype.Date `json:"date"`
	Quantity  int32       `json:"quantity"`
}

type Team struct {
	ID          string  `json:"id"`
	TeamName    string  `json:"team_name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}

type User struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Password    string  `json:"password"`
	Role        string  `json:"role"`
}

type Vendor struct {
	ID             string `json:"id"`
	VendorName     string `json:"vendor_name"`
	VendorLocation string `json:"vendor_location"`
	Description    string `json:"description"`
}
