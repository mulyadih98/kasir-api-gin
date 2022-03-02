package controller

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

func NewAuthController(service service.AuthService) AuthController {
	return authController{
		authService: service,
	}
}

func (controller authController) Register(ctx *gin.Context) {
	var inputRegister entity.User
	if err := ctx.ShouldBind(&inputRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := controller.authService.Register(inputRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "User SuccessFully Registered",
		"User ID": id,
	})
}

func (controller authController) Login(ctx *gin.Context) {
	var inputLogin entity.LoginInput
	if err := ctx.ShouldBind(&inputLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := controller.authService.Login(inputLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Loggin Successful",
		"token":   token,
	})
}
