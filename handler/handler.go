package handler

import (
	"net/http"
	"saas-estoque/usercase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	CategoryID  int64   `json:"category_id"`
}

type ProductHandler struct {
	service *usercase.ProductService
}

func (h *ProductHandler) Create(c *gin.Context) {
	product := &CreateProductRequest{}
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Create(
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (h *ProductHandler) Update(c *gin.Context) {
	product := &CreateProductRequest{}

	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(
		idInt,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CategoryID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func (h *ProductHandler) Get(c *gin.Context) {
	products, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"method": "GET"})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"method": "GET"})
		return
	}
	product := products[0]
	c.JSON(http.StatusOK, product)
}

func NewProductHandler(service *usercase.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
