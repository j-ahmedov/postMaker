package comment

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	"postMaker/internal/service/comment"
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

func (cc Controller) GetCommentList(c *gin.Context) {

	//tokenString := c.GetHeader("Authorization")
	//tokenString = strings.Split(tokenString, "Bearer ")[1]
	//
	//_, err := auth.ParseToken(tokenString)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": err.Error(),
	//		"status":  false,
	//	})
	//	return
	//}

	var filter comment.Filter
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

	list, count, err := cc.useCase.GetCommentList(ctx, filter)
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

func (cc Controller) GetCommentListByPostId(c *gin.Context) {
	var filter comment.Filter
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

	idParams := c.Param("post_id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	list, count, err := cc.useCase.GetCommentListByPostId(ctx, filter, id)
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

func (cc Controller) GetCommentById(c *gin.Context) {

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

	detail, err := cc.useCase.GetCommentById(ctx, id)
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

func (cc Controller) GetCommentByPostId(c *gin.Context) {

	idParams := c.Param("post_id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	detail, err := cc.useCase.GetCommentByPostId(ctx, id)
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

func (cc Controller) CreateComment(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	tokenData, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	var data comment.Create

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	data.UserId = tokenData.Id

	newData, err := cc.useCase.CreateComment(ctx, data)
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

func (cc Controller) UpdateComment(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	tokenData, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	var data comment.Update

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	data.UserId = tokenData.Id

	newData, err := cc.useCase.UpdateComment(ctx, data)
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

func (cc Controller) DeleteComment(c *gin.Context) {

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

	err = cc.useCase.DeleteComment(ctx, id)
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
