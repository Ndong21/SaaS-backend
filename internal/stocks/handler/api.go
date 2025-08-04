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
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "https://saas-z-t6m3.vercel.app/"},
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

	r.PUT("/blocks/sale/:id", h.handleUpdateBlockSale)

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

func (h *StockHandler) handleUpdateBlockSale(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req repo.UpdateBlockSaleParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	req.ID = id

	_, err := h.querier.UpdateBlockSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update block sale",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Block sale successfully updated",
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
