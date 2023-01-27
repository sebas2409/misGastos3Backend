package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"misGastos3Backend/database"
	"misGastos3Backend/domain"
	"strconv"
)

var db *gorm.DB

func main() {

	server := gin.Default()
	db, _ = database.GetDb()

	server.POST("/save", func(c *gin.Context) {
		var product domain.Product
		err := c.BindJSON(&product)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = database.SaveProduct(db, product)
		if err != nil {
			return
		}
		data, _ := database.GetProducts(db)

		c.JSON(200, data)
	})

	server.GET("/get", func(c *gin.Context) {
		data, _ := database.GetProducts(db)
		c.JSON(200, data)
	})

	server.GET("/products/:month", func(c *gin.Context) {
		month := c.Param("month")
		data, _ := database.GetProductsByDate(db, month)
		fmt.Println(data)
		c.JSON(200, data)
	})

	server.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		id2, _ := strconv.Atoi(id)
		err := database.DeleteProduct(db, id2)
		if err != nil {
			return
		}
		c.JSON(200, "Deleted")
	})

	err := server.Run()
	if err != nil {
		return
	}
}
