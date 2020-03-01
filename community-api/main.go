package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB
var err error

func main() {
	router := gin.Default()

	db, _ = gorm.Open("mysql", "root:admin@tcp(localhost:3306)/cfwn2?charset=utf8&parseTime=True&loc=Local")

	db.LogMode(true)
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	defer db.Close()
	//db.AutoMigrate(&TProduct{})

	ConfigRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello ,mall-api"})
	})

	router.Run(":8000")
	log.Println("Staring successed")

}
