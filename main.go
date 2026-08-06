package main

import (
	"saas-estoque/handler"
	"saas-estoque/repository/memory"
	"saas-estoque/usercase"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := memory.NewProductMemoryRepository()
	service := usercase.NewProductService(repo)

	h := handler.NewProductHandler(service)

	r := gin.Default()

	r.POST("/products", h.Create)
	r.GET("/products", h.Get)
	r.GET("/products/:id", h.Get)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)

	r.Run(":8080")

}
