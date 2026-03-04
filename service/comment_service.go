package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoutesComment(rg *gin.RouterGroup) {
	// 嵌套在 posts 下：/posts/:id/comments
	posts := rg.Group("/posts")

	posts.GET("/:id/comments", getCommentsByPostID)
	posts.POST("/:id/comments", util.TokenAuthMiddleware(), createComment)
}

// createComment godoc
// @Summary create comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param body body model.CreateCommentReq true "comment body"
// @Success 201 {object} model.MComment
// @Failure 400 {object} model.GeneralMsg
// @Failure 401 {object} model.GeneralMsg
// @Security bearerAuth
// @Router /posts/{id}/comments [post]
func createComment(c *gin.Context) {
	// 1) 从 token 中拿 userId（你现在已经用 c.Set("userId") 了）
	uidAny, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing userId in token"})
		return
	}
	userID := uidAny.(int64)

	// 2) 解析 postId
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid post id"})
		return
	}

	// 3) 绑定 body
	var req model.CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json", "error": err.Error()})
		return
	}

	// 4) 插入评论
	created, err := repository.CreateComment(postID, userID, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// getCommentsByPostID godoc
// @Summary list comments by post id
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param page query int false "page" default(1)
// @Param size query int false "size" default(10)
// @Success 200 {array} model.MComment
// @Failure 400 {object} model.GeneralMsg
// @Router /posts/{id}/comments [get]
func getCommentsByPostID(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid post id"})
		return
	}

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

	comments, err := repository.ListCommentsByPostID(postID, size, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
