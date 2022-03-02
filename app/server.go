package app

import (
	"kasir-api-gin/app/controller"
	"kasir-api-gin/config"
	"kasir-api-gin/helper"
	"kasir-api-gin/repository"
	"kasir-api-gin/service"
	"log"
	"net/http"
	"strings"

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

	// migrate database
	if err := userRepository.Migrate(); err != nil {
		log.Panic(err.Error())
	}

	// service
	authService := service.NewAuthService(userRepository, hash, token)
	// controller
	authController := controller.NewAuthController(authService)
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Kasir App With gin framework",
		})
	})

	public := server.Group("/api/v1")

	auth := public.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.POST("/check", func(c *gin.Context) {
		headerAuth := c.GetHeader("Authorization")
		if headerAuth == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Token tidak valid",
			})
			return
		}
		tokenString := strings.Split(headerAuth, " ")[1]
		id, err := token.Decode(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	return server
}
