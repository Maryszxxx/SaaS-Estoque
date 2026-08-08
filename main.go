package main

import (
	"saas-estoque/database"
	"saas-estoque/handler"
	"saas-estoque/usercase"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectPostgresProductRepository()
	defer db.Close()

	repo := database.NewPostgresProductRepository(db)

	service := usercase.NewProductService(repo)

	h := handler.NewProductHandler(service)

	r := gin.Default()

	r.POST("/products", h.Create)
	r.GET("/products", h.Get)         //findAll
	r.GET("/products/:id", h.GetById) //findById
	r.PUT("/products/:id", h.Update)  //put
	r.PATCH("/products/:id", h.Patch) //patch
	r.DELETE("/products/:id", h.Delete)

	r.Run(":8080")

}
