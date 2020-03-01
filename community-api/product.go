package main

import (
	//"log"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TProduct struct {
	//gorm.Model
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Ico        string `json:"ico"`
	WholePrice float32
	WholeNum   int
}

func getProduct(c *gin.Context) {
	var products []TProduct

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	var orderBy = c.DefaultQuery("order", "id desc")

	paginator := pagination.Paging(&pagination.Param{
		DB:      DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{orderBy},
		ShowSQL: true,
	}, &products)

	c.JSON(200, paginator)
}
