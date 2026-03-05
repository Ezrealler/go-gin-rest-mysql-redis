package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoutesCategory ...
func RoutesCategory(rg *gin.RouterGroup) {
	cat := rg.Group("/categories")
	cat.GET("/", listCategories)
	cat.POST("/", util.TokenAuthMiddleware(), createCategory) // 先要求登录
}

// listCategories godoc
// @Summary list categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} model.MCategory
// @Failure 400 {object} map[string]string
// @Router /categories/ [get]
func listCategories(c *gin.Context) {
	res, err := repository.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// createCategory godoc
// @Summary create category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body model.CreateCategoryReq true "category body"
// @Success 201 {object} model.MCategory
// @Failure 401 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Security bearerAuth
// @Router /categories/ [post]
func createCategory(c *gin.Context) {
	var req model.CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json", "error": err.Error()})
		return
	}
	created, err := repository.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}
