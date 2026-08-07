package handler

import (
	"fmt"
	"net/http"
	"saas-estoque/usercase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=4,max=90"`
	Description string  `json:"description" binding:"required,min=4"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int64   `json:"quantity" binding:"required,gt=0"`
	CategoryID  int64   `json:"category_id" binding:"required,gt=0"`
}
type PatchProductRequest struct {
	Name        *string  `json:"name" binding:"omitempty,min=4,max=90"`
	Description *string  `json:"description" binding:"omitempty,min=4"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Quantity    *int64   `json:"quantity" binding:"omitempty,gt=0"`
	CategoryID  *int64   `json:"category_id" binding:"omitempty,gt=0"`
}

type ProductHandler struct {
	service *usercase.ProductService
}

// POST
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

// PATCH
func (h *ProductHandler) Patch(c *gin.Context) {
	product := &PatchProductRequest{}

	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	err = h.service.Patch(
		idInt,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CategoryID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// PUT
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

// DELETE
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("erro no id")
		return
	}

	err = h.service.Delete(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})

}

// GetAll
func (h *ProductHandler) Get(c *gin.Context) {
	products, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"method": "GET"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetById
func (h *ProductHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("erro no id")
		return
	}

	productsById, err := h.service.FindByID(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productsById)

}
func NewProductHandler(service *usercase.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
