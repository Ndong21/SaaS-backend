package api

import (
	"net/http"
	"time"

	"github.com/Ndong21/SaaS-software/internal/stocks/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
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
	r.GET("/catalog", h.handleGetCatalog)

	//blocks module
	r.POST("/api/blocks/material", h.handleCreateMaterial)
	r.POST("/api/blocks/product", h.handleCreateBlockProduct)
	r.POST("/api/blocks/purchase", h.handleCreateMaterialPurchase)
	r.POST("/api/blocks/sale", h.handleCreateBlockSale)
	r.POST("/api/blocks/team", h.handleCreateTeam)
	r.POST("/api/blocks/session", h.handleCreateSession)
	r.POST("/api/blocks/session/material", h.handleCreateSessionMaterial)
	r.POST("/api/blocks/session/product", h.handleCreateSessionProduct)
	// r.GET("/product", h.handleGetProducts)
	// r.GET("/catalog", h.handleGetCatalog)

	// users
	r.POST("/api/auth/user", h.handleCreateUser)
	r.POST("/api/auth/login", h.handleLogin)

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
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.CategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category name is required"})
		return
	}

	category, err := h.querier.CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *StockHandler) handleCreateProduct(c *gin.Context) {
	var req repo.CreateProductParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.querier.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *StockHandler) handleCreatePurchase(c *gin.Context) {
	var req repo.CreatePurchaseParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase, err := h.querier.CreatePurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func (h *StockHandler) handleCreateSale(c *gin.Context) {
	var req repo.CreateSaleParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, err := h.querier.CreateSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sale)
}

func (h *StockHandler) handleCreateVendor(c *gin.Context) {
	var req repo.CreateVendorParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor, err := h.querier.CreateVendor(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func (h *StockHandler) handleCreateCatalog(c *gin.Context) {
	var req repo.CreateCatalogParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.querier.CreateCatalog(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
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

// endpoints for block production module
func (h *StockHandler) handleCreateMaterial(c *gin.Context) {
	var req repo.CreateMaterialParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := h.querier.CreateMaterial(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *StockHandler) handleCreateBlockProduct(c *gin.Context) {
	var req repo.CreateBlocksProductParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.querier.CreateBlocksProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *StockHandler) handleCreateMaterialPurchase(c *gin.Context) {
	var req repo.CreateMaterialPurchaseParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase, err := h.querier.CreateMaterialPurchase(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func (h *StockHandler) handleCreateBlockSale(c *gin.Context) {
	var req repo.CreateBlockSaleParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, err := h.querier.CreateBlockSale(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sale)
}

func (h *StockHandler) handleCreateTeam(c *gin.Context) {
	var req repo.CreateTeamParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.querier.CreateTeam(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *StockHandler) handleCreateSession(c *gin.Context) {
	var req repo.CreateSessionParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.querier.CreateSession(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *StockHandler) handleCreateSessionMaterial(c *gin.Context) {
	var req repo.CreateSessionMaterialsParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionMaterial, err := h.querier.CreateSessionMaterials(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessionMaterial)
}

func (h *StockHandler) handleCreateSessionProduct(c *gin.Context) {
	var req repo.CreateSessionProductsParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionProduct, err := h.querier.CreateSessionProducts(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessionProduct)
}

// users
// sign-up
func (h *StockHandler) handleCreateUser(c *gin.Context) {
	var req repo.CreateUserParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// replace the raw password with its hashed version
	req.Password = string(hashedPassword)

	user, err := h.querier.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// login
func (h *StockHandler) handleLogin(c *gin.Context) {
	// var req repo.CreateUserParams
	var body struct {
		Email    string
		Password string
	}

	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.querier.SelectRequestedUser(c, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
	}

	// compare provided password with stored hassed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	//generate token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("dsjkfliojfsldk"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	// If valid, return user info (or generate JWT/token if needed)
	c.JSON(http.StatusOK, gin.H{})

}
