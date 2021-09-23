package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/database"
	"go-gin-jwt/helpers"
	"go-gin-jwt/models"
	"net/http"
)

func UserRegistration(c *gin.Context){
	db := database.GetDB()
	ct := helpers.GetContentType(c)
	User := models.User{}

	if ct == "application/json" {
		c.ShouldBindJSON(&User)
	}else{
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": 			User.ID,
		"email":	 	User.Email,
		"full_name": 	User.FullName,
	})
}

func UserLogin(c *gin.Context)  {
	db := database.GetDB()
	ct := helpers.GetContentType(c)
	User := models.User{}

	if ct == "application/json" {
		c.ShouldBindJSON(&User)
	}else{
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	// JWT
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}