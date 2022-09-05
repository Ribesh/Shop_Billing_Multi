package controllers

import (
	"github.com/Ribesh/gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var addItem models.Customer
		if err := ctx.ShouldBindJSON(&addItem); err != nil {
			ctx.JSON(500, err.Error())
		}
		db.Create(&addItem)
		ctx.JSON(200, addItem)
	}
}
