package api

import (
	"net/http"

	"github.com/Ndong21/SaaS-software/internal/stocks/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	querier repo.Querier
}

func NewStockHandler(querier repo.Querier) *StockHandler {
	return &StockHandler{
		querier: querier,
	}
}

func (h *StockHandler) WireHttpHandler() http.Handler {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, _ any) {
		c.String(http.StatusInternalServerError, "Internal Server Error: panic")
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// Add CORS middleware
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "https://saas-z-ax4i.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	//stocks module
	r.GET("/demo", h.handleGetDBTime)
	r.POST("/product", h.handleCreateProduct)
	r.POST("/category", h.handleCreateCategory)
	r.POST("/purchase", h.handleCreatePurchase)
	r.POST("/sale", h.handleCreateSale)
	r.POST("/vendor", h.handleCreateVendor)
	r.POST("/catalog", h.handleCreateCatalog)
	r.GET("/product", h.handleGetProducts)
	r.GET("/category", h.handleGetCategories)
	r.GET("/catalog", h.handleGetCatalog)
	r.GET("/vendor", h.handleGetVendors)
	r.GET("/purchase", h.handleGetPurchases)
	r.GET("/sale", h.handleGetSales)
	r.DELETE("/catalog/:id", h.handleDeleteCatalog)
	//delete sales
	r.DELETE("/sale/:id", h.handleDeleteSale)
	//update the sales table
	r.PATCH("/sale/update", h.handleUpdateSales)

	//delete purchase
	r.DELETE("/purchase/:id", h.handleDeletePurchase)
	//update the purchases table
	r.PATCH("/purchase/update", h.handleUpdatePurchase)

	//delete product
	r.DELETE("/product/:id", h.handleDeleteProduct)
	//update the product table
	r.PATCH("/product/update", h.handleUpdateProduct)

	//delete category
	r.DELETE("/category/:id", h.handleDeleteCategory)
	//update the categories table
	r.PATCH("/category/update", h.handleUpdateCategory)

	r.GET("/reports/total-sales", h.handleGetTotalSales)
	r.GET("/reports/transactions", h.handleGetTotalTransactions)
	r.GET("/reports/top-products", h.handleGetTopProducts)

	//blocks module
	r.POST("/blocks/material", h.handleCreateMaterial)
	r.POST("/blocks/product", h.handleCreateBlockProduct)
	r.POST("/blocks/purchase", h.handleCreateMaterialPurchase)
	r.POST("/blocks/sale", h.handleCreateBlockSale)
	r.POST("/blocks/team", h.handleCreateTeam)
	r.POST("/blocks/session", h.handleCreateSession)
	r.POST("/blocks/session/material", h.handleCreateSessionMaterial)
	r.POST("/blocks/session/product", h.handleCreateSessionProduct)
	r.GET("/blocks/product", h.handleGetBlockProduct)
	r.GET("/blocks/material", h.handleGetBlockMaterial)
	r.GET("/blocks/purchase", h.handleGetBlockMaterialPurchases)
	r.GET("/blocks/team", h.handleGetTeam)
	r.GET("/blocks/session", h.handleGetSessions)
	r.GET("/blocks/sale", h.handleGetBlockSale)
	r.GET("/blocks/session/material", h.handleGetSessionMaterial)
	r.GET("/blocks/session/product", h.handleGetSessionProduct)

	r.DELETE("/blocks/material/:id", h.handleDeleteMaterial)
	r.PATCH("/blocks/material/update", h.handleUpdateMaterial)

	r.DELETE("/blocks/purchase/:id", h.handleDeleteBlockPurchase)
	r.PATCH("/blocks/purchase/update", h.handleUpdateBlockPurchase)

	r.DELETE("/blocks/product/:id", h.handleDeleteBlockProduct)
	r.PATCH("/blocks/product/update", h.handleUpdateBlockProduct)

	r.DELETE("/blocks/team/:id", h.handleDeleteTeam)
	r.PATCH("/blocks/team/update", h.handleUpdateTeam)

	r.DELETE("/blocks/session/:id", h.handleDeleteSession)
	r.PATCH("/blocks/session/update", h.handleUpdateSession)

	r.DELETE("/blocks/session/material/:id", h.handleDeleteSessionMaterial)
	r.PATCH("/blocks/session/material/update", h.handleUpdateSessionMaterial)

	r.DELETE("/blocks/session/product/:id", h.handleDeleteSessionProduct)
	r.PATCH("/blocks/session/product/update", h.handleUpdateSessionProduct)

	r.DELETE("/blocks/sale/:id", h.handleDeleteBlockSale)
	r.PATCH("/blocks/sale/update", h.handleUpdateBlockSale)

	// // users
	// r.POST("/api/auth/user", h.handleCreateUser)
	// r.POST("/api/auth/login", h.handleLogin)

	return r
}

