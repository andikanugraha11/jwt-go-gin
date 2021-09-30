package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/controllers"
	"go-gin-jwt/middlewares"
)

func StartApplication() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	userRouter.POST("/register", controllers.UserRegistration)  // users/register
	userRouter.POST("/login", controllers.UserLogin)


	productRouter := r.Group("/products")
	productRouter.Use(middlewares.Authentication())
	productRouter.POST("/", controllers.CreateProduct) // products
	productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)

	return r
}