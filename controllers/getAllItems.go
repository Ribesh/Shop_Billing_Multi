package controllers

import (
	"log"

	"github.com/Ribesh/gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllItems(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allitems models.Customer
		//var allproduct models.Product
		name := ctx.Param("name")
		log.Println(name)
		//db.Debug().Find(&allproduct, &allitems, "customers.customer_product=products.product_id = ? AND customers.customer_name=?", name)
		//db.Where("customers.customer_product=products.product_id = ? AND customers.customer_name=?", name).Find(&allproduct, allitems)
		db.Where("customer_name = ?", name).Preload("Product").Find(&allitems)
		ctx.JSON(200, allitems)

	}
}