// testing handler
func (h *StockHandler) handleGetDBTime(c *gin.Context) {
	time, err := h.querier.LoadTime(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, time)
}

func (h *StockHandler) handleCreateCategory(c *gin.Context) {
	var req repo.CreateCategoryParams

	// Bind the request body to the struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Manual validation for category name
	if req.CategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Category name is required.",
		})
		return
	}

	// Create the category in the database
	_, err = h.querier.CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create category: " + err.Error(),
		})
		return
	}

	// Respond with success message only
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Category created successfully.",
	})
}

func (h *StockHandler) handleCreateProduct(c *gin.Context) {
	var req repo.CreateProductParams

	// Bind JSON request body to the struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Optional: add your own required field checks here
	if req.ProductName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Product name is required.",
		})
		return
	}

	// Create product
	_, err = h.querier.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create product: " + err.Error(),
		})
		return
	}

	// Respond with success only
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Product created successfully.",
	})
}

func (h *StockHandler) handleCreatePurchase(c *gin.Context) {
	var req repo.CreatePurchaseParams

	// Bind the request body to the struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Optional: Add field-level validation
	if req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Quantity must be provided and greater than zero.",
		})
		return
	}

	// Attempt to create the purchase
	_, err = h.querier.CreatePurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create purchase: " + err.Error(),
		})
		return
	}

	// Respond with success message only
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Purchase created successfully.",
	})
}

func (h *StockHandler) handleCreateSale(c *gin.Context) {
	var req repo.CreateSaleParams

	// Bind JSON request to struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Optional: Validate required fields
	if req.ProductID == "" || req.Quantity <= 0 || req.UnitPrice <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Product ID, quantity, and unit price must be valid and greater than 0.",
		})
		return
	}

	// Attempt to create the sale
	_, err = h.querier.CreateSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create sale: " + err.Error(),
		})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Sale recorded successfully.",
	})
}

func (h *StockHandler) handleCreateVendor(c *gin.Context) {
	var req repo.CreateVendorParams

	// Bind the request JSON to the struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Optional: Validate required fields
	if req.VendorName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Vendor name is required.",
		})
		return
	}

	// Attempt to create the vendor
	_, err = h.querier.CreateVendor(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create vendor: " + err.Error(),
		})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Vendor created successfully.",
	})
}

func (h *StockHandler) handleCreateCatalog(c *gin.Context) {
	var req repo.CreateCatalogParams

	// Bind JSON to struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Create catalog item
	_, err = h.querier.CreateCatalog(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create catalog item: " + err.Error(),
		})
		return
	}

	// Success response without returning the object
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Catalog item created successfully.",
	})
}

func (h *StockHandler) handleGetProducts(c *gin.Context) {
	products, err := h.querier.GetAllProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *StockHandler) handleGetVendors(c *gin.Context) {
	vendors, err := h.querier.GetAllVendors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vendors": vendors,
	})
}

func (h *StockHandler) handleGetCategories(c *gin.Context) {
	categories, err := h.querier.GetAllCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *StockHandler) handleGetCatalog(c *gin.Context) {
	catalog, err := h.querier.GetCatalog(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"catalog": catalog,
	})
}

func (h *StockHandler) handleGetPurchases(c *gin.Context) {
	purchases, err := h.querier.GetAllPurchases(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"purchases": purchases,
	})
}

func (h *StockHandler) handleGetSales(c *gin.Context) {
	sales, err := h.querier.GetAllSales(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sales": sales,
	})
}

func (h *StockHandler) handleGetTotalSales(c *gin.Context) {
	total_sales, err := h.querier.TotalSales(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_sales": total_sales,
	})
}

func (h *StockHandler) handleGetTotalTransactions(c *gin.Context) {
	total_transactions, err := h.querier.CountSalesTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_transactions": total_transactions,
	})
}

func (h *StockHandler) handleGetTopProducts(c *gin.Context) {
	top_products, err := h.querier.Top5BestSellingProductsByRevenue(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"top_selling_products": top_products,
	})
}

func (h *StockHandler) handleDeleteCatalog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteCatalog(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "catalog item successfully deleted",
		"status":  "success",
	})
}

