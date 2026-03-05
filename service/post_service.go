package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesPost ...
func RoutesPost(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.GET("/:id", getPostByID)
	posts.GET("/", getPosts)
	posts.POST("/", util.TokenAuthMiddleware(), createPost)
	posts.PUT("/", util.TokenAuthMiddleware(), updatePost)
	posts.DELETE("/:id", util.TokenAuthMiddleware(), deletePostByID)
}

// getPostByID godoc
// @Summary show Post by id
// @Description get string by ID
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /posts/{id} [get]
func getPostByID(c *gin.Context) {
	var post model.MPost
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	post, err = repository.GetPostByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)

}

// getPosts godoc
// @Summary show list post
// @Description get posts
// @Tags Post
// @Accept  json
// @Produce  json
// @Success 200 {array} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Router /posts/ [get]
func getPosts(c *gin.Context) {

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

	catStr := c.Query("category_id")
	var catID int64 = 0
	if catStr != "" {
		catID, _ = strconv.ParseInt(catStr, 10, 64)
	}
	posts, err := repository.ListPosts(offset, size, catID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)

}

// createPost godoc
// @Summary create post
// @Description add by json post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post body model.MPost true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /posts/ [post]
func createPost(c *gin.Context) {
	uidAny, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing userId in token"})
		return
	}
	userId := uidAny.(int64)

	var post model.MPost
	if post.CategoryID == 0 {
		post.CategoryID = 1
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json", "error": err.Error()})
		return
	}

	post.UserID = userId

	created, err := repository.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// updatePost godoc
// @Summary update post
// @Description update by json post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post body model.MPost true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /posts/ [put]
func updatePost(c *gin.Context) {

	var post model.MPost

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	pst, err := repository.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pst)
}

// deletePostByID godoc
// @Summary delete a post by id
// @Description delete post by ID
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID" Format(int64)
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /posts/{id} [delete]
func deletePostByID(c *gin.Context) {

	var post model.MPost

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeletePostByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, post)
}
