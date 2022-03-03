package app

import (
	"kasir-api-gin/app/controller"
	"kasir-api-gin/app/middleware"
	"kasir-api-gin/config"
	"kasir-api-gin/helper"
	"kasir-api-gin/repository"
	"kasir-api-gin/service"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	server := gin.Default()
	db, err := config.MysqlConnection()
	if err != nil {
		log.Fatal("Terjadi Kesalahan Pada Koneksi Database")
	}

	//helper
	hash := helper.NewPasswordHash()
	token := helper.NewTokenJWT()

	// repository
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepositoryGorm(db)

	// service
	authService := service.NewAuthService(userRepository, hash, token)
	productService := service.NewProductService(productRepository)

	// controller
	authController := controller.NewAuthController(authService)
	productController := controller.NewProductController(productService)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Kasir App With gin framework",
		})
	})
	// public route
	public := server.Group("/api/v1")
	// auth route
	auth := public.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	// safe route for user authenticated
	authRoute := server.Group("/api/v1")
	authRoute.Use(middleware.AuthJwt())
	// route for product
	productRoute := authRoute.Group("/products")
	productRoute.POST("/", productController.PostProduct)
	productRoute.GET("/", productController.GetAllProduct)
	productRoute.GET("/:product_id", productController.GetByIdProduct)
	productRoute.PUT("/:product_id", productController.PutProduct)

	return server
}
