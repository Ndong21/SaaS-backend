package api

import (
	"net/http"

	"github.com/Ndong21/SaaS-software/internal/stocks/repo"
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
