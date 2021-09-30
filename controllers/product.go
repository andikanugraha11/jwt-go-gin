package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-gin-jwt/database"
	"go-gin-jwt/helpers"
	"go-gin-jwt/models"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context){
	db := database.GetDB()
	ct := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["id"].(float64))
	Product := models.Product{}

	if ct == "application/json" {
		c.ShouldBindJSON(&Product)
	}else{
		c.ShouldBind(&Product)
	}

	Product.UserId = userId

	err := db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context){
	db := database.GetDB()
	ct := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["id"].(float64))
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	if ct == "application/json" {
		c.ShouldBindJSON(&Product)
	}else{
		c.ShouldBind(&Product)
	}

	Product.UserId = userId
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)

}