package main

import (
	"go-gin-jwt/database"
	"go-gin-jwt/router"
)

func main()  {
	database.InitDB()

	r := router.StartApplication()

	r.Run(":8080")
}