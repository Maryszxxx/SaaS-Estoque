package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"saas-estoque/config/auth"
	"saas-estoque/database"
	_ "saas-estoque/docs"
	"saas-estoque/entity"
	"saas-estoque/handler"
	"saas-estoque/usercase"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db := database.ConnectPostgresRepository()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	r := SetupRouter(db)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Println("Servidor rodando na porta 8080")

		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("erro ao subir o servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Sinal de encerramento recebido")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("erro ao desligar servidor: %v", err)
	}

	log.Println("Servidor encerrado com sucesso")
}

func SetupRouter(db *sql.DB) *gin.Engine {

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

	// PRODUCTS

	r.POST(
		"/products",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
			entity.RoleEmployee,
		),
		h.Create,
	)

	r.GET(
		"/products",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
			entity.RoleEmployee,
		),
		h.Get,
	)

	r.GET(
		"/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
			entity.RoleEmployee,
		),
		h.GetById,
	)

	r.PUT(
		"/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
			entity.RoleEmployee,
		),
		h.Update,
	)

	r.PATCH(
		"/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
			entity.RoleEmployee,
		),
		h.Patch,
	)

	r.DELETE(
		"/products/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(
			entity.RoleAdmin,
		),
		h.Delete,
	)

	// USERS

	repoUser := database.NewPostgresUserRepository(db)

	serviceUser := usercase.NewUserService(repoUser)

	hUser := handler.NewUserHandler(serviceUser)

	r.POST("/users", hUser.Create)

	r.POST("/login", hUser.Login)

	r.POST("/refresh", hUser.Refresh)

	r.GET(
		"/users/email",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.FindByEmail,
	)

	r.GET(
		"/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.FindById,
	)

	r.DELETE(
		"/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.Delete,
	)

	r.PATCH(
		"/users/:id",
		auth.AuthMiddleware(),
		auth.RequiredRole(entity.RoleAdmin),
		hUser.Patch,
	)

	r.PATCH(
		"/users/password",
		auth.AuthMiddleware(),
		hUser.ChangePassword,
	)

	// SWAGGER

	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	return r
}