func (h *StockHandler) handleDeleteSale(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteSale(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sale successfully deleted",
		"status":  "success",
	})
}

func (h *StockHandler) handleDeletePurchase(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeletePurchase(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "purchase successfully deleted",
		"status":  "success",
	})
}

func (h *StockHandler) handleDeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.Deleteproduct(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product successfully deleted",
		"status":  "success",
	})
}

func (h *StockHandler) handleDeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteCategory(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category successfully deleted",
		"status":  "success",
	})
}

// endpoints for block production module
func (h *StockHandler) handleCreateMaterial(c *gin.Context) {
	var req repo.CreateMaterialParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateMaterial(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create material",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Material created successfully",
	})
}

func (h *StockHandler) handleCreateBlockProduct(c *gin.Context) {
	var req repo.CreateBlocksProductParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateBlocksProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create block type",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Block type created successfully",
	})
}

func (h *StockHandler) handleCreateMaterialPurchase(c *gin.Context) {
	var req repo.CreateMaterialPurchaseParams

	// Try to bind the request body to the struct
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	// Attempt to create the material purchase in the database
	_, err = h.querier.CreateMaterialPurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create material purchase: " + err.Error(),
		})
		return
	}

	// Return a success message without the created object
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Material purchase created successfully.",
	})
}

func (h *StockHandler) handleCreateBlockSale(c *gin.Context) {
	var req repo.CreateBlockSaleParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateBlockSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to record block sale",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Block sale recorded successfully",
	})
}

func (h *StockHandler) handleCreateTeam(c *gin.Context) {
	var req repo.CreateTeamParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateTeam(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create team",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Team created successfully",
	})
}

func (h *StockHandler) handleCreateSession(c *gin.Context) {
	var req repo.CreateSessionParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateSession(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create session",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Session created successfully",
	})
}

func (h *StockHandler) handleCreateSessionMaterial(c *gin.Context) {
	var req repo.CreateSessionMaterialsParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateSessionMaterials(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to record session material",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Session material recorded successfully",
	})
}

func (h *StockHandler) handleCreateSessionProduct(c *gin.Context) {
	var req repo.CreateSessionProductsParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	_, err := h.querier.CreateSessionProducts(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to record session product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Session product recorded successfully",
	})
}

func (h *StockHandler) handleGetBlockProduct(c *gin.Context) {
	products, err := h.querier.GetBlocksProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch block products",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"message":  "Block products retrieved successfully",
		"products": products,
	})
}

func (h *StockHandler) handleGetBlockMaterial(c *gin.Context) {
	materials, err := h.querier.GetMaterials(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch materials",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"message":   "materials retrieved successfully",
		"materials": materials,
	})
}

func (h *StockHandler) handleGetBlockMaterialPurchases(c *gin.Context) {
	purchases, err := h.querier.GetMaterialPurchases(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch material purchases",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"message":   "material purchases retrieved successfully",
		"purchases": purchases,
	})
}

func (h *StockHandler) handleGetTeam(c *gin.Context) {
	teams, err := h.querier.GetTeams(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch teams",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "teams retrieved successfully",
		"teams":   teams,
	})
}

func (h *StockHandler) handleGetSessions(c *gin.Context) {
	sessions, err := h.querier.GetSessions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch sessions",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"message":  "sessions retrieved successfully",
		"sessions": sessions,
	})
}

func (h *StockHandler) handleGetBlockSale(c *gin.Context) {
	sales, err := h.querier.GetBlockSales(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch sales",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "sales retrieved successfully",
		"sales":   sales,
	})
}

func (h *StockHandler) handleGetSessionMaterial(c *gin.Context) {
	session_materials, err := h.querier.GetSessionMaterials(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch session materials",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            "success",
		"message":           "session materials retrieved successfully",
		"session_materials": session_materials,
	})
}

func (h *StockHandler) handleGetSessionProduct(c *gin.Context) {
	session_products, err := h.querier.GetSessionProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch session products",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":           "success",
		"message":          "session products retrieved successfully",
		"session_products": session_products,
	})
}

// Endpoint to update sales info
func (h *StockHandler) handleUpdateSales(c *gin.Context) {
	var req repo.UpdateSalesParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requst body"})
	}

	sales, err := h.querier.UpdateSales(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating sales data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"sales":  sales,
	})
}

