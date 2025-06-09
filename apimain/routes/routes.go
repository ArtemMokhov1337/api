package main

import (
	"apimain/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.GetUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
	products := r.Group("/products")
	{
		products.POST("/", controllers.CreateProduct)
		products.GET("/", controllers.GetProduct)
		products.GET("/:id", controllers.GetProductByID)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}
	productsInWork := r.Group("/productsInWork")
	{
		productsInWork.POST("/", controllers.CreateProductInWork)
		productsInWork.GET("/:id", controllers.GetProductInWorkByID)
		productsInWork.PUT("/:id", controllers.UpdateProductInWork)
		productsInWork.DELETE("/:id", controllers.DeleteProductInWork)
	}
}
