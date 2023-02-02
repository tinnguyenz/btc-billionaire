package main

import (
	"btc-billionaire/conf"
	"btc-billionaire/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	conf.ConnectDatabase()

	router.GET("/records", controllers.GetRecords)
	router.GET("/showHistory", controllers.ShowHistory)
	router.POST("/records", controllers.CreateRecord)
	router.Run()
}
