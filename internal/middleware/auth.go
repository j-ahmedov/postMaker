package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "request does not contain an access token",
				"status":  false,
			})
			c.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  false,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
