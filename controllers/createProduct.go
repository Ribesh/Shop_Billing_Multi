package controllers

import (
	"net/http"

	"github.com/Ribesh/gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var addproduct models.Product
		//var customer models.Customer
		if err := c.ShouldBindJSON(&addproduct); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		db.Create(&addproduct)
		//c.JSON(200, customer)
		c.JSON(200, addproduct)
	}
}
