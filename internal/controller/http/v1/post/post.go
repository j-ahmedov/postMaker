package post

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"postMaker/internal/service/auth"
	"postMaker/internal/service/post"
	"postMaker/internal/service/post_like"
	post_usecase "postMaker/internal/usecase/post"
	"strconv"
	"strings"
)

type Controller struct {
	useCase *post_usecase.UseCase
}

func NewController(useCase *post_usecase.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc Controller) GetPostList(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	userData, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	// Post Filter
	var filter post.Filter
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

	// Like filter
	var likeFilter post_like.Filter
	likeQuery := c.Request.URL.Query()

	likeLimitQ := likeQuery["limit"]
	if len(likeLimitQ) > 0 {
		queryInt, err := strconv.Atoi(likeLimitQ[0])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Limit must be a number",
				"status":  false,
			})
			return
		}

		likeFilter.Limit = &queryInt
	}

	likeOffsetQ := likeQuery["offset"]
	if len(likeOffsetQ) > 0 {
		queryInt, err := strconv.Atoi(likeOffsetQ[0])
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

	list, count, err := cc.useCase.GetPostList(ctx, filter, likeFilter, userData.Id)
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

func (cc Controller) GetPostById(c *gin.Context) {
	//tokenString := c.GetHeader("Authorization")
	//tokenString = strings.Split(tokenString, "Bearer ")[1]
	//
	//claims, err := auth.ParseToken(tokenString)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": err.Error(),
	//		"status":  false,
	//	})
	//	return
	//}

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

	detail, err := cc.useCase.GetPostById(ctx, id)
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

func (cc Controller) CreatePost(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	user, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	var data post.Create

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	data.UserId = user.Id

	ctx := context.Background()

	newData, err := cc.useCase.CreatePost(ctx, data)
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

func (cc Controller) UpdatePost(c *gin.Context) {
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

	var data post.Update

	err = c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	ctx := context.Background()

	newData, err := cc.useCase.UpdatePost(ctx, data)
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

func (cc Controller) DeletePost(c *gin.Context) {
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

	err = cc.useCase.DeletePost(ctx, id)
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
