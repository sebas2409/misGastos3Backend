package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pn "github.com/pusher/push-notifications-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"misGastos3Backend/database"
	"misGastos3Backend/domain"
	"strconv"
)

var db *gorm.DB
var beamsClient pn.PushNotifications

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	beamsClient, _ = pn.New(viper.GetString("INSTANCE_ID"), viper.GetString("SECRET_KEY"))
	if err != nil {
		return
	}
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
		sendMessage("Nuevo producto", fmt.Sprintf("Se ha agregado un nuevo producto: %v, %v$", product.Name, product.Price))
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
		sendMessage("Producto eliminado", "Se ha eliminado un producto")
		c.JSON(200, "Deleted")
	})

	err = server.Run()
	if err != nil {
		return
	}
}

func sendMessage(title, body string) {
	publishRequest := map[string]interface{}{
		"fcm": map[string]interface{}{
			"notification": map[string]interface{}{
				"title": title,
				"body":  body,
			},
		},
	}
	_, err := beamsClient.PublishToInterests([]string{"familia"}, publishRequest)
	if err != nil {
		return
	}
}
