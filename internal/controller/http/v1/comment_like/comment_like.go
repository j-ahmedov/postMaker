package comment_like

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	"postMaker/internal/service/comment_like"
	comment_usecase "postMaker/internal/usecase/comment"
	"strconv"
	"strings"
)

type Controller struct {
	useCase *comment_usecase.UseCase
}

func NewController(useCase *comment_usecase.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc Controller) GetCommentLikeList(c *gin.Context) {

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

	var filter comment_like.Filter
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

	list, count, err := cc.useCase.GetCommentLikeList(ctx, filter)
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

func (cc Controller) GetCommentLikeById(c *gin.Context) {
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

	detail, err := cc.useCase.GetCommentLikeById(ctx, id)
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

func (cc Controller) CreateCommentLike(c *gin.Context) {
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

	var data comment_like.Create

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	newData, err := cc.useCase.CreateCommentLike(ctx, data)
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

func (cc Controller) UpdateCommentLike(c *gin.Context) {
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

	var data comment_like.Update

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	newData, err := cc.useCase.UpdateCommentLike(ctx, data)
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

func (cc Controller) DeleteCommentLike(c *gin.Context) {
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

	err = cc.useCase.DeleteCommentLike(ctx, id)
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
