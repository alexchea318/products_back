package private

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/database"
	"server/internal/models"
)

type ProductHandlersRoutes interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProductHandlers struct {
	ProductsTable database.ProductsTable
}

// CreateProduct creates a new product.
func (p ProductHandlers) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err := p.ProductsTable.Insert(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProducts gets all products.
func (p ProductHandlers) GetProducts(c *gin.Context) {

}

// GetProduct gets a product by its ID.
func (p ProductHandlers) GetProduct(c *gin.Context) {

}

// UpdateProduct updates a product.
func (p ProductHandlers) UpdateProduct(c *gin.Context) {

}

// DeleteProduct deletes a product.
func (p ProductHandlers) DeleteProduct(c *gin.Context) {

}
