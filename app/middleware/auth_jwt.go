package middleware

import (
	"kasir-api-gin/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := helper.NewTokenJWT()
		headerAuth := c.GetHeader("Authorization")
		if headerAuth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}
		tokenString := strings.Split(headerAuth, " ")[1]
		id, err := token.Decode(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}
