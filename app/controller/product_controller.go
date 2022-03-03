package controller

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

type ProductController interface {
	PostProduct(*gin.Context)
}

func NewProductController(service service.ProductService) ProductController {
	return productController{
		productService: service,
	}
}

func (controller productController) PostProduct(ctx *gin.Context) {
	var inputProduct entity.Product
	if err := ctx.ShouldBind(&inputProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := controller.productService.Save(inputProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product Successfully added",
		"product": product,
	})
}

func (controller productController) GetAllProduct(ctx *gin.Context) {
	products, err := controller.productService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"product": products,
	})
}
