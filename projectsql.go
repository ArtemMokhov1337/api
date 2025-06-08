package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func main() {
	dsn := "host=localhost user=postgres password=Kryptonite4ever dbname=mydb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	DB = database

	DB.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []User
		DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080") // localhost:8080
}
