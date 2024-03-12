package main

import (
	"tugas-restapi-sesi8/controllers"
	"tugas-restapi-sesi8/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()

	router := gin.Default()

	router.POST("/orders", controllers.CreateOrders)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:order_id", controllers.GetById)
	router.PUT("/orders/:order_id", controllers.UpdateOrders)
	router.DELETE("/orders/:order_id", controllers.DelteOrders)

	router.Run("localhost:8080")
}
