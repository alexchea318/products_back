package private

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/database"
	"server/internal/models"
	"strconv"
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
	var product models.NewProduct

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

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: product,
	})
}

// GetProducts gets all products.
func (p ProductHandlers) GetProducts(c *gin.Context) {
	// Get the search query from the query string.
	search := c.Query("search")
	limit, err := strconv.ParseUint(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.ParseUint(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}

	products, err := p.ProductsTable.Get(search, limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: products,
	})

}

// GetProduct gets a product by its ID.
func (p ProductHandlers) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	product, err := p.ProductsTable.GetOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: product,
	})
}

// UpdateProduct updates a product.
func (p ProductHandlers) UpdateProduct(c *gin.Context) {
	// Get the product ID from the URL.
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var product models.NewProduct

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = p.ProductsTable.Update(product, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: product,
	})
}

// DeleteProduct deletes a product.
func (p ProductHandlers) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = p.ProductsTable.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: id,
	})
}
