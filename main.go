package main

import (
	"myapi/database"
	"myapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.GET("/products_in_work", handlers.GetProductsInWork)
	r.GET("/products_in_work/:id", handlers.GetProductInWorkByID)
	r.POST("/products_in_work", handlers.CreateProductInWork)
	r.PUT("/products_in_work/:id", handlers.UpdateProductInWork)
	r.DELETE("/products_in_work/:id", handlers.DeleteProductInWork)

	r.Run(":8081")
}
