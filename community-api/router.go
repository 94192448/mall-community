package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConfigRouter(r *gin.Engine) {
	r.GET("/api/product", getProduct)
}
