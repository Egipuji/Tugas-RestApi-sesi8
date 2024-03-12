package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"tugas-restapi-sesi8/database"
	"tugas-restapi-sesi8/models"

	"github.com/gin-gonic/gin"
)

func CreateOrders(ctx *gin.Context) {
	db := database.GetDb()

	orders := models.Orders{
		CustomerName: ctx.PostForm("customer_name"),
		OrderedAt:    time.Now(),
	}

	err := db.Create(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message_error": "faild create order",
		})
		return
	}

	quantity, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		fmt.Println(err)
		return
	}
	items := models.Items{
		ItemCode:    ctx.PostForm("item_code"),
		Description: ctx.PostForm("description"),
		Quantity:    uint(quantity),
		OrderID:     orders.OrderID,
	}

	err = db.Create(&items).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message_error": "faild create order",
		})
		return
	}

	orders.Items = []models.Items{items}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDb()

	orders := []models.Orders{}

	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message_error": "fail geting order with item",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func GetById(ctx *gin.Context) {
	db := database.GetDb()

	OrderID := ctx.Param("order_id")

	orders := models.Orders{}
	items := models.Items{}

	err := db.Where("order_id = ?", OrderID).First(&orders).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Where("order_id = ?", OrderID).First(&items).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orders.Items = []models.Items{items}
	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})

}

func UpdateOrders(ctx *gin.Context) {
	db := database.GetDb()

	OrderID := ctx.Param("order_id")

	orders := models.Orders{}
	custName := ctx.PostForm("customer_name")

	err := db.Model(&orders).Where("order_id = ?", OrderID).Updates(models.Orders{CustomerName: custName}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	quantity, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		quantity = 0
	}
	items := models.Items{
		ItemCode:    ctx.PostForm("item_code"),
		Description: ctx.PostForm("description"),
		Quantity:    uint(quantity),
		OrderID:     orders.OrderID,
	}

	err = db.Model(&items).Where("order_id = ?", OrderID).
		Updates(models.Items{ItemCode: items.ItemCode, Description: items.Description, Quantity: items.Quantity}).Error

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Where("order_id = ?", OrderID).First(&orders).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Where("order_id = ?", OrderID).First(&items).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orders.Items = []models.Items{items}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "succes",
		"orders":  orders,
	})

}

func DelteOrders(ctx *gin.Context) {
	db := database.GetDb()

	OrderID := ctx.Param("order_id")

	orders := models.Orders{}
	items := models.Items{}

	err := db.Where("order_id = ?", OrderID).Delete(&items).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed delete item",
		})
		return
	}

	err = db.Where("order_id = ?", OrderID).Delete(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed delete orders",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted orders",
	})
}
