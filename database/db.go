package database

import (
	"fmt"
	"go-gin-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host 		= "127.0.0.1"
	user		= "postgres"
	password	= "postgres"
	port		= "5432"
	dbName		= "go_gin_jwt"
	db *gorm.DB
	err error
)

func InitDB()  {
	dsn := fmt.Sprintf("host=%s password=%s user=%s dbname=%s port=%s sslmode=disable",
		host, password, user, dbName, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Database error", err.Error())
	}

	fmt.Println("Database connected")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
