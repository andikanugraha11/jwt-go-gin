package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/controllers"
)

func StartApplication() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	userRouter.POST("/register", controllers.UserRegistration)  // users/register
	userRouter.POST("/login", controllers.UserLogin)

	return r
}