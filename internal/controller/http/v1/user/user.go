package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	"postMaker/internal/service/user"
	user_usecase "postMaker/internal/usecase/user"
	"strconv"
	"strings"
)

type Controller struct {
	useCase *user_usecase.UseCase
}

func NewController(useCase *user_usecase.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc Controller) GetUserList(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	_, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	var filter user.Filter
	query := c.Request.URL.Query()

	limitQ := query["limit"]
	if len(limitQ) > 0 {
		queryInt, err := strconv.Atoi(limitQ[0])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Limit must be a number",
				"status":  false,
			})
			return
		}

		filter.Limit = &queryInt
	}

	offsetQ := query["offset"]
	if len(offsetQ) > 0 {
		queryInt, err := strconv.Atoi(offsetQ[0])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Offset must be number!",
				"status":  false,
			})
			return
		}

		filter.Offset = &queryInt
	}

	ctx := context.Background()

	list, count, err := cc.useCase.GetUserList(ctx, filter)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data": map[string]interface{}{
			"results": list,
			"count":   count,
		},
	})

}

func (cc Controller) GetUserById(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	_, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	idParams := c.Param("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	detail, err := cc.useCase.GetUserById(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data":    detail,
	})

}

func (cc Controller) CreateUser(c *gin.Context) {
	var data user.CreateForm

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	newData, err := cc.useCase.CreateUser(ctx, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data":    newData,
	})
}

func (cc Controller) UpdateUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	_, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	var data user.UpdateForm

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	newData, err := cc.useCase.UpdateUser(ctx, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data":    newData,
	})
}

func (cc Controller) DeleteUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	_, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	idParams := c.Param("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	err = cc.useCase.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
	})

}
