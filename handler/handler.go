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

type UserHandler struct {
	userService *usercase.UserService
}

type UserRequest struct {
	Name         string `json:"name" binding:"required,min=4,max=90"`
	Email        string `json:"email" binding:"required,email"`
	PasswordHash string `json:"password_hash" binding:"required"`
	Role         string `json:"role" binding:"required"`
}

// Create godoc
// @Summary Cria um novo produto
// @Description Cria um produto no estoque
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Dados do produto"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]

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

// Patch godoc
// @Summary Atualiza parcialmente um produto
// @Description Atualiza somente os campos enviados
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body PatchProductRequest true "Campos para atualização"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [patch]
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

// Update godoc
// @Summary Atualiza um produto
// @Description Atualiza todos os campos de um produto
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body CreateProductRequest true "Dados do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
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

// Delete godoc
// @Summary Remove um produto
// @Description Remove um produto pelo ID
// @Tags products
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
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

// Get godoc
// @Summary Lista todos os produtos
// @Description Retorna todos os produtos cadastrados
// @Tags products
// @Produce json
// @Success 200 {array} entity.Product
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [get]

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

// GetById godoc
// @Summary Busca produto por ID
// @Description Retorna um produto específico
// @Tags products
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} entity.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
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

	return &ProductHandler{
		service: service,
	}

}

// implementação de login usuario

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := &UserRequest{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.userService.CreateUser(
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
