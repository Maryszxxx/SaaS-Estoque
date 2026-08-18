package main

import (
	"net/http"
	"saas-estoque/config/auth"
	"saas-estoque/database"
	_ "saas-estoque/docs"
	"saas-estoque/entity"
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

	r.POST("/products",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin, entity.RoleEmployee),
		h.Create,
	)

	r.GET("/products",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin, entity.RoleEmployee),
		h.Get,
	)

	r.GET("/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin, entity.RoleEmployee),
		h.GetById,
	)

	r.PUT("/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin, entity.RoleEmployee),
		h.Update,
	)
	r.PATCH("/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin, entity.RoleEmployee),
		h.Patch,
	)

	r.DELETE("/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		h.Delete,
	)

	// implementando user

	repoUser := database.NewPostgresUserRepository(db)

	serviceUser := usercase.NewUserService(repoUser)

	hUser := handler.NewUserHandler(serviceUser)

	r.POST("/users", hUser.Create)

	r.POST("/login", hUser.Login)

	r.POST("/refresh", hUser.Refresh)

	r.GET("/users/email",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.FindByEmail,
	)

	r.GET("/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.FindById,
	)

	r.DELETE("/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.Delete,
	)

	r.PATCH("/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.Patch,
	)

	r.PATCH("/users/password",
		auth.AuthMiddleware(),
		hUser.ChangePassword,
	)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

}
