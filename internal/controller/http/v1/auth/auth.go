package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	auth_usecase "postMaker/internal/usecase/auth"
)

type Controller struct {
	useCase *auth_usecase.UseCase
}

func NewController(useCase *auth_usecase.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc Controller) GenerateToken(c *gin.Context) {
	var request auth.TokenRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		c.Abort()
		return
	}

	token, err := cc.useCase.GenerateToken(c, request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"token":   token,
	})

}
