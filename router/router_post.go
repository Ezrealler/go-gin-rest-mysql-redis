package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"JWT_REST_Gin_MySQL/repository"
)

func RegisterPostRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	// GET /api/v1/posts?page=1&size=10
	v1.GET("/posts", func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		sizeStr := c.DefaultQuery("size", "10")

		page, _ := strconv.Atoi(pageStr)
		size, _ := strconv.Atoi(sizeStr)

		if page < 1 {
			page = 1
		}
		if size < 1 {
			size = 10
		}
		if size > 50 {
			size = 50
		}
		offset := (page - 1) * size

		posts, err := repository.ListPosts(offset, size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 50000, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": posts})
	})

	// GET /api/v1/posts/:id
	v1.GET("/posts/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 10001, "msg": "invalid id"})
			return
		}

		post, err := repository.GetPostByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"code": 20001, "msg": "post not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": post})
	})
}