// Endpoint to update purchase info
func (h *StockHandler) handleUpdatePurchase(c *gin.Context) {
	var req repo.UpdatePurchaseParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requst body"})
	}

	purchase, err := h.querier.UpdatePurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating purchase data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"purchase": purchase,
	})
}

// Endpoint to update product info
func (h *StockHandler) handleUpdateProduct(c *gin.Context) {
	var req repo.UpdateProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requst body"})
	}

	product, err := h.querier.UpdateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating product data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"product": product,
	})
}

// Endpoint to update category info
func (h *StockHandler) handleUpdateCategory(c *gin.Context) {
	var req repo.UpdateCategoryParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requst body"})
	}

	category, err := h.querier.UpdateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating category data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"category": category,
	})
}

// update and delete on block module
// Update material
func (h *StockHandler) handleUpdateMaterial(c *gin.Context) {
	var req repo.UpdateMaterialParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	material, err := h.querier.UpdateMaterial(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating material": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "material": material})
}

// Delete material
func (h *StockHandler) handleDeleteMaterial(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteMaterial(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting material": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "material deleted"})
}

func (h *StockHandler) handleUpdateBlockPurchase(c *gin.Context) {
	var req repo.UpdateBlockPurchaseParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	purchase, err := h.querier.UpdateBlockPurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating purchase": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "purchase": purchase})
}

func (h *StockHandler) handleDeleteBlockPurchase(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteBlockPurchase(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting purchase": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "purchase deleted"})
}

func (h *StockHandler) handleUpdateBlockProduct(c *gin.Context) {
	var req repo.UpdateBlockProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	product, err := h.querier.UpdateBlockProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating product": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "product": product})
}

func (h *StockHandler) handleDeleteBlockProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteBlockProduct(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting product": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "product deleted"})
}

func (h *StockHandler) handleUpdateTeam(c *gin.Context) {
	var req repo.UpdateTeamParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	team, err := h.querier.UpdateTeam(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating team": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "team": team})
}

func (h *StockHandler) handleDeleteTeam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteTeam(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting team": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "team deleted"})
}

func (h *StockHandler) handleUpdateSession(c *gin.Context) {
	var req repo.UpdateSessionParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	session, err := h.querier.UpdateSession(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating session": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "session": session})
}

func (h *StockHandler) handleDeleteSession(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteSession(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting session": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "session deleted"})
}

func (h *StockHandler) handleUpdateSessionMaterial(c *gin.Context) {
	var req repo.UpdateSessionMaterialParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	record, err := h.querier.UpdateSessionMaterial(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating session material": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "session_material": record})
}

func (h *StockHandler) handleDeleteSessionMaterial(c *gin.Context) {
	var req repo.DeleteSessionMaterialParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.querier.DeleteSessionMaterial(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting session material": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "session material deleted"})
}

func (h *StockHandler) handleUpdateSessionProduct(c *gin.Context) {
	var req repo.UpdateSessionProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	record, err := h.querier.UpdateSessionProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating session product": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "session_product": record})
}

func (h *StockHandler) handleDeleteSessionProduct(c *gin.Context) {
	var req repo.DeleteSessionProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.querier.DeleteSessionProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting session product": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "session product deleted"})
}

func (h *StockHandler) handleUpdateBlockSale(c *gin.Context) {
	var req repo.UpdateBlockSaleParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	sale, err := h.querier.UpdateBlockSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating sale": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "sale": sale})
}

func (h *StockHandler) handleDeleteBlockSale(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.querier.DeleteBlockSale(c, id) // typo in your SQL name, should be DeleteBlockSale
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error deleting sale": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "sale deleted"})
}

// // users
// // sign-up
// func (h *StockHandler) handleCreateUser(c *gin.Context) {
// 	var req repo.CreateUserParams
// 	err := c.ShouldBindBodyWithJSON(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Hash the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// replace the raw password with its hashed version
// 	req.Password = string(hashedPassword)

// 	user, err := h.querier.CreateUser(c, req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// // login
// func (h *StockHandler) handleLogin(c *gin.Context) {
// 	// var req repo.CreateUserParams
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	err := c.ShouldBindBodyWithJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := h.querier.SelectRequestedUser(c, body.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if user.ID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
// 	}

// 	// compare provided password with stored hassed password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
// 		return
// 	}

// 	//generate token
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte("dsjkfliojfsldk"))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
// 		return
// 	}

// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
// 	// If valid, return user info (or generate JWT/token if needed)
// 	c.JSON(http.StatusOK, gin.H{})

// }
