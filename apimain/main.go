package main

import (
	"apimain/controllers"
	"apimain/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", controllers.GetUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/products", controllers.GetProduct)
	r.GET("/products/:id", controllers.GetProductByID)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.GET("/products_in_work", controllers.GetProductsInWork)
	r.GET("/products_in_work/:id", controllers.GetProductInWorkByID)
	r.POST("/products_in_work", controllers.CreateProductInWork)
	r.PUT("/products_in_work/:id", controllers.UpdateProductInWork)
	r.DELETE("/products_in_work/:id", controllers.DeleteProductInWork)

	r.Run(":8080")
}
