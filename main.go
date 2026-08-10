package main

import (
	"net/http"
	"saas-estoque/database"
	_ "saas-estoque/docs"
	"saas-estoque/handler"
	"saas-estoque/usercase"

	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {
	db := database.ConnectPostgresRepository()
	defer db.Close()

	repo := database.NewPostgresProductRepository(db)

	service := usercase.NewProductService(repo)

	h := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":   "unhealthy",
				"database": "down",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"database": "up",
		})
	})

	r.POST("/products", h.Create)
	r.GET("/products", h.Get)         //findAll
	r.GET("/products/:id", h.GetById) //findById
	r.PUT("/products/:id", h.Update)  //put
	r.PATCH("/products/:id", h.Patch) //patch
	r.DELETE("/products/:id", h.Delete)

	// implementando user

	repoUser := database.NewPostgresUserRepository(db)

	serviceUser := usercase.NewUserService(repoUser)

	hUser := handler.NewUserHandler(serviceUser)

	r.POST("/users", hUser.Create)
	r.GET("/users", hUser.FindByEmail)
	r.GET("/users/:id", hUser.FindById)
	r.DELETE("/users/:id", hUser.Delete)
	r.PATCH("/users/:id", hUser.Patch)
	r.PATCH("/users/:id/:password", hUser.ChangePassword)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

}
