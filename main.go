package main

import (
	"github.com/Ribesh/gorm/controllers"
	"github.com/Ribesh/gorm/database"
	"github.com/Ribesh/gorm/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = database.Connect()
}

func main() {

	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	apiproduct := api.Group("/products")
	{
		apiproduct.POST("/", controllers.CreateProduct(db))
	}
	apiItem := api.Group("/items")
	{
		apiItem.POST("/", controllers.AddItem(db))
		apiItem.GET("/:name", controllers.GetAllItems(db))
	}
	router.Run(":8090")
}
